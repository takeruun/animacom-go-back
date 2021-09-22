package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	BindJSON(interface{}) error
	Status(int)
	JSON(code int, obj interface{})
	Header(key string, value string)
}
