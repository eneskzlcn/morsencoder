package morsencoder_test

import (
	"errors"
	"github.com/eneskzlcn/morsencoder/internal/mocks"
	"github.com/eneskzlcn/morsencoder/internal/morsencoder"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createServiceMocks(t *testing.T) *mocks.MockLogger {
	ctrl := gomock.NewController(t)
	mockLogger := mocks.NewMockLogger(ctrl)
	return mockLogger
}
func TestNewService(t *testing.T) {
	mockLogger := createServiceMocks(t)

	t.Run("given nil logger then it should return nil", func(t *testing.T) {
		service := morsencoder.NewService(nil)
		assert.Nil(t, service)
	})
	t.Run("given valid logger then it should return service", func(t *testing.T) {
		service := morsencoder.NewService(mockLogger)
		assert.NotNil(t, service)
	})
}
func TestService_Encode(t *testing.T) {
	mockLogger := createServiceMocks(t)
	service := morsencoder.NewService(mockLogger)
	t.Run("given invalid text to encode then it should return empty string with InvalidTextToEncode error", func(t *testing.T) {
		mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Times(2)
		givenInvalidText := "{{}}"
		encodedText, err := service.Encode(givenInvalidText)
		assert.Equal(t, encodedText, "")
		assert.True(t, errors.Is(err, morsencoder.InvalidTextToEncode))
	})
	t.Run("given valid one word text to encode then it should return morse encoded text with letters seperated by spaces", func(t *testing.T) {
		mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Times(2)
		givenText := "va"
		expectedMorseEncodedText := "...- .-"
		encodedText, err := service.Encode(givenText)
		assert.Equal(t, expectedMorseEncodedText, encodedText)
		assert.Nil(t, err)
	})
	t.Run("given valid more than one word text to encode then it should return morse encoded text with letters seperated by spaces and words seperated by ' / ' backslash between two spaces.", func(t *testing.T) {
		mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Times(2)
		givenText := "va lid"
		expectedMorseEncodedText := "...- .- / .-.. .. -.."
		encodedText, err := service.Encode(givenText)
		assert.Equal(t, expectedMorseEncodedText, encodedText)
		assert.Nil(t, err)
	})
}
