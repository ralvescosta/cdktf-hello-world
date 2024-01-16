package network

import (
	"cdk.tf/go/stack/pkg/configs"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/eip"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/natgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewNatGateway(cfgs *configs.Configs, tfStack cdktf.TerraformStack, publicA subnet.Subnet, publicB subnet.Subnet) (
	natGatewayA natgateway.NatGateway,
	natGatewayB natgateway.NatGateway,
) {
	natGatewayA = natgateway.NewNatGateway(tfStack, jsii.String("nt-gtw-a"), &natgateway.NatGatewayConfig{
		SubnetId:         publicA.Id(),
		ConnectivityType: jsii.String("public"),
	})

	natGatewayB = natgateway.NewNatGateway(tfStack, jsii.String("nt-gtw-b"), &natgateway.NatGatewayConfig{
		SubnetId:         publicB.Id(),
		ConnectivityType: jsii.String("public"),
	})

	eip.NewEip(tfStack, jsii.String("fna-eip-a"), &eip.EipConfig{
		Domain:   jsii.String("vpc"),
		Instance: natGatewayA.Id(),
	})

	eip.NewEip(tfStack, jsii.String("fna-eip-b"), &eip.EipConfig{
		Domain:   jsii.String("vpc"),
		Instance: natGatewayB.Id(),
	})

	return
}
