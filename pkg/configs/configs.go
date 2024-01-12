package configs

type Configs struct {
	Provider *ProviderConfigs

	Vpc *VpcConfigs

	PrivateSubnetA *SubnetConfigs
	PrivateSubnetB *SubnetConfigs
	PublicSubnetA  *SubnetConfigs
	PublicSubnetB  *SubnetConfigs

	InternetGateway *InternetGatewayConfigs

	PrivateRouteTable *RouteTableConfigs
	PublicRouteTable  *RouteTableConfigs
}

type ProviderConfigs struct {
	AppId                    string
	Region                   string
	AccessKey                string
	SecretKey                string
	CloudBackendHostname     string
	CloudBackendOrganization string
}

type VpcConfigs struct {
	Name      string
	CidrBlock string
}

type SubnetConfigs struct {
	Name             string
	CidrBlock        string
	AvailabilityZone string
}

type InternetGatewayConfigs struct {
	Name string
}

type RouteTableConfigs struct {
	Name                   string
	SubnetAssociationNames []string
}

func NewConfigs() *Configs {
	return &Configs{
		Provider: &ProviderConfigs{
			AppId:                    "cdktf-hello-world",
			Region:                   "us-west-1",
			AccessKey:                "",
			SecretKey:                "",
			CloudBackendHostname:     "app.terraform.io",
			CloudBackendOrganization: "ralvescostait",
		},
		Vpc: &VpcConfigs{
			Name:      "fna-vpc",
			CidrBlock: "10.0.0.0/16",
		},
		PrivateSubnetA: &SubnetConfigs{
			Name:             "fna-private-a",
			CidrBlock:        "10.0.1.0/24",
			AvailabilityZone: "us-west-1a",
		},
		PrivateSubnetB: &SubnetConfigs{
			Name:             "fna-private-b",
			CidrBlock:        "10.0.2.0/24",
			AvailabilityZone: "us-west-1c",
		},
		PublicSubnetA: &SubnetConfigs{
			Name:             "fna-public-a",
			CidrBlock:        "10.0.3.0/24",
			AvailabilityZone: "us-west-1a",
		},
		PublicSubnetB: &SubnetConfigs{
			Name:             "fna-public-b",
			CidrBlock:        "10.0.4.0/24",
			AvailabilityZone: "us-west-1c",
		},
		InternetGateway: &InternetGatewayConfigs{
			Name: "fna-igw",
		},
		PrivateRouteTable: &RouteTableConfigs{
			Name:                   "fna-private-rt",
			SubnetAssociationNames: []string{"fna-private-rt-a", "fna-private-rt-b"},
		},
		PublicRouteTable: &RouteTableConfigs{
			Name:                   "fna-public-rt",
			SubnetAssociationNames: []string{"fna-public-rt-a", "fna-public-rt-b"},
		},
	}
}
