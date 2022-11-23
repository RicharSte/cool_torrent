package main

import (
	"fmt"
	"os"
)

// /Users/nikolaistepanov/Desktop/test_video.mp4

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const BufferSize = 10485760

func main() {
	file, err := os.Open("/Users/nikolaistepanov/Desktop/test_video.mp4")
	check(err)

	defer file.Close()
	file_stat, err := file.Stat()
	check(err)
	fmt.Println(int(file_stat.Size()))
}
