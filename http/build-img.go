package http

import (
	"errors"
	"net/http"
)

var buildImgHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {

	paths, ok := r.URL.Query()["path"]
	if !ok {
		err := errors.New("can not get path in query string")
		return errToStatus(err), err
	}

	path := paths[0]

	err := d.RunHook(func() error {
		return nil
	}, "build", path, "", d.user)

	if err != nil {
		return errToStatus(err), err
	}

	return http.StatusOK, nil
})
