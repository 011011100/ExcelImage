package img

import (
	"fmt"
	"image"
	"os"
)

func GetImgFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("get file err:", err)
		return nil
	}
	return file

}

func GetImgHAndW(file *os.File) (int, int) {
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("get file err:", err)
		return 0, 0
	}
	return c.Width, c.Height
}

func ChangToHEX(r int, g int, b int) string {
	str := "0123456789ABCDEF"
	var HEXstr string

	Intcolor := []int{
		0: r,
		1: g,
		2: b,
	}

	for i, _ := range Intcolor {
		intA := Intcolor[i] / 16
		intB := Intcolor[i] % 16
		HEXstr += string(str[intA])
		HEXstr += string(str[intB])
	}

	fmt.Println(HEXstr)
	return HEXstr
}

func ChangToLetter(intParam int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var LetterStr string

	intA := (intParam-1) / 26

	intB := (intParam-1) % 26

	if intA != 0 {
		LetterStr += string(str[intA-1])
	}

	LetterStr += string(str[intB])
	return LetterStr
}

func CopyColor() {

}

