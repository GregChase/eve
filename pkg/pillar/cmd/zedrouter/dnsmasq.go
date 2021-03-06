// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// dnsmasq configlets for overlay and underlay interfaces towards domU

package zedrouter

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/lf-edge/eve/pkg/pillar/agentlog"
	"github.com/lf-edge/eve/pkg/pillar/types"
	log "github.com/sirupsen/logrus"
)

// XXX inotify seems to stop reporting any changes in some cases
// XXX avoid by start and stop dnsmasq when we add entries
// XXX KALYAN - We need to set this to have DHCP working wiht Network instances.
//		Turning this flag on temporarily till we figure out whats happening.
const dnsmasqStopStart = true // XXX change? remove?

const dnsmasqStatic = `
# Automatically generated by zedrouter
except-interface=lo
bind-interfaces
log-queries
log-dhcp
no-hosts
no-ping
bogus-priv
stop-dns-rebind
rebind-localhost-ok
neg-ttl=10
`

func dnsmasqConfigFile(bridgeName string) string {
	cfgFilename := "dnsmasq." + bridgeName + ".conf"
	return cfgFilename
}

func dnsmasqConfigPath(bridgeName string) string {
	cfgFilename := dnsmasqConfigFile(bridgeName)
	cfgPathname := runDirname + "/" + cfgFilename
	return cfgPathname
}

func dnsmasqDhcpHostDir(bridgeName string) string {
	dhcphostsDir := runDirname + "/dhcp-hosts." + bridgeName
	return dhcphostsDir
}

// createDnsmasqConfiglet
// When we create a linux bridge we set this up
// Also called when we need to update the ipsets
func createDnsmasqConfiglet(
	bridgeName string, bridgeIPAddr string,
	netconf *types.NetworkInstanceConfig, hostsDir string,
	ipsets []string, Ipv4Eid bool) {

	log.Infof("createDnsmasqConfiglet(%s, %s) netconf %v, ipsets %v\n",
		bridgeName, bridgeIPAddr, netconf, ipsets)

	cfgPathname := dnsmasqConfigPath(bridgeName)
	// Delete if it exists
	if _, err := os.Stat(cfgPathname); err == nil {
		if err := os.Remove(cfgPathname); err != nil {
			errStr := fmt.Sprintf("createDnsmasqConfiglet %v",
				err)
			log.Errorln(errStr)
		}
	}
	file, err := os.Create(cfgPathname)
	if err != nil {
		log.Fatal("createDnsmasqConfiglet failed ", err)
	}
	defer file.Close()

	// Create a dhcp-hosts directory to be used when hosts are added
	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)

	file.WriteString(dnsmasqStatic)
	for _, ipset := range ipsets {
		file.WriteString(fmt.Sprintf("ipset=/%s/ipv4.%s,ipv6.%s\n",
			ipset, ipset, ipset))
	}
	file.WriteString(fmt.Sprintf("pid-file=/var/run/dnsmasq.%s.pid\n",
		bridgeName))
	file.WriteString(fmt.Sprintf("interface=%s\n", bridgeName))
	isIPv6 := false
	if bridgeIPAddr != "" {
		ip := net.ParseIP(bridgeIPAddr)
		if ip == nil {
			log.Fatalf("createDnsmasqConfiglet failed to parse IP %s",
				bridgeIPAddr)
		}
		isIPv6 = (ip.To4() == nil)
		file.WriteString(fmt.Sprintf("listen-address=%s\n",
			bridgeIPAddr))
	} else {
		// XXX error if there is no bridgeIPAddr?
	}
	file.WriteString(fmt.Sprintf("hostsdir=%s\n", hostsDir))
	file.WriteString(fmt.Sprintf("dhcp-hostsdir=%s\n", dhcphostsDir))

	ipv4Netmask := "255.255.255.0" // Default unless there is a Subnet
	dhcpRange := bridgeIPAddr      // Default unless there is a DhcpRange

	// By default dnsmasq advertizes a router (and we can have a
	// static router defined in the NetworkInstanceConfig).
	// To support airgap networks we interpret gateway=0.0.0.0
	// to not advertize ourselves as a router. Also,
	// if there is not an explicit dns server we skip
	// advertising that as well.
	advertizeRouter := true
	var router string

	if Ipv4Eid {
		advertizeRouter = false
	} else if netconf.Gateway != nil {
		if netconf.Gateway.IsUnspecified() {
			advertizeRouter = false
		} else {
			router = netconf.Gateway.String()
		}
	} else if bridgeIPAddr != "" {
		router = bridgeIPAddr
	} else {
		advertizeRouter = false
	}
	if netconf.DomainName != "" {
		if isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:domain-search,%s\n",
				netconf.DomainName))
		} else {
			file.WriteString(fmt.Sprintf("dhcp-option=option:domain-name,%s\n",
				netconf.DomainName))
		}
	}
	advertizeDns := false
	if Ipv4Eid {
		advertizeDns = true
	}
	for _, ns := range netconf.DnsServers {
		advertizeDns = true
		file.WriteString(fmt.Sprintf("dhcp-option=option:dns-server,%s\n",
			ns.String()))
	}
	if netconf.NtpServer != nil {
		file.WriteString(fmt.Sprintf("dhcp-option=option:ntp-server,%s\n",
			netconf.NtpServer.String()))
	}
	if netconf.Subnet.IP != nil {
		ipv4Netmask = net.IP(netconf.Subnet.Mask).String()
	}
	// Special handling for IPv4 EID case to avoid ARP for EIDs.
	// We add a router for the BridgeIPAddr plus a subnet route
	// for the EID subnet, and no default route by clearing advertizeRouter
	// above. We configure an all ones netmask. In addition, since the
	// default broadcast address ends up being the bridgeIPAddr, we force
	// a bogus one as the first .0 address in the subnet.
	//
	if Ipv4Eid {
		file.WriteString("dhcp-option=option:netmask,255.255.255.255\n")
		// Onlink aka ARPing route for our IP
		route1 := fmt.Sprintf("%s/32,0.0.0.0", bridgeIPAddr)
		var route2 string
		var broadcast string
		if netconf.Subnet.IP != nil {
			route2 = fmt.Sprintf(",%s,%s", netconf.Subnet.String(),
				bridgeIPAddr)
			broadcast = netconf.Subnet.IP.String()
		}
		file.WriteString(fmt.Sprintf("dhcp-option=option:classless-static-route,%s%s\n",
			route1, route2))
		// Broadcast address option
		if broadcast != "" {
			file.WriteString(fmt.Sprintf("dhcp-option=28,%s\n",
				broadcast))
		}
	} else if netconf.Subnet.IP != nil {
		file.WriteString(fmt.Sprintf("dhcp-option=option:netmask,%s\n",
			ipv4Netmask))
	}
	if advertizeRouter {
		// IPv6 XXX needs to be handled in radvd
		if !isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:router,%s\n",
				router))
		}
	} else {
		log.Infof("createDnsmasqConfiglet: no router\n")
		if !isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:router\n"))
		}
		if !advertizeDns {
			// Handle isolated network by making sure
			// we are not a DNS server. Can be overridden
			// with the DnsServers above
			log.Infof("createDnsmasqConfiglet: no DNS server\n")
			file.WriteString(fmt.Sprintf("dhcp-option=option:dns-server\n"))
		}
	}
	if netconf.DhcpRange.Start != nil {
		dhcpRange = netconf.DhcpRange.Start.String()
	}
	if isIPv6 {
		file.WriteString(fmt.Sprintf("dhcp-range=::,static,0,10m\n"))
	} else {
		file.WriteString(fmt.Sprintf("dhcp-range=%s,static,%s,10m\n",
			dhcpRange, ipv4Netmask))
	}
}

func addhostDnsmasq(bridgeName string, appMac string, appIPAddr string,
	hostname string) {

	log.Infof("addhostDnsmasq(%s, %s, %s, %s)\n", bridgeName, appMac,
		appIPAddr, hostname)
	if dnsmasqStopStart {
		stopDnsmasq(bridgeName, true, false)
	}
	ip := net.ParseIP(appIPAddr)
	if ip == nil {
		log.Fatalf("addhostDnsmasq failed to parse IP %s", appIPAddr)
	}
	isIPv6 := (ip.To4() == nil)
	suffix := ".inet"
	if isIPv6 {
		suffix += "6"
	}

	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)
	cfgPathname := dhcphostsDir + "/" + appMac + suffix

	file, err := os.Create(cfgPathname)
	if err != nil {
		log.Fatal("addhostDnsmasq failed ", err)
	}
	defer file.Close()
	if isIPv6 {
		file.WriteString(fmt.Sprintf("%s,[%s],%s\n",
			appMac, appIPAddr, hostname))
	} else {
		file.WriteString(fmt.Sprintf("%s,id:*,%s,%s\n",
			appMac, appIPAddr, hostname))
	}
	file.Close()
	if dnsmasqStopStart {
		startDnsmasq(bridgeName)
	}
}

func removehostDnsmasq(bridgeName string, appMac string, appIPAddr string) {

	log.Infof("removehostDnsmasq(%s, %s, %s)\n",
		bridgeName, appMac, appIPAddr)
	if dnsmasqStopStart {
		stopDnsmasq(bridgeName, true, false)
	}
	ip := net.ParseIP(appIPAddr)
	if ip == nil {
		log.Fatalf("removehostDnsmasq failed to parse IP %s", appIPAddr)
	}
	isIPv6 := (ip.To4() == nil)
	suffix := ".inet"
	if isIPv6 {
		suffix += "6"
	}

	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)

	cfgPathname := dhcphostsDir + "/" + appMac + suffix
	if _, err := os.Stat(cfgPathname); err != nil {
		log.Infof("removehostDnsmasq(%s, %s) failed: %s\n",
			bridgeName, appMac, err)
	} else {
		if err := os.Remove(cfgPathname); err != nil {
			errStr := fmt.Sprintf("removehostDnsmasq %v", err)
			log.Errorln(errStr)
		}
	}
	if dnsmasqStopStart {
		startDnsmasq(bridgeName)
	}
}

func deleteDnsmasqConfiglet(bridgeName string) {

	log.Infof("deleteDnsmasqConfiglet(%s)\n", bridgeName)
	cfgPathname := dnsmasqConfigPath(bridgeName)
	if _, err := os.Stat(cfgPathname); err == nil {
		if err := os.Remove(cfgPathname); err != nil {
			errStr := fmt.Sprintf("deleteDnsmasqConfiglet %v",
				err)
			log.Errorln(errStr)
		}
	}
	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)
	if err := RemoveDirContent(dhcphostsDir); err != nil {
		errStr := fmt.Sprintf("deleteDnsmasqConfiglet %v", err)
		log.Errorln(errStr)
	}
}

func RemoveDirContent(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := dir + "/" + file.Name()
		log.Infoln("RemoveDirConent found ", filename)
		err = os.RemoveAll(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run this:
//    DMDIR=/opt/zededa/bin/
//    ${DMDIR}/dnsmasq -b -C /var/run/zedrouter/dnsmasq.${BRIDGENAME}.conf
func startDnsmasq(bridgeName string) {

	log.Infof("startDnsmasq(%s)\n", bridgeName)
	cfgPathname := dnsmasqConfigPath(bridgeName)
	name := "nohup"
	//    XXX currently running as root with -d above
	args := []string{
		"/opt/zededa/bin/dnsmasq",
		"-d",
		"-C",
		cfgPathname,
	}
	logFilename := fmt.Sprintf("dnsmasq.%s", bridgeName)
	logf, err := agentlog.InitChild(logFilename)
	if err != nil {
		log.Fatalf("startDnsmasq agentlog failed: %s\n", err)
	}
	w := bufio.NewWriter(logf)
	ts := time.Now().Format(time.RFC3339Nano)
	fmt.Fprintf(w, "%s Starting %s %v\n", ts, name, args)
	cmd := exec.Command(name, args...)
	cmd.Stderr = logf
	log.Infof("Calling command %s %v\n", name, args)
	go cmd.Run()
}

//    pkill -u nobody -f dnsmasq.${BRIDGENAME}.conf
func stopDnsmasq(bridgeName string, printOnError bool, delConfiglet bool) {

	log.Infof("stopDnsmasq(%s)\n", bridgeName)
	cfgFilename := dnsmasqConfigFile(bridgeName)
	// XXX currently running as root with -d above
	pkillUserArgs("root", cfgFilename, printOnError)

	if delConfiglet {
		deleteDnsmasqConfiglet(bridgeName)
	}
}
