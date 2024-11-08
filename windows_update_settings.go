package openuem_nats

const NOT_CONFIGURED = "windowsupdate.not_configured"
const DISABLED = "windowsupdate.disabled"
const NOTIFY_BEFORE_DOWNLOAD = "windowsupdate.notify_before_download"
const NOTIFY_BEFORE_INSTALLATION = "windowsupdate.notify_before_installation"
const NOTIFY_SCHEDULED_INSTALLATION = "windowsupdate.notify_scheduled_installation"

func WindowsUpdatePossibleStatus() []string {
	return []string{NOT_CONFIGURED, DISABLED, NOTIFY_BEFORE_DOWNLOAD, NOTIFY_BEFORE_INSTALLATION, NOTIFY_SCHEDULED_INSTALLATION}
}
