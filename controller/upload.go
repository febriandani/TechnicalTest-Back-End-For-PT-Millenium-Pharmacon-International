package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const MAX_SIZE_FILE = 8 * 1024 * 1024

type fileHandler struct {
	s Service
}

func NewFileHandler(s Service) *fileHandler {
	return &fileHandler{s}
}

func (h *fileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// db := database.GetDB()

	var input FileInput
	fmt.Println("File Upload Endpoint Hit")

	// upload of 8 MB files.
	r.ParseMultipartForm(MAX_SIZE_FILE)
	// FormFile returns the first file for the given key `data`
	// it also returns the FileHeader so we can get the Filename,
	file, handler, err := r.FormFile("data")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	//capture file name, size, content type and save to database
	input.NameFile = handler.Filename
	input.Size = int(handler.Size)
	input.ContentType = handler.Header.Get("Content-Type")

	_, err = h.s.Create(input)
	if err != nil {
		panic(err)
	}

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
