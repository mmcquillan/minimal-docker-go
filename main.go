package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting...")
	for _, env := range os.Environ() {
		fmt.Println(" - ENV: " + env)
	}
	url := "https://www.google.com/"
	fmt.Println(" - URL: " + url)
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	loc := os.TempDir() + "/out"
	fmt.Println(" - TMP: " + loc)
	ioutil.WriteFile(loc, bytes, 0644)
}
