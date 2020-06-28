package services

import (
	"encoding/json"
	"net/http"
)

func GoodResponse(w http.ResponseWriter, data interface{}) error{
	response, err := json.Marshal(data)
	if err != nil{
		return err
	}
	_, _ = w.Write(response)
	return nil
}


func BadResponse(w http.ResponseWriter, errorMessage string) error{
	response, err := json.Marshal(map[string]string{"error": errorMessage})
	if err != nil{
		return err
	}
	_, _ = w.Write(response)
	return nil
}