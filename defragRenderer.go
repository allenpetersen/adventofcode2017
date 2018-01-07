package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func renderDefrag(events <-chan diskEvent) {
	gridSize := 128
	locSize := 7
	var w, h int = locSize * 128, locSize * 128

	var palette = []color.Color{
		color.RGBA{0x2d, 0x2d, 0x2d, 0xff},
		color.RGBA{0x31, 0xd6, 0x89, 0xff},
		color.RGBA{0x32, 0x32, 0x32, 0xff},
		color.RGBA{0xff, 0xa8, 0xd9, 0xff},
	}

	var images []*image.Paletted
	var delays []int

	for e := range events {
		if e.kind == done || e.kind == groupComplete {
			continue
		}

		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				offsetX := x * locSize
				offsetY := y * locSize
				for i := 0; i < locSize; i++ {
					img.Set(offsetX+i, offsetY, palette[0])
					img.Set(offsetX+i, offsetY+locSize, palette[0])
				}
				for i := 0; i < locSize-2; i++ {
					img.Set(offsetX, offsetY, palette[0])
					img.Set(offsetX+locSize, offsetY+i+1, palette[0])
				}
				var c color.Color
				if e.disk.grid[x][y] {
					c = palette[1]
				} else {
					c = palette[2]
				}
				drawCell(img, c, offsetX, offsetY, locSize)
			}
		}

		for _, l := range e.disk.group {
			offsetX := l.x * locSize
			offsetY := l.y * locSize
			drawCell(img, palette[3], offsetX, offsetY, locSize)
		}
	}

	fmt.Printf("Done drwaning rendering %d frames\n", len(images))
	f, _ := os.OpenFile("defrag.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

}

func drawCell(img *image.Paletted, c color.Color, offsetX, offsetY, locSize int) {
	for cellX := 0; cellX < locSize-2; cellX++ {
		for cellY := 0; cellY < locSize-2; cellY++ {
			img.Set(offsetX+cellX+1, offsetY+cellY+1, c)
		}
	}
}
