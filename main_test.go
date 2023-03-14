package main

import (
	"log"
	"testing"
	"time"
)

type testArgs struct {
	image   string
	scale   int
	colors  int
	kmtimes int
}

func printArgs(args *testArgs, elapsed time.Duration) {
	log.Println(bold, "\nRunning test:\n  - image:", (*args).image, "\n  - scale:", (*args).scale, "\n  - colors:", (*args).colors, "\n  - km times:", (*args).kmtimes, reset)
	log.Println(bold, blink, green, "\nTest took", elapsed, reset)
}

func runFSDTest(args *testArgs) time.Duration {

	img, err := openImage((*args).image)
	if err != nil {
		log.Println((*args))
		log.Fatal("Couldn't get image opened")
	}

	pixels := imageToPixels(img)

	X := len(*pixels)
	Y := len((*pixels)[0])

	downscaleNoUpscale(pixels, (*args).scale)

	palettes := getPalettesFromJSON("colorpalette.json")
	palette := getPaletteWithName("GameBoy", palettes)

	start := time.Now()

	floydSteinbergDithering(pixels, palette, scale, Y, X)

	return time.Since(start)
}

func runKMTest(args *testArgs) time.Duration {
	img, err := openImage((*args).image)
	if err != nil {
		log.Println((*args))
		log.Fatal("Couldn't get image opened")
	}

	pixels := imageToPixels(img)

	downscaleNoUpscale(pixels, (*args).scale)

	start := time.Now()

	createColorPalette(pixels, *amountOfColors, 4, (*args).kmtimes)

	return time.Since(start)
}

func TestFSDSpeed(t *testing.T) {
	log.Println(bold, "\nStarting FSD Speed Test.\n", reset)

	args := testArgs{
		"src/jordgubbar.jpg",
		1,
		10,
		1,
	}
	printArgs(&args, runFSDTest(&args))
}

func TestKMSpeed(t *testing.T) {
	log.Println(bold, "\nStarting KMeans Speed Test.\n", reset)

	args := testArgs{
		"src/jordgubbar.jpg",
		15,
		5,
		20,
	}

	printArgs(&args, runKMTest(&args))
}
