package vehicle

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/ClaireGuerrini/vehicle-server/pkg/httputil"
	"github.com/ClaireGuerrini/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// http.Error(rw, "Not Implemented", http.StatusInternalServerError)
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		fmt.Println("error parsing string to int")
		return
	}
	deleteVehicle, err := d.store.Vehicle().Delete(r.Context(), id)
	if err != nil {
		d.logger.Error(
			"Could not delete the  vehicle",
			zap.Error(err),
		)
		httputil.ServeError(rw, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(deleteVehicle)
	if deleteVehicle {
		rw.WriteHeader(http.StatusNoContent)
		return
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}
