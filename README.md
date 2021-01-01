# PortScanner Module

This is a port scanner module for Go, I wrote this for another project i'm working on.

## Getting Started



### Installation and Usage


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