package main

import (
	// "os"

	"io"

	"github.com/hajimehoshi/go-mp3"
)

const sampleSize = 4

func getLength(f io.Reader) int64 {
	decoder, err := mp3.NewDecoder(f)

	if err != nil {
		panic(err)
	}

	samples := decoder.Length() / sampleSize
	return samples / int64(decoder.SampleRate()) * 1000
}