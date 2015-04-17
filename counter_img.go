package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
)

const (
	IMAGE_PATH = "images/numbers.png"
)

var spriteMap map[rune]image.Image

func init() {
	spriteMap = make(map[rune]image.Image)
	buildSpriteMap()
}

func GetCountPng(count uint64) image.Image {

	number := strconv.FormatUint(count, 10)

	counter := make([]image.Image, len(number))
	for i, c := range number {
		counter[i] = spriteMap[c]
	}

	return buildImage(counter)
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

// Image is 300px x 400px
// Each character is 100px x 100px
func buildSpriteMap() {
	fin, err := os.Open(IMAGE_PATH)
	if err != nil {
		log.Fatalf("Error opening base image: %s\n", err)
	}
	defer fin.Close()

	sourceImg, _, err := image.Decode(fin)
	if err != nil {
		log.Fatalf("Error decoding base image: %s\n", err)
	}

	one := image.NewRGBA(image.Rect(0, 0, 100, 100))
	two := image.NewRGBA(image.Rect(0, 0, 100, 100))
	three := image.NewRGBA(image.Rect(0, 0, 100, 100))
	four := image.NewRGBA(image.Rect(0, 0, 100, 100))
	five := image.NewRGBA(image.Rect(0, 0, 100, 100))
	six := image.NewRGBA(image.Rect(0, 0, 100, 100))
	seven := image.NewRGBA(image.Rect(0, 0, 100, 100))
	eight := image.NewRGBA(image.Rect(0, 0, 100, 100))
	nine := image.NewRGBA(image.Rect(0, 0, 100, 100))
	comma := image.NewRGBA(image.Rect(0, 0, 100, 100))
	zero := image.NewRGBA(image.Rect(0, 0, 100, 100))
	period := image.NewRGBA(image.Rect(0, 0, 100, 100))

	draw.Draw(one, one.Bounds(), sourceImg, image.Point{0, 0}, draw.Src)
	spriteMap['1'] = one

	draw.Draw(two, two.Bounds(), sourceImg, image.Point{100, 0}, draw.Src)
	spriteMap['2'] = two

	draw.Draw(three, three.Bounds(), sourceImg, image.Point{200, 0}, draw.Src)
	spriteMap['3'] = three

	draw.Draw(four, four.Bounds(), sourceImg, image.Point{0, 100}, draw.Src)
	spriteMap['4'] = four

	draw.Draw(five, five.Bounds(), sourceImg, image.Point{100, 100}, draw.Src)
	spriteMap['5'] = five

	draw.Draw(six, six.Bounds(), sourceImg, image.Point{200, 100}, draw.Src)
	spriteMap['6'] = six

	draw.Draw(seven, seven.Bounds(), sourceImg, image.Point{0, 200}, draw.Src)
	spriteMap['7'] = seven

	draw.Draw(eight, eight.Bounds(), sourceImg, image.Point{100, 200}, draw.Src)
	spriteMap['8'] = eight

	draw.Draw(nine, nine.Bounds(), sourceImg, image.Point{200, 200}, draw.Src)
	spriteMap['9'] = nine

	draw.Draw(comma, comma.Bounds(), sourceImg, image.Point{0, 300}, draw.Src)
	spriteMap[','] = comma

	draw.Draw(zero, zero.Bounds(), sourceImg, image.Point{100, 300}, draw.Src)
	spriteMap['0'] = zero

	draw.Draw(comma, period.Bounds(), sourceImg, image.Point{100, 300}, draw.Src)
	spriteMap['.'] = period
}
