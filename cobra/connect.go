package cobra

import (
	"errors"
	"sync"

	"github.com/mono83/radio"
	"github.com/mono83/radio/std"
	"github.com/spf13/cobra"
)

// Connect produces connector, that will be able to perform operations
func Connect(cmd *cobra.Command) (radio.Connector, error) {
	if cmd == nil {
		return nil, errors.New("no cobra command given")
	}
	connector := connector{
		cmd: cmd,
	}

	return &connector, nil
}

type connector struct {
	cmd  *cobra.Command
	lock sync.Mutex
}

// Invoke invokes command inside context
func (c *connector) Invoke(ctx radio.Context) error {
	if ctx == nil {
		return errors.New("no context given")
	}

	// Obtaining exclusive lock on command
	c.lock.Lock()
	defer c.lock.Unlock()

	// Sending indicator, that command in progress
	ctx.CommandInProgress()

	// Running with os.Std replacements
	var err error
	sout, serr := std.BindStrings(func() {
		c.cmd.SetArgs(ctx.GetArgs())
		err = c.cmd.Execute()
	})

	// Packing response
	if len(serr) > 0 {
		ctx.SendMessage(serr)
	}
	if len(sout) > 0 {
		ctx.SendMessage(sout)
	}
	if err != nil {
		ctx.SendMessage(err)
	}

	return err
}
