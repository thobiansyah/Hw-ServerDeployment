package service

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Getdata() ([]byte, error) {
	response, err := http.Get("https://jsonkeeper.com/b/DMXK")
	if err != nil {
		fmt.Errorf("404 (Not Found)", err.Error())

		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf(err.Error())

		return nil, err
	}

	return data, nil

}
