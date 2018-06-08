package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ReplaceParam struct {
	SourceName string
	TargetName string
	MainPath   string
	InFile     bool
	InDir      bool
}

func main() {
	var r ReplaceParam
	r.MainPath = "/home/mert.acel/DEVELOPMENT/GIT/AOSP/KHADAS/device"
	r.SourceName = "kvim2"
	r.TargetName = "Acel"
	r.InFile = true
	r.WalkinPath()
}

func (r ReplaceParam) WalkinPath() {
	targetDir := r.MainPath
	filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			//fmt.Println(path)
			return nil
		}

		if r.InFile {
			r.replaceFile(path)
		}
		return nil
	})
}

func (r ReplaceParam) replaceFile(path string) {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(read))
	fmt.Println(path)

	newContents := strings.Replace(string(read), r.SourceName, r.TargetName, -1)

	fmt.Println(newContents)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
