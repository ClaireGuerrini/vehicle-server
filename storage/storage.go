package storage

import (
	"github.com/ClaireGuerrini/vehicle-server/storage/vehiclestore"
)

type Store interface {
	Vehicle() vehiclestore.Store
}
