package main

import (
	"os"
	"sync"
)

// /Users/nikolaistepanov/Desktop/test_video.mp4

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_chunk(file *os.File, i int, wg *sync.WaitGroup) {
	// нужно дописать функцию, чтобы она читала трафик
}

const BufferSize = 10485760

func main() {
	file, err := os.Open("/Users/nikolaistepanov/Desktop/test_video.mp4")
	check(err)

	defer file.Close()
	file_stat, err := file.Stat()
	check(err)
	file_size := int(file_stat.Size())
	concurrency := file_size / BufferSize

	if check := file_size % BufferSize; check != 0 {
		concurrency++
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {

	}
}
