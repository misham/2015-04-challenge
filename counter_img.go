package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	IMAGE_MAX_X_PX  = 300
	IMAGE_X_STEP_PX = 100
	IMAGE_Y_STEP_PX = 100
	IMAGE_MAX_Y_PX  = 400
	IMAGE_PATH      = "images/numbers.png"
)

var spriteMap map[rune]image.Image

func init() {
	spriteMap = make(map[rune]image.Image)
	err := buildSpriteMap()
	if err != nil {
		log.Fatalf("Error initializing image library: %s\n", err)
	}
}

func GetCountPng(count int64) ([]byte, error) {

	number := humanize.Comma(count)

	counter := make([]image.Image, len(number))
	for i, c := range number {
		counter[i] = spriteMap[c]
	}

	img := buildImage(counter)
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)

	if err != nil {
		log.Printf("Error converting image to bytes: %s\n", err)
		return nil, err
	}

	return buf.Bytes(), nil
}

func buildImage(counter []image.Image) image.Image {
	// Create temp image size X: len(counter), Y: 100px
	//   Each rune is 100px by 100px. So the length is the number of runes
	//   times 100px, but the height is always 100 runes.
	// TODO Add support for comma
	result := image.NewRGBA(image.Rect(0, 0, len(counter)*100, 100))

	for i, c := range counter {

		rect := image.Rect(i*100, 0, (i+1)*100, 100)

		draw.Draw(result, rect, c, image.Point{0, 0}, draw.Src)
	}

	return result
}

// Image is 300px x 400px
// Each character is 100px x 100px
func buildSpriteMap() error {
	fin, err := os.Open(IMAGE_PATH)
	if err != nil {
		log.Fatalf("Error opening base image: %s\n", err)
	}
	defer fin.Close()

	sourceImg, _, err := image.Decode(fin)
	if err != nil {
		log.Fatalf("Error decoding base image: %s\n", err)
	}

	for x := 0; x < IMAGE_MAX_X_PX; x = x + IMAGE_X_STEP_PX {
		for y := 0; y < IMAGE_MAX_Y_PX; y = y + IMAGE_Y_STEP_PX {
			tmp := image.NewRGBA(image.Rect(0, 0, 100, 100))
			draw.Draw(tmp, tmp.Bounds(), sourceImg, image.Point{x, y}, draw.Src)

			index, err := getSpritIndex(x, y)
			if err != nil {
				return fmt.Errorf("Error getting sprite image: %s", err)
			}
			spriteMap[index] = tmp
		}
	}

	return nil
}

// Format of the image is:
//
//   | 1 | 2 | 3 |
//   | 4 | 5 | 6 |
//   | 7 | 8 | 9 |
//   | , | 0 | . |
//
// TODO Use constants to calculate the dimentions assuming the same
//      arrangement
func getSpritIndex(x int, y int) (rune, error) {
	var result rune
	var err error

	switch {
	case x == 0 && y == 0:
		result = '1'
	case x == 100 && y == 0:
		result = '2'
	case x == 200 && y == 0:
		result = '3'
	case x == 0 && y == 100:
		result = '4'
	case x == 100 && y == 100:
		result = '5'
	case x == 200 && y == 100:
		result = '6'
	case x == 0 && y == 200:
		result = '7'
	case x == 100 && y == 200:
		result = '8'
	case x == 200 && y == 200:
		result = '9'
	case x == 0 && y == 300:
		result = ','
	case x == 100 && y == 300:
		result = '0'
	case x == 200 && y == 300:
		result = '.'
	default:
		result = 0
		err = errors.New("Unkown coordinates")
	}

	return result, err
}

// For debugging only
func saveImage(name string, img image.Image) {
	fout, err := os.Create(name)
	if err != nil {
		log.Printf("Error creating image file: %s\n", err)
		return
	}

	err = png.Encode(fout, img)
	if err != nil {
		log.Printf("Error saving image: %s\n", err)
	}
}
