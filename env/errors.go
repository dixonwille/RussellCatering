package env

import "errors"

//EnvrionmentError is an error returned by this package for any environment errors
type EnvironmentErrors struct {
	e     error
	addOn string
	env   string
}

var (
	//ErrEnvNotFound is used if environment variable is not found
	ErrEnvNotFound = errors.New("Could not find environment varibale")
	//ErrEnvWrongType is used if envrionmnet variable is of wrong type
	ErrEnvWrongType = errors.New("Envrionment variable is of wrong type")
)

//NewEnvironmentError creates a new environment error
func NewEnvironmentError(e error, env, addOn string) *EnvironmentErrors {
	return &EnvironmentErrors{e: e, env: env, addOn: addOn}
}

func (e EnvironmentErrors) Error() string {
	message := e.e.Error() + ": " + e.env
	if e.addOn != "" {
		message += ": " + e.addOn
	}
	return message
}

//IsEnvNotFound returns true if error is an ErrEnvNotFound
func IsEnvNotFound(e error) bool {
	err, ok := e.(EnvironmentErrors)
	if !ok {
		return false
	}
	if err.e.Error() != ErrEnvNotFound.Error() {
		return false
	}
	return true
}

//IsEnvWrongType returns true if error is an ErrEnvWrongType
func IsEnvWrongType(e error) bool {
	err, ok := e.(EnvironmentErrors)
	if !ok {
		return false
	}
	if err.e.Error() != ErrEnvWrongType.Error() {
		return false
	}
	return true
}
