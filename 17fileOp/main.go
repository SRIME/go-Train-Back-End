// Key Points:

// 1. Create file using 							 = os.Create("./filename.txt")
// 2. Insert or append Value into file (Write) using = io.WriteString(fileVariableUsedToCreateFile, datatobeinserted)
// 3. Read the file									 = ioutil.ReadFile(fileVariable)
// 4. Closing the file using = defer				 = fileVariable.Close()

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("-----File Operations(Create, Write and Read)-----")

	//1. Creation of a file
	newFile, err := os.Create("./createFile.txt")
	checkNilError(err)
	textToBeInserted := "Hello, I'm being Inserted...."

	//2. Write to File
	lengthOfFile, err := io.WriteString(newFile, textToBeInserted)
	checkNilError(err)
	fmt.Println("Length of file is: ", lengthOfFile)

	//3. Read File
	readFile("./createFile.txt")

	//4. Closing the file at last
	defer newFile.Close()
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Data Bytes in the File: \n", data)
	fmt.Println("Contents of the File: \n", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
