package nats

type WingetPackage struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"name,omitempty"`
}

type WingetCfgProfiles struct {
	AgentID string `json:"agentID,omitempty"`
}

type WingetCfgReport struct {
	ProfileID string `json:"profileID,omitempty"`
	Success   bool   `json:"success,omitempty"`
}

type WingetCfgDeploy struct {
	PackageID string `json:"PackageID,omitempty"`
	Installed bool   `json:"installed,omitempty"`
}
