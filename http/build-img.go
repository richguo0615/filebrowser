package http

import (
	"errors"
	"net/http"
	"strings"
	"time"
)


var canBuild = time.Now().Unix()

var buildImgHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {

	if time.Now().Unix() < canBuild {
		err := errors.New("please wait for last one shell request done")
		return errToStatus(err), err
	}
	canBuild = nextTimeBuild()

	paths, ok := r.URL.Query()["path"]
	if !ok {
		err := errors.New("can not get path in query string")
		return errToStatus(err), err
	}

	path := paths[0]
	path = strings.Replace(path, "/files", "", -1) + "tmp.png"

	err := d.RunHook(func() error {
		return nil
	}, "build", path, "", d.user)

	if err != nil {
		return errToStatus(err), err
	}

	return http.StatusOK, nil
})

func nextTimeBuild() int64 {
	return time.Now().Unix() + 10
}
