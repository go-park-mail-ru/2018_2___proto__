package api

type ApiResponse struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
