package game2

import (
	"bitbucket.org/alibaba-international/ghs-common/model"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
)

type GameName string
type GameCode string

type GameObject interface {
	GetGameName() *GameName
	GetGameCode() *GameCode
}

type Codes[T GameObject] map[model.Environment]map[GameCode]T

func (c Codes[T]) GetGameName(env model.Environment, code GameCode) (*T, errorx.Error) {
	m, ok := c[env]
	if !ok {
		return nil, errorx.New("game not found")
	}
	obj, ok := m[code]
	if !ok {
		return nil, errorx.New("game not found")
	}
	return &obj, nil
}

type Games[T GameObject] map[model.Environment]map[GameName]T

func (c Games[T]) GetGame(env model.Environment, game GameName) (*T, errorx.Error) {
	m, ok := c[env]
	if !ok {
		return nil, errorx.New("game not found")
	}
	obj, ok := m[game]
	if !ok {
		return nil, errorx.New("game not found")
	}
	return &obj, nil
}
