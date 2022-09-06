package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func count(filename string) int {
	cmd := exec.Command("./xq.sh", filename)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	output := strings.TrimSuffix(out.String(), "\n")
	value, err := strconv.Atoi(output)
	if err != nil {
		fmt.Printf("err: %e\n", err)
	}
	return value
}

func processDir(files chan string) {
	path := "./lawsuit/"
	items, _ := ioutil.ReadDir(path)
	total := 0
	for _, item := range items {
		if !item.IsDir() {
			files <- path + item.Name()
			total += 1
			if total >= 100000 {
				return
			}
		}
	}
}

func readXml(files chan string, results chan int) {
	for filename := range files {
		result := count(filename)
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
		go readXml(files, results)
	}
	sum(results, numFiles)
}
