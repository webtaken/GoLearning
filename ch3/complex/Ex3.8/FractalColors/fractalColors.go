package fractalColors

import "image/color"

var Mapping [16]color.RGBA

func init() {
	Mapping[0] = color.RGBA{66, 30, 15, 255}
	Mapping[1] = color.RGBA{25, 7, 26, 255}
	Mapping[2] = color.RGBA{9, 1, 47, 255}
	Mapping[3] = color.RGBA{4, 4, 73, 255}
	Mapping[4] = color.RGBA{0, 7, 100, 255}
	Mapping[5] = color.RGBA{12, 44, 138, 255}
	Mapping[6] = color.RGBA{24, 82, 177, 255}
	Mapping[7] = color.RGBA{57, 125, 209, 255}
	Mapping[8] = color.RGBA{134, 181, 229, 255}
	Mapping[9] = color.RGBA{211, 236, 248, 255}
	Mapping[10] = color.RGBA{241, 233, 191, 255}
	Mapping[11] = color.RGBA{248, 201, 95, 255}
	Mapping[12] = color.RGBA{255, 170, 0, 255}
	Mapping[13] = color.RGBA{204, 128, 0, 255}
	Mapping[14] = color.RGBA{153, 87, 0, 255}
	Mapping[15] = color.RGBA{106, 52, 3, 255}
}
