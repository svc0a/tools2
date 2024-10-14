package game2

import (
	"bitbucket.org/alibaba-international/ghs-common/model"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
)

type gameName string
type gameCode string

type Game interface {
	GetGameCode() gameCode
}

type codes map[model.Environment]map[gameCode]gameName

func (c *codes) GetGameName(env model.Environment, code gameCode) (*gameName, errorx.Error) {
	m, ok := (*c)[env]
	if !ok {
		return nil, errorx.New("game not found")
	}
	name, ok := m[code]
	if !ok {
		return nil, errorx.New("game not found")
	}
	return &name, nil
}

type games[T Game] map[model.Environment]map[gameName]T

func (c *games[T]) GetGame(env model.Environment, game gameName) (*T, errorx.Error) {
	m, ok := (*c)[env]
	if !ok {
		return nil, errorx.New("game not found")
	}
	code, ok := m[game]
	if !ok {
		return nil, errorx.New("game not found")
	}
	return &code, nil
}

func (c *games[T]) GetGameCode(env model.Environment, game gameName) (*gameCode, errorx.Error) {
	m, ok := (*c)[env]
	if !ok {
		return nil, errorx.New("game not found")
	}
	code, ok := m[game]
	if !ok {
		return nil, errorx.New("game not found")
	}
	code1 := code.GetGameCode()
	return &code1, nil
}
