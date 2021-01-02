# PortScanner Module

This is a port scanner module for Go, I wrote this for another project I'm working on.
Still have a lot to do but this is a very early basic version.

## Getting Started

### Installation

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/0xjbb/portscanner
```

### Examples 

```go
func main(){
	
	ports := []int{1,2,3,4,5,6,7,8,9}

	Scanner := portscanner.New("127.0.0.1", ports)

	err := Scanner.Run()

	if err != nil {
		fmt.Println(err)
	}

	for _,v := range Scanner.GetResults() {
		fmt.Println("Open: ", v)
	}
}
```