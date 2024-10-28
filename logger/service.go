package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type Service interface {
	Error(err error)
}

type impl struct{}

func (i impl) Error(err error) {
	log.Error().Stack().Err(err).Msg(err.Error())
}

func Define() Service {
	return &impl{}
}

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.With().CallerWithSkipFrameCount(3).Logger()
}
