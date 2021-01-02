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
	Scanner := portscanner.New("127.0.0.2", "80,8080")

	err := Scanner.Run()

	if err != nil {
		fmt.Println(err)
	}

	for _,v := range Scanner.GetResults() {
		fmt.Println("Open: ", v)
	}
}
```