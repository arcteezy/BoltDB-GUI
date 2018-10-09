package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

// Data structure
type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// APIResponse structure
type APIResponse struct {
	Body string `json:"body"`
}

// Bolt DB
var db *bolt.DB

// Error
var err error

func main() {

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err = bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Starting web server
	http.HandleFunc("/writeData", writeData)
	http.HandleFunc("/getAllData", getAllData)
	http.HandleFunc("/getAllBuckets", getAllBuckets)
	http.HandleFunc("/createBucket", createBucket)
	http.HandleFunc("/deleteBucket", deleteBucket)
	http.HandleFunc("/getStats", getStats)
	if err = http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}

}

// Fucntion to get all buckets
func getAllBuckets(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /getAllBuckets")

	// Create transaction
	err = db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(bucketName []byte, _ *bolt.Bucket) error {
			fmt.Println(string(bucketName))
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

	var keys []string
	var key, value, bucket string

	keys = r.URL.Query()["key"]
	key = keys[0]
	fmt.Println(keys)

	keys = r.URL.Query()["value"]
	value = keys[0]
	fmt.Println(keys)

	keys = r.URL.Query()["bucket"]
	bucket = keys[0]
	fmt.Println(keys)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func getAllData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /getAllData")

	var bucketName string
	keys := r.URL.Query()["bucket"]
	bucketName = keys[0]
	fmt.Println(keys, bucketName)

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
}

func createBucket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /createBucket")

	var bucketName string
	keys := r.URL.Query()["bucket"]
	bucketName = keys[0]
	fmt.Println(keys, bucketName)

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func deleteBucket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /deleteBucket")

	var bucketName string
	if keys := r.URL.Query()["bucket"]; len(keys) > 1 {
		bucketName = keys[0]
	}

	db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func getStats(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /getStats")
}
