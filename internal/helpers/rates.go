package helpers

import "net/http"

func GetCurrentDolarRate() (*http.Response, error) {
	url := "https://uruguayapi.onrender.com/api/v1/banks/brou_rates"

	resp, err := http.Get(url)

	return resp, err
}
