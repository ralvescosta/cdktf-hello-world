package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/internetgateway"
	"cdk.tf/go/stack/generated/hashicorp/aws/routetableassociation"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/internetgateway"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v18/provider"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/routetable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/subnet"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region: jsii.String("us-west-1"),
	})

	vpc := vpc.NewVpc(stack, jsii.String("fna-vpc"), &vpc.VpcConfig{
		CidrBlock: jsii.Sprintf("10.0.0.0/16"),
	})

	privatea := subnet.NewSubnet(stack, jsii.Sprintf("fna-private-a"), &subnet.SubnetConfig{
		VpcId:            vpc.Id(),
		CidrBlock:        jsii.String("10.1.0.0/24"),
		AvailabilityZone: jsii.String("us-west-1a"),
	})

	privateb := subnet.NewSubnet(stack, jsii.Sprintf("fna-private-a"), &subnet.SubnetConfig{
		VpcId:            vpc.Id(),
		CidrBlock:        jsii.Sprintf("10.2.0.0/24"),
		AvailabilityZone: jsii.String("us-west-1b"),
	})

	publica := subnet.NewSubnet(stack, jsii.Sprintf("fna-private-a"), &subnet.SubnetConfig{
		VpcId:            vpc.Id(),
		CidrBlock:        jsii.Sprintf("10.3.0.0/24"),
		AvailabilityZone: jsii.String("us-west-1a"),
	})

	publicb := subnet.NewSubnet(stack, jsii.Sprintf("fna-private-a"), &subnet.SubnetConfig{
		VpcId:            vpc.Id(),
		CidrBlock:        jsii.Sprintf("10.4.0.0/24"),
		AvailabilityZone: jsii.String("us-west-1b"),
	})

	igw := internetgateway.NewInternetGateway(stack, jsii.String("fna-igw"), &internetgateway.InternetGatewayConfig{
		VpcId: vpc.Id(),
	})

	privateRTable := routetable.NewRouteTable(stack, jsii.String("fna-private-rt"), &routetable.RouteTableConfig{
		VpcId: vpc.Id(),
	})

	publicRTable := routetable.NewRouteTable(stack, jsii.String("fna-public-rt"), &routetable.RouteTableConfig{
		VpcId: vpc.Id(),
		Route: routetable.RouteTableRoute{
			CidrBlock: jsii.String("0.0.0.0/0"),
			GatewayId: igw.Id(),
		},
	})

	routetableassociation.NewRouteTableAssociation(stack, jsii.String("fna-private-rt-a"), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: privateRTable.Id(),
		SubnetId:     privatea.Id(),
	})
	routetableassociation.NewRouteTableAssociation(stack, jsii.String("fna-private-rt-b"), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: privateRTable.Id(),
		SubnetId:     privateb.Id(),
	})

	routetableassociation.NewRouteTableAssociation(stack, jsii.String("fna-public-rt-a"), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: publicRTable.Id(),
		SubnetId:     publica.Id(),
	})
	routetableassociation.NewRouteTableAssociation(stack, jsii.String("fna-public-rt-b"), &routetableassociation.RouteTableAssociationConfig{
		RouteTableId: publicRTable.Id(),
		SubnetId:     publicb.Id(),
	})

	// instance := awsinstance.NewInstance(stack, jsii.String("compute"), &awsinstance.InstanceConfig{
	// 	Ami:          jsii.String("ami-01456a894f71116f2"),
	// 	InstanceType: jsii.String("t2.micro"),
	// })

	// cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
	// 	Value: instance.PublicIp(),
	// })

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "cdktf-hello-world")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("ralvescostait"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("cdktf-hello-world")),
	})

	app.Synth()
}
