package utils

type Result struct {
	IsError bool
	Code int
	Message string
	Data interface{}
}