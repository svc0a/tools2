package game2

import (
	"bitbucket.org/alibaba-international/ghs-common/model"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type GameName string
type GameCode string

type GameObject interface {
	GetName() *GameName
	GetCode() *GameCode
}

type GameHelper[T GameObject] interface {
	GetByName(name GameName) (*T, errorx.Error)
	GetByCode(code GameCode) (*T, errorx.Error)
}

type gameHelper[T GameObject] struct {
	codes map[GameCode]T
	games map[GameName]T
}

type Helper[T GameObject] map[model.Environment]GameHelper[T]

func (h Helper[T]) Get(env model.Environment) GameHelper[T] {
	return h[env]
}

func NewHelpers[T GameObject](in string) Helper[T] {
	m2 := map[model.Environment][]T{}
	err := json.Unmarshal([]byte(in), &m2)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}
	m := Helper[T]{
		model.EnvTest: newHelper(m2[model.EnvTest]),
		model.EnvProd: newHelper(m2[model.EnvProd]),
	}
	return m
}

func newHelper[T GameObject](in []T) GameHelper[T] {
	g := gameHelper[T]{}
	for _, item := range in {
		g.codes[*item.GetCode()] = item
		g.games[*item.GetName()] = item
	}
	return &g
}

func (g *gameHelper[T]) GetByName(name GameName) (*T, errorx.Error) {
	r, ok := g.games[name]
	if !ok {
		return nil, errorx.New("game not found")
	}
	return &r, nil
}

func (g *gameHelper[T]) GetByCode(code GameCode) (*T, errorx.Error) {
	r, ok := g.codes[code]
	if !ok {
		return nil, errorx.New("game not found")
	}
	return &r, nil
}
