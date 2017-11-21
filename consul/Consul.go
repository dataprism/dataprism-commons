package consul

import (
	"strings"
	"golang.org/x/net/context"
	"github.com/hashicorp/consul/api"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type ConsulStorage struct {
	client *api.KV
}

func NewStorage(client *api.Client) *ConsulStorage {
	return &ConsulStorage{client: client.KV()}
}

func (s *ConsulStorage) List(ctx context.Context, prefix string, fields []string) ([]map[string]interface{}, error) {
	prefix = strings.TrimPrefix(prefix,"/")
	if ! strings.HasSuffix(prefix, "/") { prefix += "/" }

	var res []map[string]interface{}

	pairs, _, err := s.client.List(prefix, &api.QueryOptions{})
	if err != nil {
		return nil, err
	} else {
		if pairs != nil {
			for _, p := range pairs {
				idx := strings.Index(p.Key[len(prefix):], "/")

				if idx != -1 { continue }

				var data map[string]interface{}
				err := json.Unmarshal(p.Value, &data)
				if err != nil {
					logrus.Error(err)
				} else {
					res = append(res, data)
				}
			}
		}

		return res, nil
	}
}

func (s *ConsulStorage) Get(ctx context.Context, key string) ([]byte, error) {
	key = strings.TrimPrefix(key,"/")

	data, _, err := s.client.Get(key, &api.QueryOptions{})
	if err != nil {
		return nil, err
	}

	return data.Value, nil
}

func (m *ConsulStorage) Set(ctx context.Context, key string, value []byte) (error) {
	key = strings.TrimPrefix(key,"/")

	pair := &api.KVPair{Key: key, Value: value}

	_, err := m.client.Put(pair, &api.WriteOptions{})

	return err
}

func (m *ConsulStorage) Remove(ctx context.Context, key string) (error) {
	key = strings.TrimPrefix(key,"/")

	_, err := m.client.DeleteTree(key, &api.WriteOptions{})

	return err
}