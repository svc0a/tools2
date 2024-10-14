package context2

import (
	"bitbucket.org/alibaba-international/ghs-common/model"
	"bitbucket.org/alibaba-international/go-pkg/ctxx"
	"bitbucket.org/alibaba-international/go-pkg/errorx"
	"bitbucket.org/alibaba-international/go-pkg/rest/api"
	"bitbucket.org/alibaba-international/go-pkg/types"
	"bitbucket.org/alibaba-international/go-pkg/types/locale"
	"context"
	"fmt"
	"github.com/svc0a/tools2/error2"
	"time"
)

func ToAny[T any](in T) *any {
	var result any = in
	return &result
}

func GetContext(provider string, environment *model.Environment) ctxx.Context {
	return ctxx.WithMetadata(context.Background(), &ctxx.Metadata{
		TraceID:   types.NewID(),
		Operator:  types.NewID().String(),
		Caller:    fmt.Sprintf("%s.%s", provider, *environment),
		Timestamp: time.Now().UnixMilli(),
		Locale:    locale.EnUS,
	})
}

func GetEnvFromContext(ctx *api.Context) (*model.Environment, errorx.Error) {
	value, exists := ctx.Get("env")
	if !exists {
		return nil, error2.ErrEnvNotFound
	}
	env1, ok := value.(model.Environment)
	if !ok {
		return nil, error2.ErrEnvNotFound
	}
	return &env1, nil
}
