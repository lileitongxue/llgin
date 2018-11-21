package structerr

import "fmt"

type ParamBindErr struct {
	Err string
}

func (err ParamBindErr) Error() string {
	return fmt.Sprintf("resolve param error,reason: %s", err.Err)
}

type ParamInvalidErr struct {
	ParamName  string
	ParamValue string
}

func (err ParamInvalidErr) Error() string {

	return fmt.Sprintln("param %s is invalid,value is %s", err.ParamName, err.ParamValue)
}
