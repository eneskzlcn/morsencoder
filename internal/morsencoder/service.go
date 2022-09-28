package morsencoder

import "strings"

type Service struct {
	logger Logger
}

func NewService(logger Logger) *Service {
	if logger == nil {
		return nil
	}
	return &Service{
		logger: logger,
	}
}
func (s *Service) Encode(textToEncode string) (string, error) {
	s.logger.Debugf("Encoding the text %s", textToEncode)
	//morse alphabet is not case-sensitive. And our system works with uppercase.
	textToEncode = strings.ToUpper(textToEncode)
	encodedText, err := s.encodeText(textToEncode, " ", " / ")
	s.logger.Debugf("Given text %s encoded to %s", textToEncode, encodedText)
	return encodedText, err
}
func (s *Service) encodeText(textToEncode, letterSplitter, wordSplitter string) (string, error) {
	wordsInGivenText := strings.Split(textToEncode, " ")
	morseEncodedWords := make([]string, 0)
	for _, word := range wordsInGivenText {
		encodedWord, err := s.encodeWord(word, letterSplitter)
		if err != nil {
			return "", err
		}
		morseEncodedWords = append(morseEncodedWords, encodedWord)
	}
	return strings.Join(morseEncodedWords, wordSplitter), nil
}
func (s *Service) encodeWord(wordToEncode, letterSplitter string) (string, error) {
	lettersInWord := strings.Split(wordToEncode, "")
	morseEncodedLetters := make([]string, 0)
	for _, letter := range lettersInWord {
		morseEncodedLetter, err := s.encodeLetter(letter)
		if err != nil {
			return "", err
		}
		morseEncodedLetters = append(morseEncodedLetters, morseEncodedLetter)
	}
	return strings.Join(morseEncodedLetters, letterSplitter), nil
}
func (s *Service) encodeLetter(letterToEncode string) (string, error) {
	if morseEncodedLetter, exists := morseCodesMap[letterToEncode]; exists {
		return morseEncodedLetter, nil
	}
	return "", InvalidTextToEncode
}
