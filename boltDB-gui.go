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

func getAllBuckets(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /writeData")
	
	 err = db.View(func(tx *bolt.Tx) error {
        return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
            fmt.Println(string(name))
            return nil
        })
    })
    if err != nil {
        fmt.Println(err)
        return
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
