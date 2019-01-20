package alpha

import (
    "math/rand"
    "fmt"
)


func MakeMaze(rooms int, width int, height int, density float32) []Immoveable {
    meanwidth := int(float32(width)*density)
    meanheight := int(float32(height)*density)


    raster := make([][]bool, width)
    for i := range raster {
        raster[i] = make([]bool, height)
    }

    for n := 0; n < rooms; n++ {
        x := rand.Intn(width - meanwidth)
        y := rand.Intn(height - meanheight)
        w := meanwidth - 2 + rand.Intn(4)
        h := meanheight - 2 + rand.Intn(4)

        for i := x; i < x + w; i++ {
            for j := y; j < y + h; j++ {
                raster[i][j] = true
            }
        }
    }

    imv := make([]Immoveable, 0)

    for i := 0; i < width; i++ {
        for j := 0; j < height; j++ {
            if !raster[i][j] {
                imv = append(imv, Immoveable{float64(i), float64(j), 1, 1})
            }
        }
    }

    return imv
}

