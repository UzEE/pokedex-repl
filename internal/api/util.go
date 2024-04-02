package api

import "encoding/json"

func unmarshalJSON[T interface{}](data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)

	if err != nil {
		return *new(T), err
	}
	return v, nil
}
