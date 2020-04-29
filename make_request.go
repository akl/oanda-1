package oanda

import (
	//"bytes"
	//"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//ReqArgs represents the Request Arguments needed to pass to
//MakeRequest to hit the correct Oanda endpoint
type ReqArgs struct {
	ReqMethod string
	URL       string
	Body      io.Reader
}

//MakeRequeset takes a ReqArgs as an argument and uses it to hit the
//correct Oanda endpoint to retrun a []byte and an error
func MakeRequest(ra *ReqArgs) ([]byte, error) {
	req, err := http.NewRequest(ra.ReqMethod, ra.URL, ra.Body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Connection", "Keep-Alive")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respByte, nil
}