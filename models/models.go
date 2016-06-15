package models

import "encoding/json"

//Publicer is an interface to state whether a model can be made to view publicly
type Publicer interface {
	Public() interface{}
}

//Jsonify turns the interface into a json object as a slice of bytes
func Jsonify(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

//Error is used to return an error json object back
type Error struct {
	Message string `json:"message"`
}

//NewError returns a new Error struct to use
func NewError(msg string) *Error {
	return &Error{
		Message: msg,
	}
}
