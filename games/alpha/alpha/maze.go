package alpha

import (
    "math/rand"
)


func MakeMaze(rooms int, width int, height int, density float32) ([]Immoveable, func() (float64,float64)) {
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

    spawner := func() (float64, float64) {
        x := rand.Intn(width)
        y := rand.Intn(height)
        for !raster[x][y] {
            x = rand.Intn(width)
            y = rand.Intn(height)
        }
        return float64(x), float64(y)
    }

    imv := make([]Immoveable, 0)

    for i := 0; i < width; i++ {
        carry := 0
        for j := 0; j < height; j++ {
            if !raster[i][j] {
                carry++
            } else if carry > 0 {
                // optimize columns
                imv = append(imv, Immoveable{float64(i), float64(j-carry), 1, float64(carry)})
                carry = 0
            }
        }
        if (carry > 0) {
            imv = append(imv, Immoveable{float64(i), float64(height-carry), 1, float64(carry)})
        }
    }


    return imv, spawner
}

