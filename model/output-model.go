package model

type FbLoginOut struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Message string `json:"message"`
	Type string `json:"type"`
	Code int `json:"code"`
}