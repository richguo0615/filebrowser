package http

import "net/http"

var buildImgHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {

	err := d.RunHook(func() error {
		return nil
	}, "build", r.URL.Path, "", d.user)

	if err != nil {
		return errToStatus(err), err
	}

	return http.StatusOK, nil
})
