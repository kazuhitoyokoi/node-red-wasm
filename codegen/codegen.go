package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("package main")
	fmt.Println("")
	fmt.Println("func flows() string { return `")
	bytes, _ := ioutil.ReadFile(os.Args[1])
	fmt.Println((string)(bytes))
	fmt.Println("`}")
}
