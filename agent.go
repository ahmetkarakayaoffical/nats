package openuem_nats

import "time"

type Computer struct {
	Manufacturer   string `json:"manufacturer,omitempty"`
	Model          string `json:"model,omitempty"`
	Serial         string `json:"serial,omitempty"`
	Processor      string `json:"processor,omitempty"`
	ProcessorArch  string `json:"processor_arch,omitempty"`
	ProcessorCores int64  `json:"processor_cores,omitempty"`
	Memory         uint64 `json:"memory,omitempty"`
}

type Antivirus struct {
	Name      string `json:"name,omitempty"`
	IsActive  bool   `json:"is_active,omitempty"`
	IsUpdated bool   `json:"is_updated,omitempty"`
}

type OperatingSystem struct {
	Version        string    `json:"version,omitempty"`
	Description    string    `json:"description,omitempty"`
	InstallDate    time.Time `json:"install_date,omitempty"`
	Edition        string    `json:"edition,omitempty"`
	Arch           string    `json:"arch,omitempty"`
	Username       string    `json:"username,omitempty"`
	LastBootUpTime time.Time `json:"last_bootup_time,omitempty"`
}

type LogicalDisk struct {
	Label                 string `json:"label,omitempty"`
	Usage                 int8   `json:"usage,omitempty"`
	Filesystem            string `json:"filesystem,omitempty"`
	SizeInUnits           string `json:"size_in_units,omitempty"`
	RemainingSpaceInUnits string `json:"remaining_space_in_units,omitempty"`
	VolumeName            string `json:"volume_name,omitempty"`
	BitLockerStatus       string `json:"bitlocker_status,omitempty"`
}

type Monitor struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
	Serial       string `json:"serial,omitempty"`
}

type Printer struct {
	Name      string `json:"name,omitempty"`
	Port      string `json:"port,omitempty"`
	IsDefault bool   `json:"is_default,omitempty"`
	IsNetwork bool   `json:"is_network,omitempty"`
}

type Share struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Path        string `json:"path,omitempty"`
}

type SystemUpdate struct {
	Status         string    `json:"status,omitempty"`
	LastInstall    time.Time `json:"last_install,omitempty"`
	LastSearch     time.Time `json:"last_search,omitempty"`
	PendingUpdates bool      `json:"pending_updates,omitempty"`
}

type NetworkAdapter struct {
	Name              string    `json:"name,omitempty"`
	MACAddress        string    `json:"mac_address,omitempty"`
	Addresses         string    `json:"addresses,omitempty"`
	Subnet            string    `json:"subnet,omitempty"`
	DefaultGateway    string    `json:"default_gateway,omitempty"`
	DNSServers        string    `json:"dns_servers,omitempty"`
	DNSDomain         string    `json:"dns_domain,omitempty"`
	DHCPEnabled       bool      `json:"dhcp_enabled,omitempty"`
	DHCPLeaseObtained time.Time `json:"dhcp_lease_obtained,omitempty"`
	DHCPLeaseExpired  time.Time `json:"dhcp_lease_expired,omitempty"`
	Speed             string    `json:"speed,omitempty"`
}

type Application struct {
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	InstallDate string `json:"install_date,omitempty"`
	Publisher   string `json:"publisher,omitempty"`
}

type Update struct {
	Title      string    `json:"title,omitempty"`
	Date       time.Time `json:"date,omitempty"`
	SupportURL string    `json:"support_url,omitempty"`
}

type LoggedOnUser struct {
	Name      string    `json:"name,omitempty"`
	LastLogon time.Time `json:"last_logon,omitempty"`
}

type AgentReport struct {
	AgentID            string           `json:"id,omitempty"`
	OS                 string           `json:"os,omitempty"`
	Hostname           string           `json:"hostname,omitempty"`
	Version            string           `json:"version,omitempty"`
	ExecutionTime      time.Time        `json:"execution_time,omitempty"`
	Enabled            bool             `json:"enable,omitempty"`
	IP                 string           `json:"ip,omitempty"`
	MACAddress         string           `json:"mac,omitempty"`
	Computer           Computer         `json:"computer,omitempty"`
	Antivirus          Antivirus        `json:"antivirus,omitempty"`
	OperatingSystem    OperatingSystem  `json:"operatingsystem,omitempty"`
	LogicalDisks       []LogicalDisk    `json:"logicaldisks,omitempty"`
	Monitors           []Monitor        `json:"monitors,omitempty"`
	Printers           []Printer        `json:"printers,omitempty"`
	Shares             []Share          `json:"shares,omitempty"`
	SystemUpdate       SystemUpdate     `json:"systemupdate,omitempty"`
	NetworkAdapters    []NetworkAdapter `json:"networkadapters,omitempty"`
	Applications       []Application    `json:"apps,omitempty"`
	LoggedOnUsers      []LoggedOnUser   `json:"loggedonusers,omitempty"`
	SupportedVNCServer string           `json:"vnc,omitempty"`
	Updates            []Update         `json:"updates,omitempty"`
}
