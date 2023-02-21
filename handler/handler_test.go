package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/timecheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	timeNow := time.Now().Unix()
	rr := httptest.NewRecorder()
	th := TimeHandler()
	handler := http.Handler(th)
	handler.ServeHTTP(rr, req)

	var responseBody SampleResponse
	json.Unmarshal([]byte(rr.Body.String()), &responseBody)

	assert.GreaterOrEqual(t, timeNow, responseBody.Time)
	assert.Equal(t, http.StatusOK, rr.Code, "Status code didn't match")
}
