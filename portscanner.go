package main

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

func New(host string, ports string) PortScanner{
	return PortScanner{
		Host: host,
		Ports: ports,
	}
}

func (p *PortScanner) Run() error{
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

func (p *PortScanner) GetResults() []string {
	return p.Results
}
func (p *PortScanner) ParsePorts() []string {
	if p.Ports == "all"{
		return makeRange(0,65535)
	}

	ports := strings.Split(p.Ports, ",") // parse comma separated.

	var rPorts []string

	for _,port := range ports {

		if checkValidPort(port) == true{
			rPorts = append(rPorts, port)
		}
	}

	return rPorts
}

func (p *PortScanner) ConnectScan(port string){
	address := fmt.Sprintf("%s:%s", p.Host, port)
	conn, err := net.Dial("tcp", address)

	if err != nil{
		return
	}

	_ = conn.Close()

	p.Results = append(p.Results, port)
}

// helper functions
// Just checks if the port is valid.
func checkValidPort(port string) bool{
	// How does this even work loool, love go <3
	p, _ := strconv.Atoi(port)

	if p >= 0 && p <= 65535{
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

func main(){

	test := New("127.0.0.2", "80,8080")

	err := test.Run()

	if err != nil {
		fmt.Println(err)
	}

	for _,v := range test.Results {
		fmt.Println("Open: ", v)
	}
}