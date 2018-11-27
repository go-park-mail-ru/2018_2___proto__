package api

type ApiResponse struct {
	Code     int         `json:"Code"`
	Response interface{} `json:"Response"`
}
