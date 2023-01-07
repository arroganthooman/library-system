package presentation

import "fmt"

var ErrorUnauthorized error = fmt.Errorf("You are not authorized to do this action. Check your bearer token.")
