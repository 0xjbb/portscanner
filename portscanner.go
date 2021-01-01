package portscanner

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)
/*
* @TODO Refactor to add concurrency, current version is just to get the module sorted for another project.
 */
type PortScanner struct{
	Host string
	Ports string
	Results []string
	//Type string // null/syn/connect
	//Timeout int
}

func New(host string, ports string) *PortScanner{
	return &PortScanner{
		Host: host,
		Ports: ports,
	}
}

func (p PortScanner) Run() error{
	ports := p.ParsePorts()

	if len(ports) == 0 {
		return errors.New("no ports given")
	}

	// scan ports
	for _, val := range ports{
		p.ConnectScan(val)
	}

	return nil
}

func (p PortScanner) ParsePorts() []string {
	if p.Ports == "all"{
		return makeRange(0,65535)
	}

	ports := strings.Split(p.Ports, ",") // parse comma separated.

	var rPorts []string

	for _,port := range ports {
		if checkValidPort(port){
			rPorts = append(rPorts, port)
		}
	}

	return rPorts
}

func (p PortScanner) ConnectScan(port string){
	address := fmt.Sprintf("%s:%s", p.Host, port)

	conn, err := net.Dial("tcp", address)

	if err != nil{
		// not open
		return
	}

	_ = conn.Close()

	p.Results = append(p.Results, port)
}

// helper functions
// Just checks if the port is valid.
func checkValidPort(port string) bool{
	// How does this even work loool, love go <3
	if port >= "0" && port <= "65535"{
		return true
	}
	return false
}

// why doesn't Go have this built-in ree
func makeRange(start int, count int) []string{
	sli := make([]string, count + 1)

	for i := start; i <= count;i++ {
		sli[i] = strconv.Itoa(i)
	}
	return sli[start:]
}