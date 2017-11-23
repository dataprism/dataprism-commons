package consul

import (
	"strings"
	"golang.org/x/net/context"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"reflect"
	"errors"
)

type ConsulStorage struct {
	client *api.KV
}

type Pair struct {
	Key string
	Value []byte
}

func NewStorage(client *api.Client) *ConsulStorage {
	return &ConsulStorage{client: client.KV()}
}

func (s *ConsulStorage) List(ctx context.Context, prefix string) ([]*Pair, error) {
	prefix = strings.TrimPrefix(prefix,"/")
	if ! strings.HasSuffix(prefix, "/") { prefix += "/" }

	var res []*Pair

	pairs, _, err := s.client.List(prefix, &api.QueryOptions{})
	if err != nil {
		return nil, err
	} else {
		if pairs != nil {
			for _, p := range pairs {
				idx := strings.Index(p.Key[len(prefix):], "/")

				if idx == -1 { continue }

				pair := &Pair{p.Key, p.Value}
				res = append(res, pair)
			}
		}

		return res, nil
	}
}

func (s *ConsulStorage) Get(ctx context.Context, key string) (*Pair, error) {
	key = strings.TrimPrefix(key,"/")

	logrus.Info(reflect.TypeOf(key))
	data, _, err := s.client.Get(key, &api.QueryOptions{})
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("requested key returned empty")
	}

	return &Pair{data.Key, data.Value}, nil
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
