package plugins

import (
	"github.com/dataprism/dataprism-commons/api"
	"github.com/dataprism/dataprism-commons/core"
)

type DataprismPlugin interface {

	Id() string

	CreateRoutes(platform *core.Platform, API *api.Rest)

}