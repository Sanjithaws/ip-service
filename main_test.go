package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ipconfig", getIP)

	cases := []struct {
		name       string
		headers    map[string]string
		remoteAddr string
		want       string
	}{
		{"direct", nil, "1.2.3.4:5678", "1.2.3.4"},
		{"x-real-ip", map[string]string{"X-Real-IP": "5.6.7.8"}, "", "5.6.7.8"},
		{"x-forwarded-for", map[string]string{"X-Forwarded-For": "9.9.9.9, 10.0.0.1"}, "", "9.9.9.9"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/ipconfig", nil)
			for k, v := range tc.headers {
				req.Header.Set(k, v)
			}
			if tc.remoteAddr != "" {
				req.RemoteAddr = tc.remoteAddr
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Contains(t, w.Body.String(), tc.want)
		})
	}
}
