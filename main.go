package main

import (
	"fmt"

	"github.com/t3hcyberrat/learning-ai/svm"
)

func main() {

	fmt.Println("Exec SVM...")
	err := svm.Exec()
	if err != nil {
		panic(err)
	}

	fmt.Println("SVM exec successful")

}
