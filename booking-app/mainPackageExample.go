package main

import "fmt"

// func testingPackages() { // not exported
// 	fmt.Println("Packagess working testings")
// }

func TestingPackages() { // export with capitalising it
	fmt.Println("Packagess working testings")
}
