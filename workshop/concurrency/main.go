package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	frequency := make([]int64, 26)
	startIndex := 1
	endIndex := 1500
	for i := startIndex; i <= endIndex; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &wg, &mutex)
	}
	wg.Wait()
	for index, value := range allLetters {
		fmt.Printf("%c=%d\n", value, frequency[index])
	}
}

func countLetters(url string, frequency []int64, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		panic("error while getting the data from " + url + " with an error = " + err.Error())
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		for _, v := range body {
			c := strings.ToLower(string(v))
			cIndex := strings.Index(allLetters, c)
			if cIndex >= 0 {
				mutex.Lock()
				frequency[cIndex] += 1
				mutex.Unlock()
			}
		}
		fmt.Println("Completed : ", url)
	}
}
