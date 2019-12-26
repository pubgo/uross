package cmds

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/uross/cmds/uross"
)

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	defer xerror.Assert()

	cmd.AddCommand(
		uross.Server(),
	)
})
