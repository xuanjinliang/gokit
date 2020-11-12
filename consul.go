package gotool

import (
	"errors"
	"github.com/hashicorp/consul/api"
	"strings"
)

var ErrConsulKeyNotExist = errors.New("consul: key not exist")

type Consul struct {
	address string
	health  *api.Health
	kv      *api.KV
}

func NewConsul(address string) *Consul {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		panic("connect consul failed, err:" + err.Error())
	}

	consul := &Consul{
		address: address,
		health:  client.Health(),
		kv:      client.KV(),
	}
	return consul
}

func (c *Consul) GetKey(key string) (*string, error) {
	pair, _, err := c.kv.Get(key, nil)
	if err != nil {
		return nil, err
	}
	if pair == nil {
		return nil, ErrConsulKeyNotExist
	}
	str := string(pair.Value)
	return &str, nil
}

func (c *Consul) GetService(service string) ([]*api.ServiceEntry, error) {
	entry, _, err := c.health.Service(service, "", true, nil)
	return entry, err
}

func (c *Consul) PutKey(key string, str string) error {
	value := []byte(str)
	pair := &api.KVPair{Key: key, Value: value}
	_, err := c.kv.Put(pair, nil)
	return err
}

func (c *Consul) GetList(key string, cut ...string) (*map[string]string, error) {
	pair, _, err := c.kv.List(key, nil)
	if err != nil {
		return nil, err
	}
	if len(pair) == 0 {
		return nil, ErrConsulKeyNotExist
	}
	keyCut := ""
	if len(cut) > 0 {
		keyCut = cut[0]
	}

	mapData := make(map[string]string)

	for _, d := range pair {
		k := strings.ReplaceAll(d.Key, keyCut, "")
		v := string(d.Value)
		mapData[k] = v
	}

	return &mapData, nil
}