package main

import (
	"net/http"
	"tugas-14/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nilaimahasiswa", controller.GetNilaiMahasiswa)
	router.POST("/nilaimahasiswa", controller.PostNilaiMahasiswa)
	router.PUT("/nilaimahasiswa/:id", controller.PutNilaiMahasiswa)
	router.DELETE("/nilaimahasiswa/:id", controller.DeleteNilaiMahasiswa)
	http.ListenAndServe(":8080", router)
}
