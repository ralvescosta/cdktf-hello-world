package main

import (
	"cdk.tf/go/stack/pkg"
	"cdk.tf/go/stack/pkg/configs"
)

func main() {
	cfgs := configs.NewConfigs()

	appScope, tfScope := pkg.NewAWSScopeProvider(cfgs)

	pkg.ApplyStack(cfgs, tfScope)

	appScope.Synth()
}
