package handlers

import (
	"github.com/Simbory/mego"
	"github.com/Simbory/mego/session"
)

func Init(server *mego.Server) {
	server.Get("/", home)
	server.Any("/views/*pathInfo", renderView)
	server.Get("/date/<year:int>-<month:int>-<day:int>", getDate)
	server.Get("/session", testSession)
	server.Get("/uuid", testUUID)
	server.Any("/filter/*pathInfo", testFilter)
	session.Default().RegisterType(&userModel{})
}