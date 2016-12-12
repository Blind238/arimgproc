package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

// how to use the above so the output at image.Decode is converted/cast

func main() {
	f, err := os.Open("images/sample.png")
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	//nData := myImage(mData) => cannot convert mData (type image.Image) to type myImage
	n := m.(*image.NRGBA)

	fmt.Println("read and decoded file, printing colors")
	//p := []uint8{} // for R,G,B,A
	p := make([]uint8, 4) // for R,G,B,A

	for i := 0; i < len(n.Pix); i++ {
		p[i%4] = n.Pix[i]
		fmt.Println(n.Pix[i])

		if i%4 == 3 {
			fmt.Println(p)
		}
	}

	imgTpl := `<html><head></head><body>
	<img src="data:image/jpg;base64,{{.ImageOrig}}">
	<img src="data:image/jpg;base64,{{.ImageNew}}">
	</body></html>`

}
