package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	err := testDirTree(os.Stdout, ".", true)
	if err != nil {
		panic(err.Error())
	}
}

func testDirTree(out io.Writer, path string, printFiles bool) (error error)  {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("├───", file.Name())
			err := testDirTree(os.Stdout, ".", true)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("├───", file.Name())
	}
	return nil
}