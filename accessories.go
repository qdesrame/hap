package hap

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/log"

	"net/http"
)

func (s *Server) getAccessories(res http.ResponseWriter, req *http.Request) {
	if !s.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	var as []*accessory.A
	as = append(as, s.a)
	as = append(as, s.as[:]...)

	p := struct {
		Accessories []*accessory.A `json:"accessories"`
	}{as}

	log.Debug.Println(toJSON(p))
	JsonOK(res, p)
}
