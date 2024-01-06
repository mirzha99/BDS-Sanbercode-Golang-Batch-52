package main

import (
	"net/http"
	"tugas-15/controller/mahasiswa"
	"tugas-15/controller/matakuliah"
	"tugas-15/controller/nilai"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//mahasiswa
	router.GET("/mahasiswa", mahasiswa.GetMahasiswa)
	router.POST("/mahasiswa", mahasiswa.PostMahasiswa)
	router.PUT("/mahasiswa/:id", mahasiswa.PutMahasiswa)
	router.DELETE("/mahasiswa/:id", mahasiswa.DeleteMahasiswa)

	//mata kuliah
	router.GET("/matakuliah", matakuliah.GetMatakuliah)
	router.POST("/matakuliah", matakuliah.PostMatakuliah)
	router.PUT("/matakuliah/:id", matakuliah.PutMatakuliah)
	router.DELETE("/matakuliah/:id", matakuliah.DeleteMatakuliah)

	//nilai
	router.GET("/nilai", nilai.GetNilai)
	router.GET("/nilai/:id", nilai.GetNilaiById)
	router.POST("/nilai", nilai.PostNilai)
	router.PUT("/nilai/:id", nilai.PutNilai)
	router.DELETE("/nilai/:id", nilai.DeleteNilai)

	http.ListenAndServe(":8080", router)

}
