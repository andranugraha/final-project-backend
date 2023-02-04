package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"

	"final-project-backend/server"

	"github.com/gin-gonic/gin"
)

func ServeReq(opts *server.RouterConfig, req *http.Request) (*gin.Engine, *httptest.ResponseRecorder) {
	router := server.NewRouter(opts)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return router, rec
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

func MakeRequestBodyMultiPartFormData(dto map[string]string) (*bytes.Buffer, *multipart.Writer) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	for key, val := range dto {
		part, _ := writer.CreateFormField(key)
		io.Copy(part, bytes.NewReader([]byte(val)))
	}

	return body, writer
}
