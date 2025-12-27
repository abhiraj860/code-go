package main 

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("../LLD/airlinemanagementsystem")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			panic(err)
		}
	} 
	fmt.Printf("File name: %s \n",info.Name())
	fmt.Printf("File Size: %d \n", info.Size())
	fmt.Printf("File permission: %o \n", info.Mode().Perm())
	fmt.Printf("Last Modified: %s \n", info.ModTime())
	fmt.Printf("File is regular %v \n", info.Mode().IsRegular())	
	fmt.Printf("File is directory %v \n", info.IsDir())	
	fmt.Printf("File is directory %v \n", os.ModeSymlink)	
}