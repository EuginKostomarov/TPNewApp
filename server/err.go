package TPserver

import (
	"errors"
	"net/http"
)

func ReturnHTTPERR(res *http.Response) error {
	switch s := res.StatusCode; s {
	case 404:
		return errors.New("Сервис не доступен")
	case 500:
		return errors.New("Ошибка на стороне сервера")
	}
	return nil
}
