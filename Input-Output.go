/*
	Go File I/O -> In os.file Objects are used for file manipulations. os.file Objects are also called fileHandles.
	Open Function which is in s=os Package is usd to open fles in Go. ReadFile() in the io/ioutil package is used
	to read the file. This method returns []byte array of read bytes. file.WriteString method can be used to write
	to the file.

	We use defer file.close() right after opening the file to make sure that the file is closed as soon as the function
	completes. If a file does not exists or the program has not the sufficient rights to open the file then inputFile,
	inputFile = os.Open("input.dat") results in an error.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi...There")
	file.Close()
	stream, err := ioutil.ReadFile("file.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)
}
