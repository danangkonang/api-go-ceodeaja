package library

import (
	"fmt"
	"os"
)

var path = "h9h2fhfUVuS9jZ8uVbhV3vC5AWX39IVU.png"

func SendEmail() {
	deleteFile()
}

func deleteFile() {
	var err = os.Remove("files/" + path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
