package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("lets work on file")
	content := "helloo this is content"
	file, err := os.Create("./mylocalfile.txt")
	Checkerr(err)
	length, err := io.WriteString(file, content)
	Checkerr(err)
	log.Println(length)
	file.Close()
	readFile("./mylocalfile.txt")
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	Checkerr(err)
	log.Println(string(databyte))

}
func Checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
