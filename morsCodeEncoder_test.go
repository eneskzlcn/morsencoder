package main

import (
	"MorseCodeEncoder/encoder"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
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
func TestIGetAStringAsParameterSuccessfullyFromRequest(t *testing.T){
	testText := "AB"
	app := fiber.New()

	var sentTextToEncode string
	app.Post("/",func(c *fiber.Ctx) error{
		sentTextToEncode = c.Query("text")
		return c.SendStatus(200)
	})
	req := httptest.NewRequest(http.MethodPost,fmt.Sprintf("/?text=%s",testText), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	app.Test(req,1)
	assert.Equalf(t,testText,sentTextToEncode,"The string sent to the api is not taking correctly." )
}
func TestISentGivenStringInMorseEncodedFormatSuccessfully(t *testing.T){
	testText := "AB"
	app := fiber.New()

	app.Get("/",func(c *fiber.Ctx) error{
		var sentTextToEncode string
		sentTextToEncode = c.Query("text")
		log.Printf("Sended text to encode: %s",sentTextToEncode)
		if sentTextToEncode != ""{
			return c.SendString(encoder.EncodeToMorseCode(sentTextToEncode," "))
		}
		return c.SendStatus(200)
	})

	req := httptest.NewRequest(http.MethodGet,fmt.Sprintf("/?text=%s",testText), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp,_ := app.Test(req,1)

	responseEncodedTextAsByte,_ := ioutil.ReadAll(resp.Body)
	responseEncodedText := string(responseEncodedTextAsByte)

	assert.Equalf(t,responseEncodedText,encoder.EncodeToMorseCode(testText," "),"The api not correctly sent the given string in morse code encoded format" )
}
