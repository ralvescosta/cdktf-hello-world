package main

import (
	"github.com/ralvescosta/cdktf-hello-world/pkg"
	"github.com/ralvescosta/cdktf-hello-world/pkg/configs"
)

func main() {
	cfgs, logger := configs.NewConfigs()

	appScope, tfScope := pkg.NewAWSScopeProvider(logger, cfgs)
	pkg.ApplyStack(logger, cfgs, tfScope)

	appScope.Synth()
}
