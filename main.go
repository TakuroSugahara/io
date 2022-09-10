package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.FileオブジェクトをOpen関数か何かで得てるものとする
	// os.FileはREAD専用のfile descriptorだったり、I/Oできるものだったりする

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("get wd error", err.Error())
	}

	file, err := os.Open(fmt.Sprintf("%s/text.txt", wd))
	if err != nil {
		log.Fatalln("get wd error", err.Error())
	}

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("fail to read file", err)
	}

	fmt.Printf("read %d bytes: \n", count)
	fmt.Println(string(data[:count]))
}
