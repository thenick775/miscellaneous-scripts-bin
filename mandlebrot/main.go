//plots the Mandelbrot set with some added coloring
package main

import (
	"flag"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sync"
)

func mandel(x int, y int, startx float64, starty float64, deltax float64, deltay float64, maxit int, thresh float64, image *image.RGBA, wg *sync.WaitGroup) {
	defer wg.Done()
	var userx float64 = startx + (float64(x) * deltax)
	var usery float64 = starty - (float64(y) * deltay)
	var xnext, ynext, xstore, ystore, pdist float64
	xnext, ynext, xstore, ystore, pdist = 0.0, 0.0, 0.0, 0.0, 0.0

	for i := 0; i < maxit; i++ {
		xnext = xstore*xstore - ystore*ystore + userx
		ynext = 2.0*xstore*ystore + usery

		if i+1 >= maxit && xnext*xnext+ynext*ynext < thresh {
			image.Set(x, y, color.RGBA{170 - uint8(int(math.Floor(pdist))), 200 - uint8(int(math.Floor(pdist))), 255 - uint8(int(math.Floor(pdist))), 255})
		} else if i+1 < maxit+1 && xnext*xnext+ynext*ynext > thresh {
			image.Set(x, y, color.RGBA{190 - uint8(int(math.Floor(pdist)))*10, 140 - uint8(int(math.Floor(pdist)))*10, 235 - uint8(int(math.Floor(pdist)))*10, 255})
			break
		}

		pdist += math.Sqrt((xstore-xnext)*(xstore-xnext) + (ystore-ynext)*(ystore-ynext))

		xstore, ystore = xnext, ynext
	}
}

func main() {
	var maxit, antialias, widtho, heighto int
	var urad, currxuser, curryuser, thresh float64
	var fname string

	flag.IntVar(&maxit, "i", 1000, "number of iterations (max)")
	flag.IntVar(&antialias, "AA", 4, "anti alias (scaling of width/height)")
	flag.IntVar(&widtho, "w", 1080, "width of image")
	flag.IntVar(&heighto, "h", 1080, "height of image")
	flag.Float64Var(&thresh, "t", 250, "threshold number (for mandlebrot pixel value)")
	flag.Float64Var(&urad, "r", 2.5, "radius applied to start/end points")
	flag.Float64Var(&currxuser, "uxp", -0.75, "user coordinate x center position")
	flag.Float64Var(&curryuser, "uyp", 0.0, "user coordinate y center position")
	flag.StringVar(&fname, "fname", "mandle.png", "output file name")
	flag.Parse()

	widthr, heightr := widtho*antialias, heighto*antialias //anti aliasing, resized at end back to orginal
	ratio := float64(heightr) / float64(widthr)
	radius := urad * ratio
	yrad := float64(radius * ratio)

	xstart, xend := currxuser-radius/2.0, currxuser+radius/2.0
	ystart, yend := curryuser-yrad/2.0, curryuser+yrad/2.0

	var deltax, deltay float64 = float64(xend-xstart) / float64(widthr), float64(yend-ystart) / float64(heightr)

	image := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{widthr, heightr}})

	var wg sync.WaitGroup
	for x := 0; x < widthr; x++ {
		for y := 0; y < heightr; y++ {
			wg.Add(1)
			go mandel(x, y, xstart, yend, deltax, deltay, maxit, thresh, image, &wg)
		}
	}
	wg.Wait()

	image_resized := resize.Resize(uint(widtho), uint(heighto), image, resize.Lanczos3)

	out_file, _ := os.Create(fname)
	png.Encode(out_file, image_resized)

	out_file.Close()

}
