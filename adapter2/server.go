package adapter2

import (
	"bitbucket.org/alibaba-international/ghs-common/adapter"
	"bitbucket.org/alibaba-international/ghs-common/model"
	"bitbucket.org/alibaba-international/go-pkg/ctxx"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
)

type Servers map[model.Environment]adapter.Server

func NewServer(providerName string, gameGameLink getGameLinkFunc, stateFunc getBetStateByIDFunc) Servers {
	testClient := newClient(model.EnvTest, gameGameLink, stateFunc)
	testServer := adapter.NewServer(model.EnvTest, providerName, testClient)
	testClient.Server = testServer
	prodClient := newClient(model.EnvProd, gameGameLink, stateFunc)
	prodServer := adapter.NewServer(model.EnvProd, providerName, prodClient)
	prodClient.Server = prodServer
	return map[model.Environment]adapter.Server{
		model.EnvTest: testServer,
		model.EnvProd: prodServer,
	}
}

type getGameLinkFunc func(ctx ctxx.Context, cmd adapter.EnterGameCommand, env model.Environment, server adapter.Server) (*adapter.EnterGameResult, errorx.Error)
type getBetStateByIDFunc func(ctx ctxx.Context, betID string, env model.Environment) (*adapter.BetState, errorx.Error)

type client struct {
	env                 model.Environment
	getGameLink         getGameLinkFunc
	getBetStateByIDFunc getBetStateByIDFunc
	adapter.Server
}

func newClient(env model.Environment, linkFunc getGameLinkFunc, stateFunc getBetStateByIDFunc) *client {
	return &client{
		env:                 env,
		getGameLink:         linkFunc,
		getBetStateByIDFunc: stateFunc,
	}
}

func (c *client) GetGameLink(ctx ctxx.Context, cmd adapter.EnterGameCommand) (*adapter.EnterGameResult, errorx.Error) {
	return c.getGameLink(ctx, cmd, c.env, c.Server)
}

func (c *client) GetBetStateByID(ctx ctxx.Context, betID string) (*adapter.BetState, errorx.Error) {
	return c.getBetStateByIDFunc(ctx, betID, c.env)
}
