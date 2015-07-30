package sms_fx

import (
	"fmt"

	"testing"
)

func TestMyservice(t *testing.T) {

	res, err := SendCreateAccount("http://116.228.184.202:88/services/Autelan", "13051546793", 10)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Println(res.Description, res.Result, res.UserName)
}
