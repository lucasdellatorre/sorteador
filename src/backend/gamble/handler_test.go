package gamble

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type contextSetup struct {
	Method  string
	Path    string
	Body    string
	Cookies []*http.Cookie
}

func setupContext(setup contextSetup) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(setup.Method, setup.Path, strings.NewReader(setup.Body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

// func TestCreateGambleOK(t *testing.T) {
// 	setup := contextSetup{
// 		Method: http.MethodPost,
// 		Path:   "/create",
// 		Body:   `{"name": "lucas", "cpf": "14041364019", "numbers": [1,2,3,4,5]}`,
// 	}

// 	c, rec := setupContext(setup)

// 	h := NewHandler(nil) // have to mock

// 	if assert.NoError(t, h.Create(c)) {
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 	}
// }

func TestCreateGambleInvalidCPF(t *testing.T) {
	setup := contextSetup{
		Method: http.MethodPost,
		Path:   "/create",
		Body:   `{"name": "lucas", "cpf": "12345678900", "numbers": [1,2,3,4,5]}`,
	}

	c, rec := setupContext(setup)

	h := NewHandler(nil)

	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateGambleMissingName(t *testing.T) {
	setup := contextSetup{
		Method: http.MethodPost,
		Path:   "/create",
		Body:   `{"cpf": "14041364019", "numbers": [1,2,3,4,5]}`,
	}

	c, rec := setupContext(setup)

	h := NewHandler(nil)

	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateGambleMissingPayload(t *testing.T) {
	setup := contextSetup{
		Method: http.MethodPost,
		Path:   "/create",
	}

	c, rec := setupContext(setup)

	h := NewHandler(nil)

	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateGambleWrongNumbersLength(t *testing.T) {
	setup := contextSetup{
		Method: http.MethodPost,
		Path:   "/create",
		Body:   `{"name": "lucas", "cpf": "14041364019", "numbers": [1,2,3,4]}`,
	}

	c, rec := setupContext(setup)

	h := NewHandler(nil)

	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateGambleWrongNumbersValues(t *testing.T) {
	setup := contextSetup{
		Method: http.MethodPost,
		Path:   "/create",
		Body:   `{"name": "lucas", "cpf": "14041364019", "numbers": [100,1,2,4,-50]}`,
	}

	c, rec := setupContext(setup)

	h := NewHandler(nil)

	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
