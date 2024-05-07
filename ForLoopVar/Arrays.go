package main

import "fmt"

func deleteValueFromSlice(value string) {
	languages := []string{"Go", "JavaScript", "Ruby", "Python"}
	fmt.Println("Original slice:", languages)

	// Index of element to remove
	indexToRemove := -1
	for i, lang := range languages {
		if lang == value {
			indexToRemove = i
			break
		}
	}

	// If element found, remove it
	if indexToRemove != -1 {
		languages = append(languages[:indexToRemove], languages[indexToRemove+1:]...)
	}

	fmt.Println("Slice after removing 'JavaScript':", languages)
}
func main() {
	var vegetables [2]string

	vegetables[0] = "Broccoli"
	vegetables[1] = "Carrot"

	fmt.Println(vegetables)
	fmt.Println(vegetables[1]) //Carrot

	vegetableArr := [4]string{"Broccoli", "Carrot"}

	fmt.Println(vegetableArr)
	fmt.Println(len(vegetableArr))
	fmt.Println(vegetableArr[1]) //Carrot

	var vegArr []string

	vegArr = append(vegArr, "Teste 1")
	fmt.Println(len(vegArr))
	fmt.Println(vegArr)

	languages := []string{"Go", "JavaScript", "Ruby", "Python"}

	fmt.Println(languages)
	fmt.Println(len(languages))
	fmt.Println(languages[0])
	fmt.Println(languages[1:3])
	languages = append(languages, "PHP")
	fmt.Println(languages)

	deleteValueFromSlice("JavaScript")
}
