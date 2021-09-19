package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(code int, obj interface{})
}
