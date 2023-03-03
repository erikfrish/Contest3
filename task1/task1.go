package main

import "fmt"

func main() {
	var letters int
	var name string
	var colors string

	fmt.Scanln(&letters)
	fmt.Scanln(&name)
	fmt.Scanln(&colors)
	var res = beautifulWords(letters, name, colors)
	fmt.Println(res)
}

func beautifulWords(letters int, name string, colors string) uint16 {

	var colorsofwords []string
	colorsofwords = append(colorsofwords, "")
	flag := true
	count := uint16(0)

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
