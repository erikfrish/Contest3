package main

import "fmt"

func main() {
	var letters int
	var name string
	var colors string

	// fmt.Scanln(&letters)
	// fmt.Scanln(&name)
	// fmt.Scanln(&colors)

	// letters = 27
	// name = "Algorithms and Data Structures"
	// colors = "BBBBBBBBBB" + "BYB" + "YYYY" + "BBBBBBBBBB"
	// //3

	// letters = 7
	// name = "Tinkoff"
	// colors = "BYBYBYB"
	// //0

	// letters = 20
	// name = "ABBA huba buba buba huba"
	// colors = "BYBY" + "BYBY" + "BBBY" + "BYYY" + "YBBB"
	// //3

	letters = 20
	name = "ABBA huba buba buba huba GUIG"
	colors = "BYBY" + "BYBY" + "BBBY" + "YBYB" + "YBBB" + "YBYB"
	//2
	var res = beautifulWords(letters, name, colors)
	// fmt.Printf("\n%v", res)
	fmt.Print(res)

}

func beautifulWords(letters int, name string, colors string) int {

	var colorsofwords []string
	colorsofwords = append(colorsofwords, "")
	flag := true
	count := 0

	i, j := 0, 0
	for _, v := range name {
		if v == ' ' {
			colorsofwords = append(colorsofwords, "")
			j++
			continue
		}
		colorsofwords[j] += string(colors[i])
		i++
	}
	fmt.Println(colorsofwords)

	for i := range colorsofwords {
		for j := 0; j < len(colorsofwords[i])-1; j++ {
			if colorsofwords[i][j] == colorsofwords[i][j+1] {
				flag = false
			}
		}
		if !flag {
			count++
			flag = true
		}
	}
	return count
}
