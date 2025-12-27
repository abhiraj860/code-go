package main 

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("../LLD/airlinemanagementsystem/aircraft.go")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			panic(err)
		}
	} 
	fmt.Printf("File name: %s \n",info.Name())
	fmt.Printf("File Size: %d \n", info.Size())
	fmt.Printf("File permission: %s \n", info.Mode())
	fmt.Printf("Last Modified: %s \n", info.ModTime())	
}