package main

import (
	"encoding/json"
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

	// List of buckets
	var BucketList []string

	// Create transaction
	err = db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(bucketName []byte, _ *bolt.Bucket) error {
			BucketList = append(BucketList, string(bucketName))
			return nil
		})
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Jsonize bucket list
	jsonData, err := json.Marshal(BucketList)
	if err != nil {
		log.Println(err)
	}

	// Response variable
	var response APIResponse
	response.Body = string(jsonData)

	// Jsonize response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	// Sending response
	if _, err = w.Write(jsonResponse); err != nil {
		log.Println(err)
	}
}

// Function to write data into bucket
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

// Function to get data from a bucket
func getAllData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /getAllData")

	// Bucket name variable
	var bucketName string
	// Get bucket name from parameters
	keys := r.URL.Query()["bucket"]
	bucketName = keys[0]

	// Empty bucket content
	var bucketContent []Data

	// Open transaction
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))
		// Iterate over key-value pair
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			bucketContent = append(bucketContent, Data{Key: string(k), Value: string(v)})
			return nil
		})
		return nil
	})

	// Marshal bucket content
	jsonData, _ := json.Marshal(bucketContent)

	// Build response
	var response APIResponse
	response.Body = string(jsonData)

	// Marshal response
	responseJSON, _ := json.Marshal(response)

	// Return response
	w.Write(responseJSON)
}

// Function to create a new bucket
func createBucket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /createBucket")

	// Bucket name variable
	var bucketName string
	// Get bucket name from request
	keys := r.URL.Query()["bucket"]
	bucketName = keys[0]
	fmt.Println(keys, bucketName)

	// Open transaction
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Response variable
	var response APIResponse
	response.Body = "SUCCESSFUL"

	// Jsonize response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	// Sending response
	if _, err = w.Write(jsonResponse); err != nil {
		log.Println(err)
	}
}

// Function to delete a bucket
func deleteBucket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /deleteBucket")

	// Bucket name variable
	var bucketName string
	// Get bucket name from request
	keys := r.URL.Query()["bucket"]
	bucketName = keys[0]
	fmt.Println(keys, bucketName)

	// Open transaction
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			fmt.Println("delete bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ppppp")
	// Response variable
	var response APIResponse
	response.Body = "SUCCESSFUL"

	// Jsonize response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	// Sending response
	if _, err = w.Write(jsonResponse); err != nil {
		log.Println(err)
	}
}

// Function to get stats
func getStats(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Page hit : /getStats")
}
