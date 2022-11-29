package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// /Users/nikolaistepanov/Desktop/test_video.mp4

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_chunk(file *os.File, offset, size, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer := make([]byte, size)
	file_chunk, err := file.ReadAt(buffer, int64(offset))
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	fmt.Println("bytes read, string(bytestream): ", file_chunk, "chunk nuber: ", i)
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

	for i := 1; i <= concurrency; i++ {
		if i == 1 {
			start := 0
			go read_chunk(file, start, BufferSize, i, &wg)
		} else {
			start := (BufferSize * (i - 1)) + 1
			go read_chunk(file, start, BufferSize, i, &wg)
		}
	}
	wg.Wait()
}
