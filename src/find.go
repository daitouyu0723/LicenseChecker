package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const samplePath string = "/home/wangwei/workspace/go/LicenseChecker/sample"

func main() {

	err := filepath.Walk(samplePath, walk)
	if err != nil {
		log.Fatal(err)
	}
}

func walk(path string, f os.FileInfo, err error) error {
	fmt.Println(path)

	if f.IsDir() {
		return nil
	}

	fc, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", md5.Sum(fc))
	return nil
}
