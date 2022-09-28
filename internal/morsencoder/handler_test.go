package morsencoder_test

import (
	"errors"
	"fmt"
	"github.com/eneskzlcn/morsencoder/internal/mocks"
	"github.com/eneskzlcn/morsencoder/internal/morsencoder"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createHandlerMocks(t *testing.T) (*mocks.MockLogger, *mocks.MockMorseEncoderService) {
	ctrl := gomock.NewController(t)
	mockLogger := mocks.NewMockLogger(ctrl)
	mockService := mocks.NewMockMorseEncoderService(ctrl)
	return mockLogger, mockService
}

func setupFiberAppForHandler(handler *morsencoder.Handler) *fiber.App {
	app := fiber.New()
	app.Get("/encode", handler.Encode)
	return app
}

func makeTestRequestWithoutBody(app *fiber.App, method string, route string) (*http.Response, error) {
	req := httptest.NewRequest(method, route, nil)
	return app.Test(req)
}
func assertStringBodyEqual(t *testing.T, responseBody io.Reader, expectedValue string) {
	bodyBytes, err := io.ReadAll(responseBody)
	assert.Nil(t, err)
	actualValue := string(bodyBytes)
	assert.Equal(t, actualValue, expectedValue)
}
func TestNewHandler(t *testing.T) {
	mockLogger, mockService := createHandlerMocks(t)
	t.Run("given nil logger then it should return nil", func(t *testing.T) {
		handler := morsencoder.NewHandler(mockService, nil)
		assert.Nil(t, handler)
	})
	t.Run("given nil service then it should return nil", func(t *testing.T) {
		handler := morsencoder.NewHandler(nil, mockLogger)
		assert.Nil(t, handler)
	})
	t.Run("given not nil service and logger then it should return handler", func(t *testing.T) {
		handler := morsencoder.NewHandler(mockService, mockLogger)
		assert.NotNil(t, handler)
	})
}

func TestHandler_Encode(t *testing.T) {
	mockLogger, mockService := createHandlerMocks(t)
	handler := morsencoder.NewHandler(mockService, mockLogger)
	mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).AnyTimes()
	app := setupFiberAppForHandler(handler)
	t.Run("given invalid text to encode then it should return status bad request", func(t *testing.T) {
		givenInvalidText := "invalid"
		mockService.EXPECT().Encode(givenInvalidText).Return("", morsencoder.InvalidTextToEncode)
		resp, err := makeTestRequestWithoutBody(app, fiber.MethodGet, fmt.Sprintf("/encode?text=%s", givenInvalidText))
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("given valid text to encode but an error occurred on server then it should return internal server error", func(t *testing.T) {
		givenValidText := "valid"
		mockService.EXPECT().Encode(givenValidText).Return("", errors.New("any error different than InvalidTextToEncode"))
		resp, err := makeTestRequestWithoutBody(app, fiber.MethodGet, fmt.Sprintf("/encode?text=%s", givenValidText))
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
	t.Run("given valid text to encode then it should return morse encoded text and status ok", func(t *testing.T) {
		givenValidText := "valid"
		expectedEncodedText := "...-+.-+.-..+..+-.."
		mockService.EXPECT().Encode(givenValidText).Return(expectedEncodedText, nil)
		resp, err := makeTestRequestWithoutBody(app, fiber.MethodGet, fmt.Sprintf("/encode?text=%s", givenValidText))
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		assertStringBodyEqual(t, resp.Body, expectedEncodedText)
	})
}

func TestHandler_RegisterRoutes(t *testing.T) {
	mockLogger, mockService := createHandlerMocks(t)
	handler := morsencoder.NewHandler(mockService, mockLogger)
	app := fiber.New()
	handler.RegisterRoutes(app)
	mockLogger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Times(1)
	mockService.EXPECT().Encode(gomock.Any()).Return("", nil).Times(1)
	req := httptest.NewRequest(fiber.MethodGet, "/encode", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.NotEqual(t, fiber.StatusNotFound, resp.Status)

}
