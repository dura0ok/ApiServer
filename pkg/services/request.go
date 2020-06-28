package services

import "net/http"

func CheckRequestMethod(req http.Request, need string) bool {
	if req.Method != need {
		return false
	}
	return true
}