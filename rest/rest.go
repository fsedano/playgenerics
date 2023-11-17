package rest

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func GetData[T any](url string, val *T) error {

	client := resty.New()

	resp, err := client.R().
		SetResult(val).
		Get(url)

	if err != nil {
		log.Printf("Some erro: %s", err)
		return err
	}
	if resp.StatusCode() > 399 {
		log.Printf("Error ins tatus code: %d", resp.StatusCode())
		return fmt.Errorf("error in status code: %d", resp.StatusCode())
	}

	//log.Printf("Resp=%s", resp)
	//log.Printf("Parsed=%s", *val)
	return nil
}
