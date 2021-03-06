package containerserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/atc/api/present"
	"github.com/pivotal-golang/lager"
)

func (s *Server) GetContainer(w http.ResponseWriter, r *http.Request) {
	handle := r.FormValue(":id")

	hLog := s.logger.Session("container", lager.Data{
		"handle": handle,
	})

	container, found, err := s.db.GetContainer(handle)
	if err != nil {
		hLog.Error("failed-to-lookup-container", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !found {
		hLog.Debug("container-not-found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	hLog.Debug("found-container")

	presentedContainer := present.Container(container)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(presentedContainer)
}
