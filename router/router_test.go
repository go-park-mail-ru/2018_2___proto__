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
	ctx.StatusCode(http.StatusOK)
}

func Panic(ctx IContext) {
	ctx.StatusCode(http.StatusInternalServerError)
	panic("panic")
}

func TestRouting(t *testing.T) {
	apiRouter := NewRouter(nil)
	apiRouter.AddHandlerGet("/test", Authorize)
	apiRouter.AddHandlerPost("/panic", Panic)
	apiRouter.AddHandlerDelete("/empty", Authorize)
	apiRouter.AddHandlerOptions("/empty", Authorize)
	apiRouter.AddHandlerPut("/empty", Authorize)
	apiRouter.AddHandler("/all", Authorize)

	testCases := []*HttpTestCase{
		&HttpTestCase{http.MethodGet, "/test", http.StatusOK},
		&HttpTestCase{http.MethodPost, "/panic", http.StatusInternalServerError},
		&HttpTestCase{http.MethodPost, "/empty", http.StatusNotFound},
		&HttpTestCase{http.MethodDelete, "/empty", http.StatusOK},
		&HttpTestCase{http.MethodOptions, "/empty", http.StatusOK},
		&HttpTestCase{http.MethodGet, "/all", http.StatusOK},
		&HttpTestCase{http.MethodPost, "/all", http.StatusOK},
		&HttpTestCase{http.MethodPut, "/all", http.StatusOK},
		&HttpTestCase{http.MethodDelete, "/all", http.StatusOK},
		&HttpTestCase{http.MethodOptions, "/all", http.StatusOK},
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
