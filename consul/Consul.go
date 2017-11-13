package consul

import (
	"strings"
	"golang.org/x/net/context"
	"github.com/hashicorp/consul/api"
)

type ConsulStorage struct {
	client *api.KV
}

func NewStorage(client api.Client) *ConsulStorage {
	return &ConsulStorage{client: client.KV()}
}

func (s *ConsulStorage) List(ctx context.Context, prefix string) ([]string, error) {
	prefix = strings.TrimPrefix(prefix,"/")
	if ! strings.HasSuffix(prefix, "/") { prefix += "/" }

	pairs, _, err := s.client.Keys(prefix, "/", &api.QueryOptions{})
	if err != nil {
		return nil, err
	} else {
		if pairs == nil {
			return []string{}, nil
		}

		var res []string
		for _, p := range pairs {
			idx := strings.Index(p, "/")

			if idx == -1 {
				continue
			}

			idx2 := strings.Index(p[idx + 1:], "/")

			if idx2 == -1 {
				idx2 = len(p[idx + 1:])
			}

			res = append(res, p[idx + 1:][:idx2])
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