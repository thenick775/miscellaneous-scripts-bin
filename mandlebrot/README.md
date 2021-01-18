## Mandelbrot Set Plotting Utility


![](graphics/mandel_defaul.png)
![](graphics/mandel_zoomed.png)

This quick utility plots the Mandelbrot set concurrently, allowing for customizable zoom, positioning, anti aliasing, and image size.

Usage:
  -AA int
    	anti alias (scaling of width/height) (default 4)
  -fname string
    	output file name (default "mandle.png")
  -h int
    	height of image (default 1080)
  -i int
    	number of iterations (max) (default 1000)
  -r float
    	radius applied to start/end points (default 2.5)
  -t float
    	threshold number (for mandlebrot pixel value) (default 250)
  -uxp float
    	user coordinate x center position (default -0.75)
  -uyp float
    	user coordinate y center position
  -w int
    	width of image (default 1080)

