package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errros "github.com/SlavaculoN/golang-todoapp/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf(
			"no key='%s' in path values: %w",
			key,
			core_errros.ErrInvalidArgument,
		)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"path value'%s' by key=''%s not a valid integer: %v: %w",
			pathValue,
			key,
			err,
			core_errros.ErrInvalidArgument,
		)
	}

	return val, nil
}
