package models

type Response struct {
	Errors interface{} `json:"code"`
	Data   interface{} `json:"data"`
}
