package main

import (
	"fmt"
	"multipleFileTestGo/testPackage"
)

func main() {
	printTest()
	testPackage.Test()
	fmt.Println("Worked!")
}
