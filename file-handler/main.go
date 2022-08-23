package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Opens a file for read only
	f1, err := os.Open("test1.txt")
	PrintFatalError(err)
	defer f1.Close() // Close resource

	// Create a new file
	f2, err := os.Create("test2.txt")
	PrintFatalError(err)
	defer f2.Close()

	// Open file for read write
	f3, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_RDWR, 0666) // We can OR multiple os operation on file
	PrintFatalError(err)
	defer f3.Close()

	// err = os.Rename("test3.txt", "test1New.txt")
	// PrintFatalError(err)

	// move a file
	err = os.Rename("./test4move.txt", "./testfolder/test.txt")
	PrintFatalError(err)

	// err = os.Remove("test1New.txt")
	// PrintFatalError(err)

	// tempFile, err := os.Create("test4.txt")
	// PrintFatalError(err)
	// defer tempFile.Close()

	CopyFile("test1.txt", "test4.txt")

	// read a file
	bytes, _ := ioutil.ReadFile("test1.txt")
	fmt.Println(string(bytes))

	scanner := bufio.NewScanner(f3)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println("Found line: ", count, scanner.Text())
	}

	// buffered writer, content stored in memory and saves disk I/O
	writerBuffer := bufio.NewWriter(f3)
	for i := 1; i <= 5; i++ {
		writerBuffer.WriteString(fmt.Sprintln("added line", i))
	}
	writerBuffer.Flush()

	GenerateFileStatusReport("test1.txt")
}

func CopyFile(fname1, fname2 string) {
	fOld, err := os.Open(fname1)
	PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	// copy bytes from source to destinataion
	_, err = io.Copy(fNew, fOld)
	PrintFatalError(err)

	// flush file contents to disc
	err = fNew.Sync()
	PrintFatalError(err)
}

func GenerateFileStatusReport(fname string) {
	filestats, err := os.Stat(fname)
	PrintFatalError(err)

	fmt.Println("File Name: ", filestats.Name())
	fmt.Println("Is Dir? ", filestats.IsDir())
	fmt.Println("Permissions: ", filestats.Mode())
	fmt.Println("File Size: ", filestats.Size())
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error while processing file: ", err)
	}
}
