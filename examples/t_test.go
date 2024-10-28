package examples

import (
	"github.com/pkg/errors"
	"github.com/svc0a/tools2/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.Init()
	err := outer()
	logger.Define().Error(err)
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}
