package pkg

import (
	"cdk.tf/go/stack/pkg/configs"
	"github.com/aws/jsii-runtime-go"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v18/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewAWSScopeProvider(cfgs *configs.Configs) (cdktf.App, cdktf.TerraformStack) {
	appScope := cdktf.NewApp(nil)

	tfScope := cdktf.NewTerraformStack(appScope, jsii.String(cfgs.AppId))

	awsprovider.NewAwsProvider(tfScope, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region:    jsii.String(cfgs.Region),
		AccessKey: jsii.String(cfgs.AccessKey),
		SecretKey: jsii.String(cfgs.SecretKey),
	})

	cdktf.NewCloudBackend(tfScope, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("ralvescostait"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String(cfgs.AppId)),
	})

	return appScope, tfScope
}
