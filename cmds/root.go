package cmds

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/uross/version"

	// 导入db migrations
	//_ "github.com/pubgo/g/migrations"
	"github.com/pubgo/g/xcmd"
)

const Service = "uross"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	defer xerror.Assert()

	cmd.Use = Service
	cmd.Version = version.Version

	cmd.AddCommand()
})

func init() {
	xenv.InitVersion("Uross", version.Version, version.BuildV, version.CommitV)
}
