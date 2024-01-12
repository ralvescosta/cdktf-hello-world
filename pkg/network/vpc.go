package network

import (
	"cdk.tf/go/stack/pkg/configs"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewVpc(cfgs *configs.Configs, tfStack cdktf.TerraformStack) (fnaVpc vpc.Vpc) {
	fnaVpc = vpc.NewVpc(tfStack, jsii.String(cfgs.Vpc.Name), &vpc.VpcConfig{
		CidrBlock: jsii.Sprintf(cfgs.Vpc.CidrBlock),
	})

	return
}
