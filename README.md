# FinderAndChanger
This application is able to change the name and content in the given parameter.

Development :
```bash
go get -u github.com/KMACEL/FinderAndChanger
```
Download :
```bash
https://github.com/KMACEL/FinderAndChanger/releases
```

Example :
```bash
./FinderAndChanger -path=/home/mert.acel/DEVELOPMENT/GIT/AOS -source=Test -target=Acel -file=true -dir=false
```

Params :
```text
  -dir bool
    	If the parameter is true, the folder names and filenames
    	in the folder are searched and changed.
    	Attention ! dir and file should not be the same and true or false.

  -file bool
    	If the file parameter is true, it searches for and modifies data in files.
    	Attention ! dir and file should not be the same and true or false.

  -path path
    	The path param is changed in the
    	file in the folder or folder under the folder or in the file
    	inside the folder.
    	For example `/home/mert.acel/DEVELOPMENT/GIT/AOSP`

  -source string
    	Sourcename specifies the parameter to be
    	searched and changed on the specified path.

  -target string
    	targetName is the name to be changed
    	in the specified path and the name being searched.
```
