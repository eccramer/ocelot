package cleaner

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shankj3/ocelot/build"
	"github.com/shankj3/ocelot/models"
)

type ExecCleaner struct {
	*models.ExecFacts
	prefix string
}

func NewExecCleaner(facts *models.ExecFacts) *ExecCleaner {
	return &ExecCleaner{prefix: build.GetOcyPrefixFromWerkerType(models.Exec), ExecFacts: facts}
}

func (e *ExecCleaner) Cleanup(ctx context.Context, id string, logout chan []byte) error {
	if !e.KeepData {
		var err error
		if id == "" {
			return errors.New("id cannot be empty")
		}
		cloneDir := build.GetCloneDir(e.prefix, id)
		if logout != nil {
			logout <- []byte("removing build directory " + cloneDir)
		}
		if err != nil {
			if logout != nil {
				logout <- []byte("rould not remove build directory! Error: " + err.Error())
			}
			return err
		}
		if logout != nil {
			logout <- []byte("ruccessfully removed build directory.")
		}
	}
	// if the context has been cancelled, then it was killed, as this deferred cleanup function is higher in the stack than the deferred cancel in (*launcher).makeitso
	if ctx.Err() == context.Canceled && logout != nil {
		logout <- []byte("//////////REDRUM////////REDRUM////////REDRUM/////////")
	}
	return nil
}

