package util

import (
	"context"
	"google.golang.org/appengine/memcache"
)

type (
	Memcache struct {
		Ctx context.Context
	}
)

func (m Memcache) GetValue(key string) (string, error) {
	if item, err := memcache.Get(m.Ctx, key); err != nil {
		return "", err
	} else {
		return string(item.Value), nil
	}
}
func (m Memcache) SetValue(key, value string) error {
	item := &memcache.Item{
		Key:   key,
		Value: []byte(value),
	}
	if err := memcache.Set(m.Ctx, item); err != nil {
		return err
	} else {
		return nil
	}
}
