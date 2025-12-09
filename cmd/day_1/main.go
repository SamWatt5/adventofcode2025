package main

import (
	"fmt"
	"bufio" 
	"os"
	"strings"
)

func check(err error){
	if err != nil{
		panic(err)
	}
}

func openFile(filename string){
	file, err := os.ReadFile(filename)
	check(err)

	for _, data := range file {
		if !bytes.Equal([]byr)
	}
}



func main(){
	openFile(input.txt)
}

