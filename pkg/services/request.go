package services

import (
	"errors"
	"net/http"
)

func CheckRequestMethod(req http.Request, w http.ResponseWriter, need string) (bool, error) {
	if req.Method != need {
		methodError := errors.New("this method for this route is not supported")
		err := BadResponse(w, methodError.Error())
		if err != nil{
			return false, methodError
		}
		return false, methodError
	}
	return true, nil
}