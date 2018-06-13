package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*
███████╗██╗███╗   ██╗██████╗ ███████╗██████╗          █████╗ ███╗   ██╗██████╗          ██████╗██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███████╗██████╗
██╔════╝██║████╗  ██║██╔══██╗██╔════╝██╔══██╗        ██╔══██╗████╗  ██║██╔══██╗        ██╔════╝██║  ██║██╔══██╗████╗  ██║██╔════╝ ██╔════╝██╔══██╗
█████╗  ██║██╔██╗ ██║██║  ██║█████╗  ██████╔╝        ███████║██╔██╗ ██║██║  ██║        ██║     ███████║███████║██╔██╗ ██║██║  ███╗█████╗  ██████╔╝
██╔══╝  ██║██║╚██╗██║██║  ██║██╔══╝  ██╔══██╗        ██╔══██║██║╚██╗██║██║  ██║        ██║     ██╔══██║██╔══██║██║╚██╗██║██║   ██║██╔══╝  ██╔══██╗
██║     ██║██║ ╚████║██████╔╝███████╗██║  ██║        ██║  ██║██║ ╚████║██████╔╝        ╚██████╗██║  ██║██║  ██║██║ ╚████║╚██████╔╝███████╗██║  ██║
╚═╝     ╚═╝╚═╝  ╚═══╝╚═════╝ ╚══════╝╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝          ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝
*/

/*
██████╗ ███████╗██████╗ ██╗      █████╗  ██████╗███████╗        ██████╗  █████╗ ██████╗  █████╗ ███╗   ███╗
██╔══██╗██╔════╝██╔══██╗██║     ██╔══██╗██╔════╝██╔════╝        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗████╗ ████║
██████╔╝█████╗  ██████╔╝██║     ███████║██║     █████╗          ██████╔╝███████║██████╔╝███████║██╔████╔██║
██╔══██╗██╔══╝  ██╔═══╝ ██║     ██╔══██║██║     ██╔══╝          ██╔═══╝ ██╔══██║██╔══██╗██╔══██║██║╚██╔╝██║
██║  ██║███████╗██║     ███████╗██║  ██║╚██████╗███████╗        ██║     ██║  ██║██║  ██║██║  ██║██║ ╚═╝ ██║
╚═╝  ╚═╝╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝ ╚═════╝╚══════╝        ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝
*/

//ReplaceParam is the data to be processed is the hosting structure
type ReplaceParam struct {
	SourceName string
	TargetName string
	MainPath   string
	InFile     bool
	InDir      bool
}

var (
	path       = ""
	sourceName = ""
	targetName = ""
	inFile     = false
	inDir      = false
)

/*
██╗███╗   ██╗██╗████████╗
██║████╗  ██║██║╚══██╔══╝
██║██╔██╗ ██║██║   ██║
██║██║╚██╗██║██║   ██║
██║██║ ╚████║██║   ██║
╚═╝╚═╝  ╚═══╝╚═╝   ╚═╝
*/

func init() {

	_path := flag.String("path", "", "The `path` param is changed in the"+
		" \nfile in the folder or folder under the folder or in the file"+
		" \ninside the folder."+
		" \nFor example `/home/mert.acel/DEVELOPMENT/GIT/AOSP`\n")

	_sourceName := flag.String("source", "", "Sourcename specifies the parameter to be"+
		" \nsearched and changed on the specified path.\n")

	_targetName := flag.String("target", "", "targetName is the name to be changed"+
		" \nin the specified path and the name being searched.\n")

	_inFile := flag.Bool("file", false, "If the file parameter is true, it searches for and modifies data in files."+
		" \nAttention ! dir and file should not be the same and true or false.\n")

	_inDir := flag.Bool("dir", false, "If the parameter is true, the folder names and filenames"+
		" \nin the folder are searched and changed."+
		" \nAttention ! dir and file should not be the same and true or false.\n")

	flag.Parse()
	fmt.Println("Get Param : ")
	fmt.Println("Path : ", *_path)
	fmt.Println("Source Name : ", *_sourceName)
	fmt.Println("Target Name : ", *_targetName)
	fmt.Println("Dir: ", *_inDir)
	fmt.Println("File : ", *_inFile)

	if *_path == "" || *_targetName == "" || *_inDir == *_inFile || len(os.Args) < 2 {
		defaultError()
	}

	path = *_path
	sourceName = *_sourceName
	targetName = *_targetName
	inFile = *_inFile
	inDir = *_inDir

}

/*
███╗   ███╗ █████╗ ██╗███╗   ██╗
████╗ ████║██╔══██╗██║████╗  ██║
██╔████╔██║███████║██║██╔██╗ ██║
██║╚██╔╝██║██╔══██║██║██║╚██╗██║
██║ ╚═╝ ██║██║  ██║██║██║ ╚████║
╚═╝     ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝
*/

func main() {
	var r ReplaceParam

	r.MainPath = path
	r.SourceName = sourceName
	r.TargetName = targetName
	r.InFile = inFile
	r.InDir = inDir

	r.WalkinPath()
}

/*
██╗    ██╗ █████╗ ██╗     ██╗  ██╗██╗███╗   ██╗        ██████╗  █████╗ ████████╗██╗  ██╗
██║    ██║██╔══██╗██║     ██║ ██╔╝██║████╗  ██║        ██╔══██╗██╔══██╗╚══██╔══╝██║  ██║
██║ █╗ ██║███████║██║     █████╔╝ ██║██╔██╗ ██║        ██████╔╝███████║   ██║   ███████║
██║███╗██║██╔══██║██║     ██╔═██╗ ██║██║╚██╗██║        ██╔═══╝ ██╔══██║   ██║   ██╔══██║
╚███╔███╔╝██║  ██║███████╗██║  ██╗██║██║ ╚████║        ██║     ██║  ██║   ██║   ██║  ██║
 ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝        ╚═╝     ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝
*/

//WalkinPath is function to navigate through the folder
func (r ReplaceParam) WalkinPath() {
	var folderSize int
	targetDir := r.MainPath

	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		beginFolder, lastFolder := filepath.Split(path)
		if lastFolder != ".git" {
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

			if !info.IsDir() {
				if r.InFile {
					r.replaceInFile(path)
				}
			}
		}
		return nil
	})
	errorHandler(err)

	fmt.Println("Folder Size : ", folderSize)
	for i := 0; i < folderSize; i++ {
		errFolder := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return nil
			}

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
		errorHandler(errFolder)
	}
}

/*
██████╗ ███████╗██████╗ ██╗      █████╗  ██████╗███████╗        ███████╗██╗██╗     ███████╗
██╔══██╗██╔════╝██╔══██╗██║     ██╔══██╗██╔════╝██╔════╝        ██╔════╝██║██║     ██╔════╝
██████╔╝█████╗  ██████╔╝██║     ███████║██║     █████╗          █████╗  ██║██║     █████╗
██╔══██╗██╔══╝  ██╔═══╝ ██║     ██╔══██║██║     ██╔══╝          ██╔══╝  ██║██║     ██╔══╝
██║  ██║███████╗██║     ███████╗██║  ██║╚██████╗███████╗        ██║     ██║███████╗███████╗
╚═╝  ╚═╝╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝ ╚═════╝╚══════╝        ╚═╝     ╚═╝╚══════╝╚══════╝
*/

//replaceFile is change specified parameter in a file
func (r ReplaceParam) replaceInFile(path string) {
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

/*
██████╗ ███████╗███████╗ █████╗ ██╗   ██╗██╗  ████████╗        ███████╗██████╗ ██████╗  ██████╗ ██████╗
██╔══██╗██╔════╝██╔════╝██╔══██╗██║   ██║██║  ╚══██╔══╝        ██╔════╝██╔══██╗██╔══██╗██╔═══██╗██╔══██╗
██║  ██║█████╗  █████╗  ███████║██║   ██║██║     ██║           █████╗  ██████╔╝██████╔╝██║   ██║██████╔╝
██║  ██║██╔══╝  ██╔══╝  ██╔══██║██║   ██║██║     ██║           ██╔══╝  ██╔══██╗██╔══██╗██║   ██║██╔══██╗
██████╔╝███████╗██║     ██║  ██║╚██████╔╝███████╗██║           ███████╗██║  ██║██║  ██║╚██████╔╝██║  ██║
╚═════╝ ╚══════╝╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝           ╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝
*/
// defaultError is contains the data to be written to the screen in case of an error
func defaultError() {
	fmt.Printf("Please Control Param\n\n")
	fmt.Printf("For Help -help \n\n")
	fmt.Printf("Example : \n\n")
	fmt.Printf("./FinderAndChanger -path/your/find/path -source=source_name  -target=target_name -file=in_file -dir=in_directory\n")
	fmt.Printf("./FinderAndChanger -path=/home/mert.acel/DEVELOPMENT/GIT/AOS -source=Test -target=Acel -file=true -dir=false\n\n")

	fmt.Printf("Params : \n\n")
	flag.PrintDefaults()
	os.Exit(1)
}

/*
███████╗██████╗ ██████╗  ██████╗ ██████╗         ██╗  ██╗ █████╗ ███╗   ██╗██████╗ ██╗     ███████╗██████╗
██╔════╝██╔══██╗██╔══██╗██╔═══██╗██╔══██╗        ██║  ██║██╔══██╗████╗  ██║██╔══██╗██║     ██╔════╝██╔══██╗
█████╗  ██████╔╝██████╔╝██║   ██║██████╔╝        ███████║███████║██╔██╗ ██║██║  ██║██║     █████╗  ██████╔╝
██╔══╝  ██╔══██╗██╔══██╗██║   ██║██╔══██╗        ██╔══██║██╔══██║██║╚██╗██║██║  ██║██║     ██╔══╝  ██╔══██╗
███████╗██║  ██║██║  ██║╚██████╔╝██║  ██║        ██║  ██║██║  ██║██║ ╚████║██████╔╝███████╗███████╗██║  ██║
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝        ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝
*/
// errorHandler is the operation to be performed when an error occurs
func errorHandler(err error) {
	if err != nil {
		defaultError()
		os.Exit(3)
	}
}
