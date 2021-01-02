package portscanner

import (
	"errors"
	"fmt"
	"net"
)
/*
* @TODO Refactor to add concurrency, current version is just to get the module sorted for another project.
 */
type PortScanner struct{
	Host string
	Ports []int
	Results []int
	//Type string // null/syn/connect
	//Timeout int
}

func New(host string, ports []int) PortScanner{
	return PortScanner{
		Host: host,
		Ports: ports,
	}
}

func (p *PortScanner) Run() error{

	if len(p.Ports) == 0 {
		return errors.New("no ports given")
	}

	// scan ports
	for _, port := range p.Ports{
		// switch p.Type{}
		if checkValidPort(port){
			p.connectScan(port)
		}
	}

	return nil
}

func (p *PortScanner) GetResults() []int {
	return p.Results
}

func (p *PortScanner) connectScan(port int){
	address := fmt.Sprintf("%s:%d", p.Host, port)
	conn, err := net.Dial("tcp", address)

	if err != nil{
		return
	}

	_ = conn.Close()

	p.Results = append(p.Results, port)
}

// helper functions
// Just checks if the port is valid.
func checkValidPort(port int) bool{
	if port >= 0 && port <= 65535{
		return true
	}
	return false
}
