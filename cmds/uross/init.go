package uross

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/uross/version"
)

func init() {
	xenv.InitVersion("Uross", version.Version, version.BuildV, version.CommitV)
}
