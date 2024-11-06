package models

type Response struct {
	Code         int         `json:"-"`
	Response     interface{} `json:"response"`
	ResponseTime int64       `json:"response_time"`
	Length       int         `json:"length"`
}
