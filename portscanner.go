package portscanner

import (
	"fmt"
	"strconv"
	"strings"
)

type PortScanner struct{
	Host string
	Ports string
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
	ports, err := p.ParsePorts()

	if err != nil{
		return err
	}

	// scan ports
	for _, val := range ports{
		fmt.Println(val)


	}
}

func (p PortScanner) GetResults(){}
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

// Just checks if the port is valid.
func checkValidPort(port string) bool{
	// How does this even work loool, love go <3
	if port >= "0" && port <= "65535"{
		return true
	}
	return false
}
func makeRange(start int, count int) []string{
	sli := make([]string, count + 1)

	for i := start; i <= count;i++ {
		sli[i] = strconv.Itoa(i)
	}
	return sli[start:]
}