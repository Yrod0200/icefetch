package main

import (
	"icefetch/src/helpers"
)

func main() {

	info := helpers.GenerateInfoSlice()
	logo := helpers.GenerateLogoSlice()

	helpers.Render(logo, info)
}
