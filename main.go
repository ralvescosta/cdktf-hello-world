package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	awsinstance "github.com/cdktf/cdktf-provider-aws-go/aws/v18/instance"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v18/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region: jsii.String("us-west-1"),
	})

	instance := awsinstance.NewInstance(stack, jsii.String("compute"), &awsinstance.InstanceConfig{
		Ami:          jsii.String("ami-01456a894f71116f2"),
		InstanceType: jsii.String("t2.micro"),
	})

	println(instance)

	    cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
	        Value: instance.PublicIp(),
	    })

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
