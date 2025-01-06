package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	frequency := make([]int64, 26)
	startIndex := 1
	endIndex := 20
	for i := startIndex; i <= endIndex; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency)
	}
	for index, value := range allLetters {
		fmt.Printf("%c=%d\n", value, frequency[index])
	}
}

func countLetters(url string, frequency []int64) {
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
				frequency[cIndex] += 1
			}
		}
		fmt.Println("Completed : ", url)
	}
}
