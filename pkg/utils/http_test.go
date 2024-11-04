package utils_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go_api_template/pkg/utils"
)

func TestRespondJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{"message": "test"}

	utils.RespondJSON(w, http.StatusOK, data)

	var response utils.JSONResponse
	err := json.NewDecoder(w.Body).Decode(&response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, data, response.Data)
}

func TestRespondError(t *testing.T) {
	w := httptest.NewRecorder()
	message := "error message"

	utils.RespondError(w, http.StatusInternalServerError, message)

	var response utils.JSONError
	err := json.NewDecoder(w.Body).Decode(&response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.Status)
	assert.Equal(t, message, response.Message)
}

func TestParseJSON(t *testing.T) {
	data := map[string]string{"message": "test"}
	body, _ := json.Marshal(data)
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

	var result map[string]string
	err := utils.ParseJSON(r, &result)

	assert.NoError(t, err)
	assert.Equal(t, data, result)
}
