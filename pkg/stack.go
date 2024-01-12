package pkg

import (
	"cdk.tf/go/stack/pkg/configs"
	"cdk.tf/go/stack/pkg/network"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func ApplyStack(cfgs *configs.Configs, tfStack cdktf.TerraformStack) {
	fnaVpc := network.NewVpc(cfgs, tfStack)
	privateA, privateB, publicA, publicB := network.NewSubnets(cfgs, tfStack, fnaVpc)
	igw := network.NewInternetGateway(cfgs, tfStack, fnaVpc)
	_, _ = network.NewRouteTables(cfgs, tfStack, fnaVpc, igw, privateA, privateB, publicA, publicB)
}
