package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestHttpWriter struct {
	code int
}

func (w *TestHttpWriter) Header() http.Header {
	return nil
}

func (w *TestHttpWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

func (w *TestHttpWriter) WriteHeader(statusCode int) {
	w.code = statusCode
}

type HttpTestCase struct {
	method string
	url    string
	code   int
}

func Authorize(ctx IContext) {
	ctx.Write([]byte("OK"))
	ctx.StatusCode(http.StatusOK)
}

func TestRouting(t *testing.T) {
	apiRouter := NewRouter(nil)
	apiRouter.AddHandlerGet("/test", Authorize)

	testCases := []*HttpTestCase{
		&HttpTestCase{http.MethodGet, "/test", http.StatusOK},
		&HttpTestCase{http.MethodPost, "/empty", http.StatusNotFound},
	}

	for _, testCase := range testCases {
		request := httptest.NewRequest(testCase.method, testCase.url, nil)
		writer := &TestHttpWriter{}

		apiRouter.ServeHTTP(writer, request)

		if writer.code != testCase.code {
			t.Fatalf("Expected status: %v\nGot status: %v", testCase.code, writer.code)
		}
	}
}
