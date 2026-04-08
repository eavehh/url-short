package req

import (
	"net/http"
	"stepik_1/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	// валидация по структурным тегам которые мы указали в payload.go (в этом же пакете)
	body, err := Decode[T](r.Body)
	if err != nil {
		println("Error decode ", err.Error())
		res.SendJson(*w, err.Error(), 402)
		return nil, err
	}

	err = IsValide(body)
	if err != nil {
		println("validation error: ", err.Error())
		res.SendJson(*w, err.Error(), 402)
		return nil, err
	}
	return &body, err
}
