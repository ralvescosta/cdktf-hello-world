package network

import (
	"cdk.tf/go/stack/pkg/configs"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/internetgateway"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetableassociation"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewRouteTables(
	cfgs *configs.Configs,
	tfStack cdktf.TerraformStack,
	fnaVpc vpc.Vpc,
	igw internetgateway.InternetGateway,
	privateA subnet.Subnet,
	privateB subnet.Subnet,
	publicA subnet.Subnet,
	publicB subnet.Subnet,
) (privateRouteTable routetable.RouteTable, publicRouteTable routetable.RouteTable) {
	privateRouteTable = routetable.NewRouteTable(tfStack, jsii.String(cfgs.PrivateRouteTable.Name), &routetable.RouteTableConfig{
		VpcId: fnaVpc.Id(),
	})

	publicRouteTable = routetable.NewRouteTable(tfStack, jsii.String(cfgs.PublicRouteTable.Name), &routetable.RouteTableConfig{
		VpcId: fnaVpc.Id(),
		Route: []*routetable.RouteTableRoute{
			{
				CidrBlock: jsii.String("0.0.0.0/0"),
				GatewayId: igw.Id(),
			},
		},
	})

	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PrivateRouteTable.SubnetAssociationNames[0]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: privateRouteTable.Id(),
		SubnetId:     privateA.Id(),
	})
	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PrivateRouteTable.SubnetAssociationNames[1]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: privateRouteTable.Id(),
		SubnetId:     privateB.Id(),
	})

	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PublicRouteTable.SubnetAssociationNames[0]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: publicRouteTable.Id(),
		SubnetId:     publicA.Id(),
	})
	routetableassociation.NewRouteTableAssociation(tfStack, jsii.String(cfgs.PublicRouteTable.SubnetAssociationNames[1]), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: publicRouteTable.Id(),
		SubnetId:     publicB.Id(),
	})

	return
}
