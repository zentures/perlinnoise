package perlinnoise

import (
    "testing"
    "fmt"
)

func TestPerlin1DGenerator(t *testing.T) {
    fmt.Println("var noiseData = [");
    for i := 0; i < 360; i++ {
        p := PerlinNoise1d(float64(i), 0.2, 0.5)
        fmt.Println("{x:",i,",y:",p,"},");
    }
    fmt.Println("];");
}
