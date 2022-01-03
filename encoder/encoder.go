package encoder

import (
	"strings"
)

var morseCodesMap = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	".": ".-.-.-",
	",": "--..--",
	"?": "..--..",
	"!": "-.-.--",
	"-": "-....-",
	"/": "-..-.",
	"@": ".--.-.",
	"(": "-.--.",
	")": "-.--.-",
}

func EncodeToMorseCode(text string,letterSplitter string) string{
	var encodedText string
	text = strings.ToUpper(text)
	seperatedText := strings.Split(text, " ")
	for _,word := range seperatedText{
		encodedWord := encodeWord(word,letterSplitter)
		if encodedWord != ""{
			encodedText+= encodedWord
		}
	}
	return encodedText
}
func encodeWord(word string,letterSplitter string) string{
	var encodedWord string
	for i:= 0 ; i < len(word); i++{
		morseCode:= morseCodesMap[word[i:i+1]]
		if morseCode != ""{
			encodedWord+= morseCode + letterSplitter
		}
	}
	return encodedWord
}