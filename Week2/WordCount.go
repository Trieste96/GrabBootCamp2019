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

//IncCounter increases the occurence of a word by 1
func (sc *SafeCounter) IncCounter(c chan string) {
	for {
		w := <-c

		sc.mux.Lock()
		sc.occ[w]++
		sc.mux.Unlock()
	}
}

func countWord(filePath string, c chan string, wg *sync.WaitGroup) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	words := strings.Fields(string(content))
	for _, w := range words {
		c <- w
	}
	wg.Done()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing directory path, go run [GO_FILE_NAME] [DIR_PATH] ")
	}
	dir := os.Args[1]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan string)
	var wg sync.WaitGroup

	sc := SafeCounter{occ: make(map[string]int)}
	go sc.IncCounter(c)

	for _, file := range files {
		filePath := dir + "/" + file.Name()
		wg.Add(1)
		go countWord(filePath, c, &wg)
	}
	wg.Wait()

	for k, v := range sc.occ {
		fmt.Printf("%s: %d\n", k, v)
	}

}
