package linear

import (
	"image/color"
	"math"
	"testing"
)

func TestRGB(t *testing.T) {

	t.Run("RGBFromLinearNRGBA64()", func(t *testing.T) {

		t.Run("returns correct results", func(t *testing.T) {
			cases := []struct {
				InputNRGBA64  color.NRGBA64
				ExpectedColor RGB
				ExpectedAlpha float32
			}{
				{
					InputNRGBA64:  color.NRGBA64{R: 0, G: 0, B: 0, A: 0},
					ExpectedColor: RGB{0, 0, 0},
					ExpectedAlpha: 0,
				},
				{
					InputNRGBA64:  color.NRGBA64{R: 32768, G: 32768, B: 32768, A: 32768},
					ExpectedColor: RGB{0.5, 0.5, 0.5},
					ExpectedAlpha: 0.5,
				},
				{
					InputNRGBA64:  color.NRGBA64{R: 65535, G: 0, B: 0, A: 0},
					ExpectedColor: RGB{1, 0, 0},
					ExpectedAlpha: 0,
				},
				{
					InputNRGBA64:  color.NRGBA64{R: 0, G: 65535, B: 0, A: 0},
					ExpectedColor: RGB{0, 1, 0},
					ExpectedAlpha: 0,
				},
				{
					InputNRGBA64:  color.NRGBA64{R: 0, G: 0, B: 65535, A: 0},
					ExpectedColor: RGB{0, 0, 1},
					ExpectedAlpha: 0,
				},
				{
					InputNRGBA64:  color.NRGBA64{R: 0, G: 0, B: 0, A: 65535},
					ExpectedColor: RGB{0, 0, 0},
					ExpectedAlpha: 1,
				},
			}

			for _, c := range cases {
				col, alpha := RGBFromLinearNRGBA64(c.InputNRGBA64)

				if expected, actual := c.ExpectedColor.R, col.R; math.Abs(float64(expected)-float64(actual)) > 0.001 {
					t.Errorf("Expected red component %v for input colour %+v but got %v", expected, c.InputNRGBA64, actual)
				}
				if expected, actual := c.ExpectedColor.G, col.G; math.Abs(float64(expected)-float64(actual)) > 0.001 {
					t.Errorf("Expected green component %v for input colour %+v but got %v", expected, c.InputNRGBA64, actual)
				}
				if expected, actual := c.ExpectedColor.B, col.B; math.Abs(float64(expected)-float64(actual)) > 0.001 {
					t.Errorf("Expected blue component %v for input colour %+v but got %v", expected, c.InputNRGBA64, actual)
				}
				if expected, actual := c.ExpectedAlpha, alpha; math.Abs(float64(expected)-float64(actual)) > 0.001 {
					t.Errorf("Expected alpha component %v for input colour %+v but got %v", expected, c.InputNRGBA64, actual)
				}
			}
		})
	})

	t.Run("ToLinearNRGBA64()", func(t *testing.T) {

		t.Run("returns correct results", func(t *testing.T) {
			cases := []struct {
				InputColor      RGB
				InputAlpha      float32
				ExpectedNRGBA64 color.NRGBA64
			}{
				{
					InputColor:      RGB{0, 0, 0},
					InputAlpha:      0,
					ExpectedNRGBA64: color.NRGBA64{R: 0, G: 0, B: 0, A: 0},
				},
				{
					InputColor:      RGB{0.5, 0.5, 0.5},
					InputAlpha:      0.5,
					ExpectedNRGBA64: color.NRGBA64{R: 32768, G: 32768, B: 32768, A: 32768},
				},
				{
					InputColor:      RGB{1, 0, 0},
					InputAlpha:      0,
					ExpectedNRGBA64: color.NRGBA64{R: 65535, G: 0, B: 0, A: 0},
				},
				{
					InputColor:      RGB{0, 1, 0},
					InputAlpha:      0,
					ExpectedNRGBA64: color.NRGBA64{R: 0, G: 65535, B: 0, A: 0},
				},
				{
					InputColor:      RGB{0, 0, 1},
					InputAlpha:      0,
					ExpectedNRGBA64: color.NRGBA64{R: 0, G: 0, B: 65535, A: 0},
				},
				{
					InputColor:      RGB{0, 0, 0},
					InputAlpha:      1,
					ExpectedNRGBA64: color.NRGBA64{R: 0, G: 0, B: 0, A: 65535},
				},
			}

			for _, c := range cases {
				col := c.InputColor.ToLinearNRGBA64(c.InputAlpha)

				if expected, actual := c.ExpectedNRGBA64, col; expected != actual {
					t.Errorf("Expected colour %+v for input colour %+v and alpha %v but got %+v", expected, c.InputColor, c.InputAlpha, actual)
				}
			}
		})
	})
}
