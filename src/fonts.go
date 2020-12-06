package main

import (
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// LoadTTF loads a ttf font from source
// Code from https://github.com/faiface/pixel/wiki/Typing-text-on-the-screen
func LoadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
