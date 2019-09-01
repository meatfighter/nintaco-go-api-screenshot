package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/meatfighter/nintaco-go-api/nintaco"
)

const (
	imageWidth  = 256
	imageHeight = 240
)

type screenshot struct {
	api           nintaco.API
	pixels        []int
	img           *image.NRGBA
	buttonPressed bool
}

func newScreenshot() *screenshot {
	return &screenshot{
		api:    nintaco.GetAPI(),
		pixels: make([]int, imageWidth*imageHeight),
		img:    image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight)),
	}
}

func (s *screenshot) launch() {
	s.api.AddFrameListener(s)
	s.api.AddStatusListener(s)
	s.api.AddActivateListener(s)
	s.api.AddDeactivateListener(s)
	s.api.AddStopListener(s)
	s.api.Run()
}

// API.GetPixels obtains 9-bit extended palette indices. The lower 6 bits represent one
// of the 64 colors and the upper 3 bits describe if and how that color is emphasized.
// Usually, there is no emphasis (the upper 3 bits will all be 0). This method uses a
// table to convert the extended palette indices into RGBA's.
func (s *screenshot) convertToRGBA() {
	for i := len(s.pixels) - 1; i >= 0; i-- {
		p := extendedPalette[s.pixels[i]]
		j := i << 2
		s.img.Pix[j] = p.R
		s.img.Pix[j|1] = p.G
		s.img.Pix[j|2] = p.B
		s.img.Pix[j|3] = p.A
	}
}

func (s *screenshot) saveScreenshot() {
	fileName := time.Now().Format("screenshot-20060102-150405.png")
	file, e := os.Create(fileName)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	e = png.Encode(writer, s.img)
	if e != nil {
		fmt.Println(e)
	}
	writer.Flush()
	fmt.Printf("Captured image to %s\n", fileName)
}

func (s *screenshot) FrameRendered() {
	if s.api.ReadGamepad(0, nintaco.GamepadButtonSelect) {
		if !s.buttonPressed {
			s.buttonPressed = true
			s.api.GetPixels(s.pixels)
			s.convertToRGBA()
			s.saveScreenshot()
		}
	} else {
		s.buttonPressed = false
	}
}

func (s *screenshot) APIEnabled() {
	fmt.Println("API enabled")
}

func (s *screenshot) APIDisabled() {
	fmt.Println("API disabled")
}

func (s *screenshot) Dispose() {
	fmt.Println("API stopped")
}

func (s *screenshot) StatusChanged(message string) {
	fmt.Printf("Status message: %s\n", message)
}

func main() {
	nintaco.InitRemoteAPI("localhost", 9999)
	newScreenshot().launch()
}
