package network

import (
	"cdk.tf/go/stack/pkg/configs"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/internetgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewInternetGateway(cfgs *configs.Configs, tfStack cdktf.TerraformStack, fnaVpc vpc.Vpc) (igw internetgateway.InternetGateway) {
	igw = internetgateway.NewInternetGateway(tfStack, jsii.String("fna-igw"), &internetgateway.InternetGatewayConfig{
		VpcId: fnaVpc.Id(),
	})

	return
}
