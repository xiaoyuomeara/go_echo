package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetGhost(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/echo/statement")

	// Assertions
	if assert.NoError(t, getGhost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "oooOOOoooOOOooo\n", rec.Body.String())
	}

}

func TestGetEcho(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/echo/statement", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/echo/")
	c.SetParamNames("statement")
	c.SetParamValues("statement")

	// Assertions
	if assert.NoError(t, getEcho(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "statement statement\n", rec.Body.String())
	}

}
