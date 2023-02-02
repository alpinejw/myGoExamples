package main

import "fmt"
import "structure_test/internal/pkg1"
import "structure_test/internal/pkg2"
import "structure_test/internal/pkg3"

func main() {
	fmt.Println("hello world")
	fmt.Println("The result is ", pkgone.MyFunc(5))
	fmt.Println("Package 2 says:", pkgtwo.ReturnLocalVar())
	fmt.Println("Package 3 adds two ints: 3 + 6 =", hellothree.AddTwoInt(3, 6))
}
