package soal2

import (
	"fmt"
	"strconv"
)

func getMovies(ch chan string, movies ...string) {
	defer close(ch)
	fmt.Println("List Movies :")
	// Mengirim data movie ke channel
	for i, movie := range movies {
		ch <- strconv.Itoa(i+1) + ". " + movie
	}

}
func Soal2() {
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)
	for value := range moviesChannel {
		fmt.Println(value)
	}
}
