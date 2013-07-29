/*
 * Based on algorithm given at http://freespace.virgin.net/hugo.elias/models/m_perlin.htm
 *
 * Borrowed heavily from https://github.com/iand/perlin/blob/master/perlin.go
 *
 * MIT License http://opensource.org/licenses/MIT
 */

package perlinnoise

import "math"

func noise(x uint64) float64 {
    fn := (x << 13) ^ x
    return (1.0 - float64((fn*(fn*fn*15731+789221)+1376312589)&0x7fffffff)/1073741824.0)
}

func smoothedNoise(x float64) float64 {
    //return Noise(x)/2  +  Noise(x-1)/4  +  Noise(x+1)/4
    xint := uint64(math.Trunc(x));

    return noise(xint)/2 + noise(xint-1)/4 + noise(xint+1)/4;
}

func interpolate(a, b, x float64) float64 {
    /*
        ft = x * 3.1415927
        f = (1 - cos(ft)) * .5

        return  a*(1-f) + b*f
    */

    ft := x * math.Pi
    f := (1 - math.Cos(ft)) * 0.5
    return a*(1-f) + b*f
}

func interpolateNoise(x float64) float64 {
    /*
        integer_X    = int(x)
        fractional_X = x - integer_X
        v1 = SmoothedNoise1(integer_X)
        v2 = SmoothedNoise1(integer_X + 1)
        return Interpolate(v1 , v2 , fractional_X)
    */

    xint := math.Trunc(x);
    xfrac := x - xint;

    v1 := smoothedNoise(xint)
    v2 := smoothedNoise(xint+1)

    return interpolate(v1, v2, xfrac);
}

func PerlinNoise1d(x, f, a float64) (value float64) {
    /*
        total = 0
        p = persistence
        n = Number_Of_Octaves - 1

        loop i from 0 to n

            frequency = 2i
            amplitude = p^i

            total = total + InterpolatedNoisei(x * frequency) * amplitude

        end of i loop

        return total
    */

    //var frequency float64 = 0.1;
    //var amplitude float64 = 1;
    octaves := 6;

    for i := 0; i < octaves; i++ {
        value += interpolateNoise(x * f) * a;
        f *= 2;
        a /= 2;
    }

    return
}
