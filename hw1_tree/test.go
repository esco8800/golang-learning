package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var res string
var org int = 0

func main() {
	err := testDirTree(os.Stdout, "testdata", true)
	if err != nil {
		panic(err.Error())
	}
}

func testDirTree(out io.Writer, path string, printFiles bool) (error error)  {

	var level int = -1
	var pos int = 1

	err := testDirTreeRecursive(os.Stdout, path, printFiles, level, pos)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	return nil
}

func testDirTreeRecursive(out io.Writer, path string, printFiles bool, level, pos int) (error error)  {
	files, err := ioutil.ReadDir(path)
	lenFiles := len(files)
	if err != nil {
		log.Fatal(err)
	}

	level++

	for i := 0; i < lenFiles; i++ {

		if lenFiles == i + 1 {
			pos = 2
		} else  {
			pos = 1
		}

		if files[i].IsDir() {
			makeByLevelAndPos(level, pos, files[i].Name())
			err2 := testDirTreeRecursive(os.Stdout, concat(path + "/", files[i].Name()), printFiles, level, pos)
			if err2 != nil {
				log.Fatal(err)
			}
		} else {
			if printFiles {
				makeByLevelAndPos(level, pos, files[i].Name())
			}
		}
	}
	return nil
}

func makeByLevelAndPos(level, pos int, name string) {
	for i := 0; i < level; i++ {

		if org != 0  && org <= i {
			res = concat(res, "\t")
		} else {
			res = concat(res, "|\t")
		}

	}

	if level == org {
		org = 0
	}

	if pos == 1 {
		res = concat(res, "├───" + name + "\n")
	}
	if pos == 2 {
		res = concat(res, "└───" + name + "\n")
		org = level
	}
}

func concat(str, concat string) string  {
	return str + concat
}