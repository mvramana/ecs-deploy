package api

type Flags struct {
	Bootstrap             bool
	DisableEcsDeploy      bool
	Profile               string
	Region                string
	ClusterName           string
	Environment           string
	AlbSecurityGroups     string
	EcsSubnets            string
	CloudwatchLogsPrefix  string
	CloudwatchLogsEnabled bool
	KeyName               string
	InstanceType          string
	EcsSecurityGroups     string
	EcsMinSize            string
	EcsMaxSize            string
	EcsDesiredSize        string
	ParamstoreEnabled     bool
	ParamstoreKmsArn      string
	ParamstorePrefix      string
	LoadbalancerDomain    string
	Server                bool
	DeleteCluster         string
}

func NewFlags() *Flags {
	f := Flags{}
	return &f
}
