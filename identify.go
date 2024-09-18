package hap

import (
	"github.com/brutella/hap/log"

	"net/http"
)

func (s *Server) identify(res http.ResponseWriter, req *http.Request) {
	if s.IsPaired() {
		log.Info.Printf("request only valid if unpaired")
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	if s.a.IdentifyFunc != nil {
		s.a.IdentifyFunc(req)
	}

	res.WriteHeader(http.StatusNoContent)
}
