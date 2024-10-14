package adapter2

import (
	"bitbucket.org/alibaba-international/ghs-common/adapter"
	"bitbucket.org/alibaba-international/ghs-common/model"
	"bitbucket.org/alibaba-international/go-pkg/ctxx"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
)

type Servers map[model.Environment]adapter.Server

func NewServer(providerName string, gameGameLink getGameLinkFunc, stateFunc getBetStateByIDFunc) Servers {
	return map[model.Environment]adapter.Server{
		model.EnvTest: adapter.NewServer(model.EnvTest, providerName, newClient(model.EnvTest, gameGameLink, stateFunc)),
		model.EnvProd: adapter.NewServer(model.EnvProd, providerName, newClient(model.EnvProd, gameGameLink, stateFunc)),
	}
}

type getGameLinkFunc func(ctx ctxx.Context, cmd adapter.EnterGameCommand, env model.Environment) (*adapter.EnterGameResult, errorx.Error)
type getBetStateByIDFunc func(ctx ctxx.Context, betID string, env model.Environment) (*adapter.BetState, errorx.Error)

type client struct {
	env                 model.Environment
	getGameLink         getGameLinkFunc
	getBetStateByIDFunc getBetStateByIDFunc
}

func newClient(env model.Environment, linkFunc getGameLinkFunc, stateFunc getBetStateByIDFunc) *client {
	return &client{
		env:                 env,
		getGameLink:         linkFunc,
		getBetStateByIDFunc: stateFunc,
	}
}

func (c *client) GetGameLink(ctx ctxx.Context, cmd adapter.EnterGameCommand) (*adapter.EnterGameResult, errorx.Error) {
	return c.getGameLink(ctx, cmd, c.env)
}

func (c *client) GetBetStateByID(ctx ctxx.Context, betID string) (*adapter.BetState, errorx.Error) {
	return c.getBetStateByIDFunc(ctx, betID, c.env)
}
