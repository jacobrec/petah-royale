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

        if hasOverlap(raster, x, y, w, h) {
            n--
            continue
        }

        for i := x; i < x + w; i++ {
            for j := y; j < y + h; j++ {
                raster[i][j] = true
            }
        }

        if n > 0 {
            drillForcedPath(raster, x, y, w, h, width, height)
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

func drillForcedPath(raster[][]bool, xb int, yb int, w int, h int, width int, height int) {
    i := false
    for !i {
        x := xb + rand.Intn(w)
        y := yb + rand.Intn(h)

        switch rand.Intn(4) {
            case 0:
                i = drillPath(raster, width, height, x, y, 1, 0)
            case 1:
                i = drillPath(raster, width, height, x, y, -1, 0)
            case 2:
                i = drillPath(raster, width, height, x, y, 0, 1)
            case 3:
                i = drillPath(raster, width, height, x, y, 0, -1)
        }

    }
}

func hasOverlap(raster [][]bool, x int, y int, w int, h int) bool {
    for i := x; i < x+w; i++ {
        for j := y; j < y+h; j++ {
            if raster[i][j] {
                return true;
            }
        }
    }
    return false;
}

func drillPath(raster [][]bool, xdim int, ydim int, x int, y int, xdir int, ydir int) bool {
    state := raster[x][y]
    if !state {
        return false
    }

    startx := x
    starty := y

    for true {
        x += xdir
        y += ydir
        if x < 0 || x >= xdim || y < 0 || y >= ydim {
            return false
        }

        if !state && raster[x][y] {
            break
        }

        state = raster[x][y]
    }

    for startx != x || starty != y {
        raster[x][y] = true
        raster[x+1][y] = true
        raster[x+1][y+1] = true
        raster[x][y+1] = true
        x -= xdir
        y -= ydir
    }

    return true
}


