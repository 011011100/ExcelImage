package main

import (
	"ExcelText/img"
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"image"
	_ "image/color"
	_ "image/jpeg"
	"io/ioutil"
	"strconv"
)

func main() {

	//这里填写你的图片名字
	//目前只支持500*500像素的图片
	var imgName = "HaiLunAndPBoos.jpg"

	exclFile := excelize.NewFile()
	sheet := "sheet2"
	exclFile.NewSheet(sheet)

	err := exclFile.SetColWidth(sheet,"A","XFD",2.2)
	if err != nil {
		fmt.Println("@",err)
		return
	}

	imgfile,err := ioutil.ReadFile(imgName)

	if err != nil {
		fmt.Println("@",err)
		return
	}

	byteImgFile := bytes.NewBuffer(imgfile)

	imgD ,_,err := image.Decode(byteImgFile)

	if err != nil {
		fmt.Println("@",err)
		return
	}

	bounds := imgD.Bounds()

	image.NewRGBA(bounds)

	dx := bounds.Dx()
	dy := bounds.Dy()

	for i := 1; i <= dx; i++ {
		for j := 1; j <= dy; j++ {
			colorRGB := imgD.At(i,j)
			r,g,b,_ := colorRGB.RGBA()

			r_uint8 := uint8(r >>8)
			g_uint8 := uint8(g >>8)
			b_uint8 := uint8(b >>8)

			style ,err := exclFile.NewStyle(&excelize.Style{
				Fill: excelize.Fill{
					Type: "pattern",
					Color:[]string{"#"+img.ChangToHEX(int(r_uint8),int(g_uint8),int(b_uint8))},
					Pattern: 1,
				},
			})

			if err != nil {
				fmt.Println(err)
				return
			}
			strj :=strconv.Itoa(j)
			exclFile.SetCellStyle(sheet,img.ChangToLetter(i)+strj,img.ChangToLetter(i)+strj,style)

			fmt.Println(img.ChangToLetter(i)+strj)
		}
	}

	err = exclFile.SaveAs("img.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}
