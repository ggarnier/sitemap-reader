package main

import (
	"fmt"
	"io/ioutil"
)

func processDir(files chan string) {
	path := "./lawsuit/"
	items, _ := ioutil.ReadDir(path)
	for _, item := range items {
		if !item.IsDir() {
			files <- path + item.Name()
		}
	}
}

func readXmlFiles(files chan string, results chan int) {
	for filename := range files {
		result := parseXml(filename)
		results <- result
	}
}

func sum(results chan int, numFiles int) {
	total := 0
	for i := 0; i < numFiles; i++ {
		val := <-results
		total += val
		if i%20 == 0 {
			fmt.Printf("foi %d\n", i)
		}
	}
	fmt.Printf("total = %d\n", total)
}

func main() {
	files := make(chan string, 100000)
	parallel := 40
	results := make(chan int, parallel)

	processDir(files)
	numFiles := len(files)
	for i := 0; i < parallel; i++ {
		go readXmlFiles(files, results)
	}
	sum(results, numFiles)
}
