package main

import (
	"MorseCodeEncoder/encoder"
	"log"
	"testing"
)

/* TEST IS PROGRAM ABLE TO ENCODE STRING TO MORSE CODE FORMAT
	* we need a library for morse alphabet to encode
	* we need a function for encode operation.
*/

// TEST IS PROGRAM ABLE TO TAKE A STRING AS PARAMETER
// TEST IS PROGRAM ABLE TO SENT BACK ENCODED MORSE CODE


func TestIEncodeGivenStringToMorseCodeSuccessfully(t *testing.T){
	testString := "Convert Me To Morse Code"
	encodedString := encoder.EncodeToMorseCode(testString," ")
	log.Println(encodedString)
}
//func TestIGetAStringAsParameterSuccessfullyFromRequest(t *testing.T){
//
//}