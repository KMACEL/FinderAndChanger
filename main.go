package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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
	args := os.Args

	if len(args) < 5 {
		fmt.Println("Please Control Param")
		fmt.Println("Example : ")
		fmt.Println("./Recursive_Finder /your/find/path source_name target_name in_file in_directory")
		fmt.Println("./Recursive_Finder /home/mert.acel/DEVELOPMENT/GIT/AOSP test Acel true false")
		return
	}
	fileBool, _ := strconv.ParseBool(args[4])
	dirBool, _ := strconv.ParseBool(args[5])

	if fileBool == dirBool {
		fmt.Println("./Recursive_Finder /your/find/path source_name target_name in_file in_directory")
		fmt.Println("Not Equal target_name in_file in_directory")
		return
	}

	r.MainPath = args[1]
	r.SourceName = args[2]
	r.TargetName = args[3]
	r.InFile = fileBool
	r.InDir = dirBool

	/*r.MainPath = "/home/mert.acel/DEVELOPMENT/GIT/AOSP/KHADAS"
	r.SourceName = "test"
	r.TargetName = "acel1"
	r.InFile = false
	r.InDir = true*/

	r.WalkinPath()
}

var folderSize int

func (r ReplaceParam) WalkinPath() {
	targetDir := r.MainPath

	filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		beginFolder, lastFolder := filepath.Split(path)

		if r.InDir {
			if strings.Contains(lastFolder, r.SourceName) {
				fmt.Println(beginFolder, "-", lastFolder)
				if !info.IsDir() {
					os.Rename(path, beginFolder+strings.Replace(lastFolder, r.SourceName, r.TargetName, -1))
				} else {
					folderSize++
				}
			}
		}

		if lastFolder != ".git" {
			if !info.IsDir() {
				if r.InFile {
					r.replaceFile(path)
				}
			}
		}
		return nil
	})

	fmt.Println("Folder Size : ", folderSize)
	for i := 0; i < folderSize; i++ {
		filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
			beginFolder, lastFolder := filepath.Split(path)
			if info.IsDir() {
				if strings.Contains(lastFolder, r.SourceName) {
					fmt.Println("Folder : ", path)
					err := os.Rename(path, beginFolder+strings.Replace(lastFolder, r.SourceName, r.TargetName, -1))
					if err != nil {
						fmt.Println(err.Error())
					}
					return nil
				}
				return nil
			}
			return nil
		})
	}
}

func (r ReplaceParam) replaceFile(path string) {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if strings.Contains(string(read), r.SourceName) {
		fmt.Println(path)
		newContents := strings.Replace(string(read), r.SourceName, r.TargetName, -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}
}
