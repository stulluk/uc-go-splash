package main

import (
	"os"
	"fmt"
	"flag"
	"time"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"image/gif"
	"src/framebuffer"
)

func die(err interface{}) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 { die("usage: fbshow file") }
	name := flag.Args()[0]
	file, err := os.Open(name)
	if err != nil { die(err) }
	_, str, err := image.DecodeConfig(file)
	if err != nil { die(err) }
	_, err = file.Seek(0, 0)
	if err != nil { die(err) }
	fb, err := framebuffer.Open("/dev/fb0")
	if err != nil { die(err) }
	if str == "gif" {
		all, err := gif.DecodeAll(file)
		if err != nil { die(err) }
		acc := image.NewRGBA(all.Image[0].Bounds())
		dst := acc.Bounds().Sub(acc.Bounds().Min).Add(fb.Bounds().Min).Add(fb.Bounds().Size().Sub(acc.Bounds().Size()).Div(2))
		src := acc.Bounds().Min
		for {
			for idx, img := range all.Image {
				begin := time.Now()
				draw.Draw(acc, acc.Bounds(), img, src, draw.Over)
				draw.Draw(fb, dst, acc, src, draw.Src)
				time.Sleep(time.Duration(all.Delay[idx]) * 10 * time.Millisecond - time.Since(begin))
			}
		}
		return
	}
	img, _, err := image.Decode(file)
	if err != nil { die(err) }
	dst := img.Bounds().Sub(img.Bounds().Min).Add(fb.Bounds().Min).Add(fb.Bounds().Size().Sub(img.Bounds().Size()).Div(2))
	src := img.Bounds().Min
	draw.Draw(fb, dst, img, src, draw.Src)
}

