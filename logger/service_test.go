package logger

import (
	"github.com/pkg/errors"
	"testing"
)

func Test1(t *testing.T) {
	Init()
	err := outer()
	Define().Error(err)
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
