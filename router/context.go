package router

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
	Logger() ILogger

	Request() *http.Request

	Writer() http.ResponseWriter

	Body() ([]byte, error)

	ReadJSON(json.Unmarshaler) error

	Write(data []byte) (int, error)

	Header(key string, value string)

	GetHeader(key string) string

	URL() *url.URL

	RequestURI() string

	//возвращает параметры запроса из урлы (/user?id=12&var=val) по имени
	//если отсутствует параметр, то возвращает дефолтное значение
	QueryParamDefault(key string, def string) string

	QueryParam(key string) string

	QueryParamInt(key string) (int, error)

	//возвращает все параметры из api урлы
	//тип урлы /user/2
	UrlParams() map[string]string

	SetApiParser(parser IApiUrlParser)

	//парсит тело запроса и вовращает все найденные параметры
	Form() (map[string][]string, bool)

	//возвращает параметр из тела запроса по имени
	//если параметр с таким иеменем отсутствует, то возвращает дефолтное значение
	PostParamDefault(key string, def string) string

	PostParam(key string) string

	PostParamInt(key string) (int, error)

	//возвращает данные, которые хранятся в контексте
	//можно храить любые типы, но придется кастовать их
	//будет полезно, если нужно передать по конвейеру http запроса какие-нибудь данные
	CtxParam(key string) (interface{}, bool)

	AddCtxParam(key string, value interface{})

	GetCookie(key string) (*http.Cookie, error)

	SetCookie(cookie *http.Cookie) error

	//возвращает метод запроса
	//можно использовать для фильтрации запросов по типу
	Method() string

	StatusCode(code int)

	ContentType(content string)
}

type Context struct {
	statusChanged bool
	contextData   map[string]interface{}
	apiUrlParser  IApiUrlParser
	log           ILogger
	r             *http.Request
	w             http.ResponseWriter
}

func NewContext(w http.ResponseWriter, r *http.Request, log ILogger) *Context {
	return &Context{
		statusChanged: false,
		contextData:   make(map[string]interface{}),
		log:           log,
		r:             r,
		w:             w,
	}
}

func (c *Context) Logger() ILogger {
	return c.log
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

func (c *Context) ReadJSON(entity json.Unmarshaler) error {
	body, err := c.Body()

	if err == nil {
		err = entity.UnmarshalJSON(body)
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

func (c *Context) QueryParamDefault(key string, def string) string {
	if value := c.r.URL.Query().Get(key); value != def {
		return value
	}

	return def
}

func (c *Context) QueryParam(key string) string {
	return c.QueryParamDefault(key, defaultQueryParam)
}

func (c *Context) QueryParamInt(key string) (int, error) {
	value := c.QueryParam(key)

	if value != defaultQueryParam {
		param, err := strconv.Atoi(value)
		if err != nil {
			return defaultQueryIntParam, errors.New("Wrong param format")
		}

		return param, nil
	}

	return defaultQueryIntParam, errors.New("URL Param not found")
}

func (c *Context) Form() (map[string][]string, bool) {
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
	form, exist := c.Form()
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

func (c *Context) AddCtxParam(key string, value interface{}) {
	c.contextData[key] = value
}

func (c *Context) GetCookie(key string) (*http.Cookie, error) {
	return c.r.Cookie(key)
}

func (c *Context) SetCookie(cookie *http.Cookie) error {
	if c.statusChanged {
		return errors.New("Cookie cannot be set after status code")
	}

	http.SetCookie(c.w, cookie)
	return nil
}

func (c *Context) Method() string {
	return c.r.Method
}

func (c *Context) StatusCode(code int) {
	c.statusChanged = true
	c.w.WriteHeader(code)
}

func (c *Context) UrlParams() map[string]string {
	return c.apiUrlParser.Parse(c.r.URL.Path)
}

func (c *Context) SetApiParser(parser IApiUrlParser) {
	c.apiUrlParser = parser
}

func (c *Context) ContentType(cType string) {
	c.w.Header().Set("Content-Type", cType)
}
