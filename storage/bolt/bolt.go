package bolt

import (
	"github.com/asdine/storm"
	"github.com/filebrowser/filebrowser/auth"
	"github.com/filebrowser/filebrowser/settings"
	"github.com/filebrowser/filebrowser/share"
	"github.com/filebrowser/filebrowser/storage"
	"github.com/filebrowser/filebrowser/users"
)

// NewStorage creates a storage.Storage based on Bolt DB.
func NewStorage(db *storm.DB) (*storage.Storage, error) {
	users := users.NewStorage(usersBackend{db: db})
	share := share.NewStorage(shareBackend{db: db})
	settings := settings.NewStorage(settingsBackend{db: db})
	auth := auth.NewStorage(authBackend{db: db}, users)

	err := save(db, "version", 2)
	if err != nil {
		return nil, err
	}

	return &storage.Storage{
		Auth:     auth,
		Users:    users,
		Share:    share,
		Settings: settings,
	}, nil
}
