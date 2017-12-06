package core

import (
	"github.com/dataprism/dataprism-commons/storage"
	"github.com/dataprism/dataprism-commons/execute"
	consul "github.com/hashicorp/consul/api"
	nomad "github.com/hashicorp/nomad/api"
	"github.com/dataprism/dataprism-commons/config"
)

type Platform struct {
	Settings *config.Configuration
	KV *storage.ConsulStorage
	Scheduler *execute.NomadScheduler
}

func InitializePlatform() (*Platform, error) {
	// -- create the consul client
	consulClient, err := consul.NewClient(consul.DefaultConfig())
	if err != nil { return nil, err }

	// -- create the consul storage
	store := storage.NewConsulStorage(consulClient.KV())

	// -- use consul to retrieve the configuration
	settings, err := config.Load(store)
	if err != nil { return nil, err }

	// -- initialize nomad
	nomadClient, err := nomad.NewClient(nomad.DefaultConfig())
	if err != nil { return nil, err }

	// -- create the scheduler
	scheduler := execute.NewNomadScheduler(nomadClient, settings.JobsDir)

	return &Platform{
		Settings: settings,
		KV: store,
		Scheduler: scheduler,
	}, nil
}
