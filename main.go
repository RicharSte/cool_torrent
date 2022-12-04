package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

// /Users/nikolaistepanov/Desktop/test_video.mp4

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write_down(data []byte, wg *sync.WaitGroup) {
	hasher := sha256.New() // fix later
	hasher.Write(data)
	file, err := os.Create(hex.EncodeToString(hasher.Sum(nil)))
	check(err)
	defer file.Close()

	file.Write(data)
	defer wg.Done()
}

func read_chunk(file *os.File, offset, size, i int, ch chan []byte, ch1 chan string) {
	buffer := make([]byte, size)
	_, err := file.ReadAt(buffer, int64(offset))
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	hasher := sha256.New() //fix later
	hasher.Write(buffer)
	ch <- buffer
	ch1 <- strconv.Itoa(i) + "__" + hex.EncodeToString(hasher.Sum(nil))
}

const BufferSize = 10485760

func main() {

	var wg sync.WaitGroup
	ch_byte := make(chan []byte)
	ch_hash := make(chan string)
	var block = make(map[string]int)

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
	wg.Add(concurrency * 2)

	for i := 1; i <= concurrency; i++ {
		if i == 1 {
			start := 0
			go read_chunk(file, start, BufferSize, i, ch_byte, ch_hash)
		} else {
			start := (BufferSize * (i - 1)) + 1
			go read_chunk(file, start, BufferSize, i, ch_byte, ch_hash)
		}
	}
	go func() {
		for c := range ch_byte {
			write_down(c, &wg)
		}
	}()
	go func() {
		for c := range ch_hash {
			splitted_string := strings.Split(c, "__")
			int_val, _ := strconv.Atoi(splitted_string[0])
			block[splitted_string[1]] = int_val
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println(block)
}
