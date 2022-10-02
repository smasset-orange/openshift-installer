package flexibleengine

import (
	"encoding/json"
)

type config struct {
	AccessKey   string            `json:"fe_access_key,omitempty"`
	SecretKey   string            `json:"fe_secret_key,omitempty"`
	DomainName  string            `json:"fe_domain_name"`
	Region      string            `json:"fe_region"`
	ProjectName string            `json:"fe_project_name,omitempty"`
	CommonTags  map[string]string `json:"fe_common_tags,omitempty"`
}

// TFVarsSources contains the parameters to be converted into Terraform variables
type TFVarsSources struct {
	AccessKey   string
	SecretKey   string
	DomainName  string
	Region      string
	ProjectName string
	CommonTags  map[string]string
}

// TFVars generates Azure-specific Terraform variables launching the cluster.
func TFVars(sources TFVarsSources) ([]byte, error) {
	cfg := &config{
		AccessKey:   sources.AccessKey,
		SecretKey:   sources.SecretKey,
		DomainName:  sources.DomainName,
		Region:      sources.Region,
		ProjectName: sources.ProjectName,
		CommonTags:  sources.CommonTags,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
