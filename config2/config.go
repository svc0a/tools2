package config2

import (
	"bitbucket.org/alibaba-international/ghs-common/model"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
	"github.com/svc0a/tools2/error2"
)

type Configs[T any] struct {
	Data map[model.Environment]T
}

func (c *Configs[T]) GetConfig(env model.Environment) (*T, errorx.Error) {
	conf1, ok := c.Data[env]
	if !ok {
		return nil, error2.ErrEnvNotFound
	}
	return &conf1, nil
}
