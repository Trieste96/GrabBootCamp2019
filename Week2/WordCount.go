package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

//SafeCounter is safe when being accessed concurrently
type SafeCounter struct {
	occ map[string]int
	mux sync.Mutex
}

//IncCounter get a word from the channel and increases its occurence by 1
func (sc *SafeCounter) IncCounter(c chan string) {
	for {
		w := <-c
		sc.mux.Lock()
		sc.occ[w]++
		sc.mux.Unlock()
	}
}

//Get all words in a file and insert to the channel
func countWord(filePath string, c chan string, wg *sync.WaitGroup) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}

	//fmt.Printf("Content %40s: %s\n", filePath, content)
	words := strings.Fields(string(content))
	for _, w := range words {
		c <- w
	}
	wg.Done()
}

//ReadDir recursively returns all file names in the current directory and in its sub-directories
func ReadDir(dir string) (filePath []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			subFiles, err := ReadDir(dir + "/" + f.Name())
			if err != nil {
				return nil, err
			}
			filePath = append(filePath, subFiles...)
		} else {
			filePath = append(filePath, dir+"/"+f.Name())
		}
	}
	return filePath, nil
}

func printOccurence(m map[string]int) {
	fmt.Printf("Occurence: ")
	for k, v := range m {
		fmt.Printf("%s=%d, ", k, v)
	}
	fmt.Println("")
}

//Create one gorountine for each file to push their words to the channel
func countWordAllFiles(dir string) {
	files, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan string)
	counter := SafeCounter{occ: make(map[string]int)}
	go counter.IncCounter(ch)

	var wg sync.WaitGroup
	for _, f := range files {
		fmt.Printf("File %s\n", f)
		wg.Add(1)
		go countWord(f, ch, &wg)
	}
	wg.Wait()
	printOccurence(counter.occ)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing directory path, go run [GO_FILE_NAME] [DIR_PATH] ")
	}

	dir := os.Args[1]
	countWordAllFiles(dir)
}
