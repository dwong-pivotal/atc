package pipelineserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/atc/auth"
)

func (s *Server) ListAllPipelines(w http.ResponseWriter, r *http.Request) {
	logger := s.logger.Session("list-all-pipelines")
	teamName := auth.GetAuthOrDefaultTeamName(r)

	pipelines, err := s.getPipelinesForTeam(teamName)
	if err != nil {
		logger.Error("failed-to-get-all-active-pipelines", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(pipelines)
}