package main

import (
	"image"
	"image/draw"
	"image/jpeg"

	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func main() {
	imgFile, err := os.Open("certificate.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	// new img
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	// font
	fontBytes, err := os.ReadFile("CrimsonText-Bold.ttf")
	if err != nil {
		log.Fatal(err)
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	// ctx for drawing
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.Black)

	// name
	c.SetFontSize(40)
	_, err = c.DrawString("john doe", freetype.Pt(400, 300))
	if err != nil {
		log.Fatal(err)
	}

	// link
	c.SetFontSize(20)
	_, err = c.DrawString("certisure.com/credentials/123456", freetype.Pt(300, 400))
	if err != nil {
		log.Fatal(err)
	}

	// qr
	qr, err := qr.Encode("https://certisure.vercel.app/credential/123456", qr.M, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}
	qr, err = barcode.Scale(qr, 200, 200)
	if err != nil {
		log.Fatal(err)
	}
	draw.Draw(rgba, image.Rect(50, 50, 250, 250), qr, image.Point{}, draw.Over)

	outFile, err := os.Create("certificate_new.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	jpeg.Encode(outFile, rgba, nil)

	log.Println("check certificate_new")
}
