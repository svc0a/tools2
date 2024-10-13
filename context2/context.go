package context2

import (
	"bitbucket.org/alibaba-international/ghs-common/adapter"
	"bitbucket.org/alibaba-international/go-pkg/ctxx"
	"bitbucket.org/alibaba-international/go-pkg/types"
	"bitbucket.org/alibaba-international/go-pkg/types/locale"
	"context"
	"fmt"
	"time"
)

func ToAny[T any](in T) *any {
	var result any = in
	return &result
}

func GetContext(provider string, environment adapter.Environment) ctxx.Context {
	return ctxx.WithMetadata(context.Background(), &ctxx.Metadata{
		TraceID:   types.NewID(),
		Operator:  types.NewID().String(),
		Caller:    fmt.Sprintf("%s.%s", provider, environment),
		Timestamp: time.Now().UnixMilli(),
		Locale:    locale.EnUS,
	})
}
