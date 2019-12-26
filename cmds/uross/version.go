package uross

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xenv"
)

func Version() *xcmd.Command {
	return &xcmd.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "version info",
		Run: func(cmd *xcmd.Command, args []string) {
			xenv.Version()
		},
	}
}
