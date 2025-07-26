package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
		fmt.Println("File name required!");
		return
	}
    var fileName string = os.Args[1];
    
	file, err := os.Create(fileName);
	if err != nil {
         fmt.Println("Error ", err)
		 return
	}

	defer file.Close()

	println(fileName, " created successfully")
}