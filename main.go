package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("../Buffers/huge-allocation.s")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}
		fmt.Println("Error: ", err)
		return
	}
	
	permissions := fileInfo.Mode().Perm()
	permissionString := fmt.Sprintf("%o", permissions)
	fmt.Printf("Permission %s\n", permissionString)
}