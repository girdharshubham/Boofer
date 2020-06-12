package service

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	req         *http.Request
	err         error
	username             = "girdharshubham"
	password             = "password"
	routesToHit []string = []string{"/healthz", "/about", "/v1/"}
)

func beforeAll() (r *gin.Engine) {
	gin.SetMode(gin.TestMode)
	r = gin.Default()
	return
}

func TestV1Routes(t *testing.T) {
	r := beforeAll()
	V1Routes(r)

	for _, routeToHit := range routesToHit {
		w := httptest.NewRecorder()
		if req, err = http.NewRequest(http.MethodGet, routeToHit, nil); err != nil {
			t.Fail()
		}
		req.SetBasicAuth(username, password)
		r.ServeHTTP(w, req)

		assert.Equalf(t, http.StatusOK, w.Code, "Expected %d, got %d", http.StatusOK, w.Code)

	}
}
