package main

import (
	"MorseCodeEncoder/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

/* TEST IS PROGRAM ABLE TO ENCODE STRING TO MORSE CODE FORMAT
	* we need a library for morse alphabet to encode
	* we need a function for encode operation.
*/

// TEST IS PROGRAM ABLE TO TAKE A STRING AS PARAMETER
// TEST IS PROGRAM ABLE TO SENT BACK ENCODED MORSE CODE


func TestIEncodeGivenStringToMorseCodeSuccessfully(t *testing.T){
	testString := "AB"
	encodedString := encoder.EncodeToMorseCode(testString," ")
	assert.Equalf(t, ".- -...",encodedString,"The encoding operation not correct. Expected morse encoded string not equals actual one")
}
//func TestIGetAStringAsParameterSuccessfullyFromRequest(t *testing.T){
//
//}