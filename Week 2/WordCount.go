package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

//Safe counter map is safe when using concurrently
type SafeCounter struct {
	occ map[string]int
	mux sync.Mutex
}

func (sc *SafeCounter) updateWordMap(c chan string) {
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
	fmt.Printf("Content of file %s: %s\n", filePath, content)
	words := strings.Fields(string(content))
	for _, w := range words {
		c <- w
	}
	wg.Done()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing path of folder")
	}
	dir := os.Args[1]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	c := make(chan string, 100)
	var wg sync.WaitGroup

	sc := SafeCounter{occ: make(map[string]int)}
	go sc.updateWordMap(c)
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
