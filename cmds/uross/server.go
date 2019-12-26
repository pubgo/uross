package uross

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xdi"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/uross/uross/app"
)

func Server() *xcmd.Command {
	var (
		args = xcmd.Args(func(cmd *xcmd.Command) {

		})
	)

	return args(&xcmd.Command{
		Use:   "s",
		Short: "server",
		RunE: func(cmd *xcmd.Command, args []string) (err error) {
			defer xerror.RespErr(&err)
			xerror.Panic(xdi.Invoke(func(ctrl app.App) {
				ctrl.InitRoutes()
			}))
			return
		},
	})
}
