package main

import (
	"log"
	"fmt"
	"strings"
)

func sb(s string) []uint8 {
	
    return []uint8(s)
}

func bs (b []uint8) string {
    return string(b)
}
 
func main() {
	log.Println()
	text1 := "Мы+получили+текст,+в+котором+по+какой-то+ошибке+все+пробелы+заменились+на+плюсы.+Надо+это+исправить"
    bytes := sb(text1)
    fmt.Println(bytes)

    text2 := bs(bytes)
    fmt.Println(text2)

	plusText := "Мы+получили+текст,+в+котором+по+какой-то+ошибке+все+пробелы+заменились+на+плюсы.+Надо+это+исправить"
	correctedText := ReplacePlusWithSpace(plusText)
	fmt.Println(correctedText)
	
	
	/*text := "Hello"
 	encryptedText := shifrC(text)
 	fmt.Println("Text1:", text)
	fmt.Println("Text2:", encryptedText)*/
}	
	// Функция для замены символов "+" на пробел
func ReplacePlusWithSpace(text string) string {
		replacedText := strings.ReplaceAll(text, "+", " ")
		return replacedText
}


/*// Функция для шифрования текста по Шифру Цезаря (сдвиг на 3 буквы)
func shifrC(text string) string {
 shift := +3
 runes := []rune(text)
 for i, char := range runes {
  if char >= 'A' && char <= 'Z' {
   runes[i] = 'A' + (char-'A'+ rune(shift))
  } else if char >= 'a' && char <= 'z' {
   runes[i] = 'a' + (char-'a'+ rune(shift))
  }
 }
 return string(runes)
}
*/
