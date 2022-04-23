package main

import (
	"fmt"
	"log"
)
import "github.com/ahmed9905/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	fileNew, err := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileNew)

	fmt.Println("it uploaded", file)
}
