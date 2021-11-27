package main

import (
	"fmt"
	"image"
	"os"
	"time"

	"github.com/golang/freetype/truetype"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
)

const (
	wait   = 26300 * time.Nanosecond
	Width  = 128
	Height = 64
)

func main() {

	board := raspi.NewAdaptor()
	oled := i2c.NewSSD1306Driver(board)

	work := func() {
		oled.Clear()
		img := createImage("hello world!!")
		if err := oled.ShowImage(img); err != nil {
			fmt.Println(err)
		}
	}
	robot := gobot.NewRobot("screenBot",
		[]gobot.Connection{board},
		[]gobot.Device{oled},
		work,
	)

	go robot.Start()
	time.Sleep(5 * time.Second)

	fmt.Println("robot end")
	robot.Stop()
}
func createImage(text string) image.Image {
	img := image.NewGray(image.Rect(0, 0, Width, Height))
	textSize := 15
	opt := truetype.Options{
		Size: float64(textSize),
	}
	ft, err := truetype.Parse(gobold.TTF)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	face := truetype.NewFace(ft, &opt)

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	dr.Dot.X = (fixed.I(Width) - dr.MeasureString(text)) / 2
	dr.Dot.Y = (fixed.I(Height) + fixed.I(textSize)) / 2
	dr.DrawString(text)
	return img
}
