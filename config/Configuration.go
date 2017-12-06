package config

import (
	"encoding/json"
	"github.com/dataprism/dataprism-commons/storage"
	"context"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	store *storage.ConsulStorage	`json:"-"`
	KafkaCluster *KafkaCluster		`json:"cluster"`
	JobsDir string					`json:"jobs_dir"`
}

type KafkaCluster struct {
	Servers []string				`json:"servers"`
	KafkaBufferMaxMs int			`json:"buffer_max_ms"`
	KafkaBufferMinMsg int			`json:"buffer_min_msg"`
}

func Load(store *storage.ConsulStorage) (*Configuration, error) {
	pair, err := store.Get(context.Background(), "config")
	if err != nil { return nil, err }

	var result Configuration
	if pair == nil {
		result = Configuration{}
	} else {
		logrus.Info(string(pair.Value))
		err = json.Unmarshal(pair.Value, &result)
		if err != nil { return nil, err }
	}

	result.store = store

	return &result, nil
}

func (m *Configuration) Store(configuration *Configuration) (error) {
	data, err := json.Marshal(configuration)
	if err != nil { return err }

	return m.store.Set(context.Background(), "config", data)
}
