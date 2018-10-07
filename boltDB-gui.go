package main

import (
	"fmt"
	"log"
	"net/http"
)

// Error
var err error

func main() {

	// Starting web server
	http.HandleFunc("/writeData", writeData)
	http.HandleFunc("/getAllData", getAllData)
	http.HandleFunc("/createBucket", createBucket)
	http.HandleFunc("/deleteBucket", deleteBucket)
	http.HandleFunc("/getStats", getStats)

	if err = http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}

}

func writeData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /writeData")
}

func getAllData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /getAllData")
}

func createBucket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /createBucket")
}

func deleteBucket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /deleteBucket")
}

func getStats(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /getStats")
}
