package portscanner

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/ip_helpers"

	log "github.com/sirupsen/logrus"
	"net"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// IPv4 is the type used for IP addresses.
//type IPv4 ip_helpers.IPv4

type PortList struct {
	ServiceName string `json:"service_name"`
	Port        string `json:"port"`
}

type Host struct {
	IP    string     `json:"ip"`
	Ports []PortList `json:"ports"`
}

type Hosts struct {
	Hosts []Host `json:"hosts"`
}

// IPScanner scans all IP addresses in ipList for every port in portList.
func IPScanner(ips []string, portStr []string, printResults bool) (hostsFound Hosts) {
	var ipList []ip_helpers.IPv4
	var portList []string
	var wg sync.WaitGroup

	if len(portStr) == 1 {
		portList = ParsePortList(portStr[0])
	} else {
		portList = portStr
	}
	if len(ips) == 0 {
		ipList = append(ipList, ip_helpers.IPv4{127, 0, 0, 1})
	} else {
		for _, i := range ips {
			if strings.Contains(i, "-") {
				ipList = append(ipList, ip_helpers.ParseIPSequence(i)...)
				fmt.Println(ipList)
			} else {
				ip := ip_helpers.ToIPv4(i)
				if ip.IsValid() {
					ipList = append(ipList, ip)
				}
			}
		}
	}
	for _, ip := range ipList {
		wg.Add(1)
		go func(ip ip_helpers.IPv4) {
			defer wg.Done()
			ports := PortScanner(ip, portList)
			if len(ports) > 0 {
				var pl []PortList
				for _, port := range ports {
					pl = append(pl, PortList{
						Port:        port,
						ServiceName: portShortList[port],
					})
				}
				hostsFound.Hosts = append(hostsFound.Hosts, Host{IP: ip.ToString(), Ports: pl})
				if printResults {
					PresentResults(ip, ports)
				}
			}
		}(ip)
	}
	wg.Wait()
	return hostsFound
}

// ParsePortList gets a port list from its port entry in arguments.
func ParsePortList(rawPorts string) []string {
	var ports []string
	individuals, _ := regexp.Compile("([0-9]+)[,]*")
	series, _ := regexp.Compile("([0-9]+)[-]([0-9]+)")

	// For individual ports, separated by ','
	lIndividuals := individuals.FindAllStringSubmatch(rawPorts, -1)

	// For sequence ports, using '-'
	lSeries := series.FindAllStringSubmatch(rawPorts, -1)

	if len(lSeries) > 0 {
		for _, s := range lSeries {
			init, _ := strconv.Atoi(s[1])
			end, _ := strconv.Atoi(s[2])
			for i := init + 1; i < end; i++ {
				ports = append(ports, strconv.Itoa(i))
			}
		}
	}
	for _, port := range lIndividuals {
		ports = append(ports, port[1])
	}
	sort.Strings(ports)
	return ports
}

// PortScanner scans IP:port pairs looking for open ports on IP addresses.
func PortScanner(ip ip_helpers.IPv4, portList []string) []string {
	var open []string
	for _, port := range portList {
		conn, err := net.DialTimeout("tcp",
			ip.ToString()+":"+port,
			300*time.Millisecond)
		if err == nil {
			conn.Close()
			open = append(open, port)
		}
	}
	return open
}

// PresentResults presents all results in console.
func PresentResults(ip ip_helpers.IPv4, ports []string) int {
	log.Println(" \n>" + ip.ToString())
	log.Println(" Port:	Description:")
	for _, port := range ports {
		log.Println(" " + port + "\t" + portShortList[port])
	}
	return 0
}
