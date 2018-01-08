package cmd

import (
	"queue/core"
	"queue/model"
)

type Cmd struct {
}

func (cmd *Cmd) Run() {
	core.Push(model.Job{})
}
