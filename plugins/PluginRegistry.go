package plugins

import (
	"errors"
	"github.com/lytics/logrus"
)

type PluginRegistry struct {
	plugins map[string]DataprismPlugin
}

func NewDataprismPluginRegistry() *PluginRegistry {
	return &PluginRegistry{
		make(map[string]DataprismPlugin, 1),
	}
}

func (r *PluginRegistry) Add(p DataprismPlugin) error {
	if _, ok := r.plugins[p.Id()]; ok {
		return errors.New("a plugin with id " + p.Id() + " has already been registered")
	}

	logrus.Infof("Added the %s plugin", p.Id())
	r.plugins[p.Id()] = p

	return nil
}



func (r *PluginRegistry) Plugins() map[string]DataprismPlugin {
	return r.plugins
}

func (r *PluginRegistry) Get(id string) DataprismPlugin {
	return r.plugins[id]
}