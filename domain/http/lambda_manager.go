package http

type LambdaManager struct {
	Errors  interface{} `json:"errors"`
	Success bool        `json:"success"`
	Data    Data        `json:"data"`
}

type Data struct {
	AppName            string         `json:"app_name"`
	AppGroupID         string         `json:"app_group_id"`
	AppGroupName       string         `json:"app_group_name"`
	AppShortCode       string         `json:"app_short_code"`
	ServiceAccount     string         `json:"service_account"`
	TeamCode           string         `json:"team_code"`
	Pdgc               string         `json:"pdgc"`
	DataClassification string         `json:"data_classification"`
	MigrationCommand   interface{}    `json:"migration_command"`
	ResourceLabels     ResourceLabels `json:"resource_labels"`
	DeploymentTarget   string         `json:"deployment_target"`
	Region             string         `json:"region"`
	CommonNamespaceID  string         `json:"common_namespace_id"`
	Subnetwork         string         `json:"subnetwork"`
	CommonNamespace    string         `json:"common_namespace"`
	Team               string         `json:"team"`
	TeamID             string         `json:"team_id"`
	ProductGroupID     string         `json:"product_group_id"`
	AppID              string         `json:"app_id"`
	AppType            string         `json:"app_type"`
	Procs              []Procs        `json:"procs"`
	InfraHash          InfraHash      `json:"infra_hash"`
	DeploymentType     interface{}    `json:"deployment_type"`
	Members            []Members      `json:"members"`
}

type Members struct {
	Env Env `json:"env"`
}

type Env struct {
	ProdEnv EnvValues `json:"p"`
	IntEnv  EnvValues `json:"i"`
	DevEnv  EnvValues `json:"d"`
}

type EnvValues struct {
	TenantAdminsGroup string `json:"tenant_admins_group"`
	TenantUsersGroup  string `json:"tenant_users_group"`
}

type Procs struct {
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	InitCommand   string      `json:"init_command"`
	HealthCheck   interface{} `json:"health_check"`
	CustomRunlist string      `json:"custom_runlist"`
	Env           Env         `json:"env"`
}

type InfraHash struct{}

type ResourceLabels struct {
	Type                  string `json:"type"`
	AssetExposure         string `json:"asset-exposure"`
	IntegrityImpact       string `json:"integrity-impact"`
	AssetCriticality      string `json:"asset-criticality"`
	AvailabilityImpact    string `json:"availability-impact"`
	ConfidentialityImpact string `json:"confidentiality-impact"`
}
