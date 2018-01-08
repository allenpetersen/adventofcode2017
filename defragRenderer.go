package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x2d, 0x2d, 0x2d, 0xff},
	color.RGBA{0x31, 0xd6, 0x89, 0xff},
	color.RGBA{0x32, 0x32, 0x32, 0xff},
	color.RGBA{0xff, 0xa8, 0xd9, 0xff},
	color.RGBA{0xea, 0x9c, 0xc8, 0xff},
	color.RGBA{0xd6, 0x90, 0xb7, 0xff},
	color.RGBA{0xc1, 0x84, 0xa6, 0xff},
	color.RGBA{0xad, 0x78, 0x96, 0xff},
	color.RGBA{0x98, 0x6d, 0x85, 0xff},
	color.RGBA{0x84, 0x61, 0x74, 0xff},
	color.RGBA{0x6f, 0x55, 0x64, 0xff},
	color.RGBA{0x5b, 0x49, 0x53, 0xff},
	color.RGBA{0x46, 0x3d, 0x42, 0xff},
}
var locSize = 7

type oldGroup struct {
	frame int
	locs  []diskLocation
}

func renderDefrag(events <-chan diskEvent) {
	gridSize := 128
	var w, h int = locSize * 128, locSize * 128

	var images []*image.Paletted
	var delays []int
	oldGroups := []oldGroup{}
	frame := 0
	for e := range events {
		if e.kind == groupComplete {
			oldGroups = append(oldGroups, oldGroup{frame, e.disk.group})
			continue
		} else if e.kind == groupCleared {
			continue
		}

		frame++

		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				offsetX := x * locSize
				offsetY := y * locSize

				var c color.Color
				if e.disk.grid[x][y] {
					c = palette[1]
				} else {
					c = palette[2]
				}

				drawCell(img, c, offsetX, offsetY)
			}
		}

		for _, l := range e.disk.group {
			offsetX := l.x * locSize
			offsetY := l.y * locSize
			drawCell(img, palette[3], offsetX, offsetY)
		}

		for _, g := range oldGroups {
			for _, l := range g.locs {
				offsetX := l.x * locSize
				offsetY := l.y * locSize
				drawCell(img, getFadeColor(frame-g.frame), offsetX, offsetY)
			}
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

func getFadeColor(age int) color.Color {
	switch {
	case age < 5:
		return palette[4]
	case age < 10:
		return palette[5]
	case age < 15:
		return palette[6]
	case age < 20:
		return palette[7]
	case age < 30:
		return palette[8]
	case age < 40:
		return palette[9]
	case age < 60:
		return palette[10]
	case age < 80:
		return palette[11]
	default:
		return palette[12]
	}
}

func drawCell(img *image.Paletted, cell color.Color, offsetX, offsetY int) {
	for x := 0; x < locSize; x++ {
		for y := 0; y < locSize; y++ {
			if x == 0 || x == locSize-1 || y == 0 || y == locSize-1 {
				// border
				img.Set(offsetX+x, offsetY+y, palette[0])
			} else {
				// cell
				img.Set(offsetX+x, offsetY+y, cell)
			}
		}
	}
}

func getCellColor(event *diskEvent, x, y int) color.Color {
	for _, l := range event.disk.group {
		if l.x == x && l.y == y {
			if event.kind == groupCleared {
				return palette[4]
			}
			return palette[3]
		}
	}

	if event.disk.grid[x][y] {
		return palette[1]
	}
	return palette[2]
}
