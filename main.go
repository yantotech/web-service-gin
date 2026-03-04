package main

/*
package main berfungsi untuk,
1. menandakan sebagai program utama
2. agar program pada file ini bisa dieksekusi dengan menggunakan go run
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
import berfungsi untuk
1. mengambil package lain diluar file itu sendiri
2. struktur import :
   import (""package yg mau diambil"")

net/http merupakan
1. package bawaan Go agar bisa menggunakan HTTP

github.com/gin-gonic/gin merupakan
1. framework Gin yang web framework berkinerja tinggi yang dirancang untuk bahasa pemrograman Go (Golang),
   yang utamanya digunakan untuk membangun API RESTful, layanan mikro (microservices),
   dan aplikasi web yang cepat dan ringan.
*/

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

/*
type merupakan
1. type merupakan keyword untuk membuat tipe data baru
2. struktur type
   type "nama type" struct {...}
3. golang memiliki jenis type selayaknya pemrograman biasa,
   namun di Golang mendukung pembuatan type data (user-defined type)

ID  string  `json:"id"
1. ID merupakan nama field dengan tipe data string
2. `json "id"` merupakan struct tag yang berfungsi untuk mengatur nama field saat dikonversi ke JSON(dapat dilihatpada haisl konversi)
*/

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

/*
1. var merupakan keyword untuk mendeklarasikan variabel
   struktur umum var:
   var "nama variabel" "tipe data" = nilai
2. slice merupakan tipe data di Go yang digunakan untuk menyimpan kumpulan data (seperti array),
   tetapi ukurannya bisa berubah-ubah (dinamis).
   contoh penulisan:
   []"nama variabel"
*/

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

/*
1. func merupakan keyword untuk membaut function
2. struktur function
   func "nama function" ("nama parameter" "tipe parameter") "tipe return" {
	  "instruksi"
   }
3. main() menjadikan function ini merupakan function yang pertama dijalankan saat program dieksekusi
4. router := gin.Default() untuk membuat engine/router yang berfungsi sebagai mesin utama dari server
5. router.GET("/albums", getAlbums) untuk membuat endpoint HTTP GET dengan ketentuan
   a) jika ada request GET ke URL /albums, maka akan menjalankan function getAlbums
6. router.Run("localhost:8080") berfungsi menjalankan server, sehingga bisa diakses lewat browser
*/

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

/*
1. getAlbums merupakan fungsi yang digunakan untuk menangani HTTP request dan mengirimkan response dalam format JSON
2. c *gin.Context merupakan parameter berjenis pointer yang mengarah ke gin.Context
3. c.IndentedJSON(http.StatusOK, albums) bagian ini berfungsi untuk mengirimkan response ke client
   a. http.StatusOK untuk mengirimkan kode status HTTP 200 OK
   b. albums merupakan jenis data yang akan dikirimkan
   c. IndentedJSON untuk mengirimkan response dalam bentuk JSON dengan indentasi
*/

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/*
1. var newAlbum album adalah kode untuk membuat variabel baru bernama newAlbum dengan tipe album
2. c.BindJSON(&newAlbum) berfungsi untuk membaca JSON dari request body dan mengubah JSON tersebut menjadi struct pada newAlbum
3. albums = append(albums, newAlbum) berfungsi untuk menambahkan newAlbum kedalam slice albums
*/

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

/*
1. id := c.Param("id") berfungsi untuk mengambil parameter id dari URL
2. for _, a := range albums berfungsi untuk melakukan perulangan terhadap semua data di slice albums.
3. if a.ID == id berfungsi untuk mengecek apakah ID album cocok dengan ID dari URL
4.
*/
