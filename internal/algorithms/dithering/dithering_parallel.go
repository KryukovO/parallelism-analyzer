package dithering

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sync"
)

// Функция, реализующая параллельную версию упорядоченного размытия
func OrderedDitheringParallel(srcImgPath string, dstImgPath string, order int, threadCount int) (errF error) {
	defer func() {
		if msg := recover(); msg != nil {
			errF = fmt.Errorf("%v", msg)
		}
	}()

	D, err := ditheringMatrix(order)
	if err != nil {
		return err
	}

	srcImgReader, err := os.Open(srcImgPath)
	if err != nil {
		return err
	}
	defer srcImgReader.Close()
	m, _, err := image.Decode(srcImgReader)
	if err != nil {
		return err
	}
	bounds := m.Bounds()

	dstImg := image.NewRGBA(bounds)

	if threadCount < 1 {
		return errors.New("число потоков должно быть больше нуля")
	}

	stepY := int(math.Max(math.Ceil(float64(bounds.Max.Y-bounds.Min.Y)/float64(threadCount)), 1))
	var wg sync.WaitGroup

	for thread := 0; thread < threadCount; thread++ {
		startY := bounds.Min.Y + stepY*thread
		if startY < bounds.Max.Y {
			wg.Add(1)
			go func() {
				for y := startY; y < int(math.Min(float64(startY+stepY), float64(bounds.Max.Y))); y++ {
					for x := bounds.Min.X; x < bounds.Max.X; x++ {
						i := (y - bounds.Min.Y) % order
						j := (x - bounds.Min.X) % order
						r, g, b, a := m.At(x, y).RGBA()
						intens := int(math.Max(float64(r/256), math.Max(float64(g/256), float64(b/256))))
						if (intens*order*order+1)/256 > D[i][j] {
							dstImg.Set(x, y, color.RGBA{255, 255, 255, uint8(a / 256)})
						} else {
							dstImg.Set(x, y, color.RGBA{0, 0, 0, uint8(a / 256)})
						}
					}
				}
				wg.Done()
			}()
		}
	}

	dstImgFile, err := os.Create(dstImgPath)
	if err != nil {
		return err
	}

	wg.Wait()

	png.Encode(dstImgFile, dstImg)

	return nil
}
