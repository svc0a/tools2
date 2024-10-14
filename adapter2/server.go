package adapter2

import (
	"bitbucket.org/alibaba-international/ghs-common/adapter"
	"bitbucket.org/alibaba-international/go-pkg/ctxx"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
)

func NewServer(providerName string, gameGameLink getGameLinkFunc, checkOrderState checkOrderStateFunc) adapter.Server {
	server1 := adapter.NewServer(providerName, map[adapter.Environment]adapter.Client{
		adapter.EnvTest: newClient(adapter.EnvTest, gameGameLink, checkOrderState),
		adapter.EnvProd: newClient(adapter.EnvProd, gameGameLink, checkOrderState),
	})
	return server1
}

type getGameLinkFunc func(ctx ctxx.Context, cmd adapter.EnterGameCommand, env adapter.Environment) (*adapter.EnterGameResult, errorx.Error)
type checkOrderStateFunc func(ctx ctxx.Context, cmd adapter.CheckOrderStateCommand, env adapter.Environment) (*adapter.CheckOrderStateResult, errorx.Error)

type client struct {
	env             adapter.Environment
	getGameLink     getGameLinkFunc
	checkOrderState checkOrderStateFunc
}

func newClient(env adapter.Environment, linkFunc getGameLinkFunc, stateFunc checkOrderStateFunc) *client {
	return &client{
		env:             env,
		getGameLink:     linkFunc,
		checkOrderState: stateFunc,
	}
}

func (c *client) GetGameLink(ctx ctxx.Context, cmd adapter.EnterGameCommand) (*adapter.EnterGameResult, errorx.Error) {
	return c.getGameLink(ctx, cmd, c.env)
}

func (c *client) CheckOrderState(ctx ctxx.Context, cmd adapter.CheckOrderStateCommand) (*adapter.CheckOrderStateResult, errorx.Error) {
	return c.checkOrderState(ctx, cmd, c.env)
}
