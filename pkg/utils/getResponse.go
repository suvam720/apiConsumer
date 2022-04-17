package utils

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

//function to get http response from a url
func Get(ctx context.Context, url string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return checkStatus(resp)
}

func checkStatus(h *http.Response) ([]byte, error) {
	if h.StatusCode == http.StatusOK {
		dataByte, err := ioutil.ReadAll(h.Body)
		if err != nil {
			return []byte{}, err
		}
		return dataByte, nil
	} else {
		return []byte{}, errors.New(strconv.Itoa(http.StatusNotFound))
	}
}