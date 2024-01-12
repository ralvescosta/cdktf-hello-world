package network

import (
	"cdk.tf/go/stack/pkg/configs"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewSubnets(cfgs *configs.Configs, tfStack cdktf.TerraformStack, fnaVpc vpc.Vpc) (
	privateA subnet.Subnet, privateB subnet.Subnet, publicA subnet.Subnet, publicB subnet.Subnet,
) {
	privateA = subnet.NewSubnet(tfStack, jsii.Sprintf("fna-private-a"), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.String("10.0.1.0/24"),
		AvailabilityZone: jsii.String("us-west-1a"),
	})

	privateB = subnet.NewSubnet(tfStack, jsii.Sprintf("fna-private-b"), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.Sprintf("10.0.2.0/24"),
		AvailabilityZone: jsii.String("us-west-1c"),
	})

	publicA = subnet.NewSubnet(tfStack, jsii.Sprintf("fna-public-a"), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.Sprintf("10.0.3.0/24"),
		AvailabilityZone: jsii.String("us-west-1a"),
	})

	publicB = subnet.NewSubnet(tfStack, jsii.Sprintf("fna-public-b"), &subnet.SubnetConfig{
		VpcId:            fnaVpc.Id(),
		CidrBlock:        jsii.Sprintf("10.0.4.0/24"),
		AvailabilityZone: jsii.String("us-west-1c"),
	})

	return
}
