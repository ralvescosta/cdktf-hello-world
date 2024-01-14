package ecs

import (
	"cdk.tf/go/stack/pkg/configs"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecscluster"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecsservice"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecstaskdefinition"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewECSFargate(cfgs *configs.Configs, tfStack cdktf.TerraformStack, fnaVpc vpc.Vpc) {
	cluster := ecscluster.NewEcsCluster(tfStack, jsii.String("fna-ecs-cluster"), &ecscluster.EcsClusterConfig{
		Setting: []*ecscluster.EcsClusterSetting{
			{
				Name:  jsii.String("containerInsights"),
				Value: jsii.String("enabled"),
			},
		},
	})

	td := ecstaskdefinition.NewEcsTaskDefinition(tfStack, jsii.String("fna-td"), &ecstaskdefinition.EcsTaskDefinitionConfig{
		Family:                  jsii.String("service"),
		Cpu:                     jsii.String("0.5"),
		Memory:                  jsii.String("128M"),
		NetworkMode:             jsii.String("awsvpc"),
		RequiresCompatibilities: jsii.Strings("FARGATE"),
		ContainerDefinitions: jsii.String(`
		<<TASK_DEFINITION
		[
			{
				"image": "nginx",
				"name": "nginx",
				"portMappings": [
					{
						"containerPort": 80,
					}
				],
			}
		]
		TASK_DEFINITION
		`),
	})

	ecsservice.NewEcsService(tfStack, jsii.String("fna-service"), &ecsservice.EcsServiceConfig{
		Cluster:        cluster.Id(),
		TaskDefinition: td.Arn(),
	})
}
