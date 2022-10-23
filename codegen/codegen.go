package main;

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("package main; func flows() string { return `")
	bytes, _ := ioutil.ReadFile(os.Args[1])
	fmt.Println((string)(bytes))
	fmt.Println("`}")
}
