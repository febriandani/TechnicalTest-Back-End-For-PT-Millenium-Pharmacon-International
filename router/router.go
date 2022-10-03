package router

import (
	"fmt"
	"html/template"
	"net/http"
	"technical-test-02-10-22/controller"
	"technical-test-02-10-22/database"
)

func SetupRoutes() {

	db := database.GetDB()

	fileRepository := controller.NewRepository(db)

	fileService := controller.NewService(fileRepository)

	fileHandler := controller.NewFileHandler(fileService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		upload := fileHandler.UploadFile

		var t, err = template.ParseFiles("web/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		t.Execute(w, upload)
	})

	http.HandleFunc("/upload", fileHandler.UploadFile)
	http.ListenAndServe(":8080", nil)
}
