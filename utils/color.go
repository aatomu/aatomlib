package utils

import "fmt"

// 文字の色を変える
func SetPrintWordColor(r int, g int, b int) {
	//文字色指定
	fmt.Print("\x1b[38;2;" + fmt.Sprint(r) + ";" + fmt.Sprint(g) + ";" + fmt.Sprint(b) + "m")
}

// 文字の色を元に戻す
func ResetPrintWordColor() {
	//文字色リセット
	fmt.Print("\x1b[39m")
}

// 背景の色を変える
func SetPrintBackColor(r int, g int, b int) {
	//背景色指定
	fmt.Print("\x1b[48;2;" + fmt.Sprint(r) + ";" + fmt.Sprint(g) + ";" + fmt.Sprint(b) + "m")
}

// 背景の色を元に戻す
func ResetPrintBackColor() {
	//背景色リセット
	fmt.Print("\x1b[49m")
}
