package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultQueryParam    = ""
	defaultQueryIntParam = 0
	maxPostMemmory       = 1024
)

type IContext interface {
	Request() *http.Request

	Writer() http.ResponseWriter

	Body() ([]byte, error)

	ReadJSON(interface{}) error

	Write(data []byte) (int, error)

	Header(key string, value string)

	GetHeader(key string) string

	URL() *url.URL

	RequestURI() string

	URLParamDefault(key string, def string) string

	URLParam(key string) string

	URLParamInt(key string) (int, error)

	form() (map[string][]string, bool)

	PostParamDefault(key string, def string) string

	PostParam(key string) string

	PostParamInt(key string) (int, error)

	CtxParam(key string) (interface{}, bool)

	GetCookie(key string) (*http.Cookie, error)

	SetCookie(cookie *http.Cookie)
}

type Context struct {
	contextData map[string]interface{}
	r           *http.Request
	w           http.ResponseWriter
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		contextData: make(map[string]interface{}),
		r:           r,
		w:           w,
	}
}

func (c *Context) Request() *http.Request {
	return c.r
}

func (c *Context) Writer() http.ResponseWriter {
	return c.w
}

func (c *Context) Body() ([]byte, error) {
	return ioutil.ReadAll(c.r.Body)
}

func (c *Context) ReadJSON(dest interface{}) error {
	body, err := c.Body()

	if err == nil {
		err = json.Unmarshal(body, dest)
	}

	return err
}

func (c *Context) Write(data []byte) (int, error) {
	return c.w.Write(data)
}

func (c *Context) GetHeader(key string) string {
	return c.r.Header.Get(key)
}

func (c *Context) Header(key string, value string) {
	if value == "" {
		c.w.Header().Del(key)
		return
	}

	c.w.Header().Add(key, value)
}

func (c *Context) URL() *url.URL {
	return c.r.URL
}

func (c *Context) RequestURI() string {
	return c.r.RequestURI
}

func (c *Context) URLParamDefault(key string, def string) string {
	if value := c.r.URL.Query().Get(key); value != def {
		return value
	}

	return def
}

func (c *Context) URLParam(key string) string {
	return c.URLParamDefault(key, defaultQueryParam)
}

func (c *Context) URLParamInt(key string) (int, error) {
	value := c.URLParam(key)

	if value != defaultQueryParam {
		param, err := strconv.Atoi(value)
		if err != nil {
			return defaultQueryIntParam, errors.New("Wrong param format")
		}

		return param, nil
	}

	return defaultQueryIntParam, errors.New("URL Param not found")
}

func (c *Context) form() (map[string][]string, bool) {
	c.r.ParseMultipartForm(maxPostMemmory)

	/*if form := c.r.Form; len(form) > 0 {
		return form, true
	}*/

	if form := c.r.PostForm; len(form) > 0 {
		return form, true
	}

	if m := c.r.MultipartForm; m != nil {
		if len(m.Value) > 0 {
			return m.Value, true
		}
	}

	return nil, false
}

func (c *Context) PostParamDefault(key string, def string) string {
	form, exist := c.form()
	if !exist {
		return def
	}

	values, ok := form[key]
	if !ok {
		return def
	}

	return values[0]
}

func (c *Context) PostParam(key string) string {
	return c.PostParamDefault(key, defaultQueryParam)
}

func (c *Context) PostParamInt(key string) (int, error) {
	value := c.PostParam(key)

	if value != defaultQueryParam {
		param, err := strconv.Atoi(value)

		if err != nil {
			return defaultQueryIntParam, errors.New("Wrong param format")
		}

		return param, nil
	}

	return defaultQueryIntParam, errors.New("POST param not found")
}

func (c *Context) CtxParam(key string) (interface{}, bool) {
	value, ok := c.contextData[key]
	return value, ok
}

func (c *Context) GetCookie(key string) (*http.Cookie, error) {
	return c.r.Cookie(key)
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	c.r.AddCookie(cookie)
}