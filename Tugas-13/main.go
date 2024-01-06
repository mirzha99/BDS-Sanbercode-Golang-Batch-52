package main

import (
	"net/http"
	"tugas-13/controller"
	"tugas-13/middleware"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.Handle("/nilaimahasiswa", middleware.Auth(http.HandlerFunc(controller.NilaiMahasiswa)))
	http.ListenAndServe(":8080", nil)
}
