package tencentcloud

const (
	CBS_STORAGE_TYPE_CLOUD_BASIC   = "CLOUD_BASIC"
	CBS_STORAGE_TYPE_CLOUD_PREMIUM = "CLOUD_PREMIUM"
	CBS_STORAGE_TYPE_CLOUD_SSD     = "CLOUD_SSD"

	CBS_STORAGE_USAGE_SYSTEM_DISK = "SYSTEM_DISK"
	CBS_STORAGE_USAGE_DATA_DISK   = "DATA_DISK"

	CBS_STORAGE_STATUS_UNATTACHED  = "UNATTACHED"
	CBS_STORAGE_STATUS_ATTACHING   = "ATTACHING"
	CBS_STORAGE_STATUS_ATTACHED    = "ATTACHED"
	CBS_STORAGE_STATUS_EXPANDING   = "EXPANDING"
	CBS_STORAGE_STATUS_ROLLBACKING = "ROLLBACKING"
	CBS_STORAGE_STATUS_TORECYCLE   = "TORECYCLE"

	CBS_SNAPSHOT_STATUS_NORMAL   = "NORMAL"
	CBS_SNAPSHOT_STATUS_CREATING = "CREATING"
)

var CBS_STORAGE_TYPE = []string{
	CBS_STORAGE_TYPE_CLOUD_BASIC,
	CBS_STORAGE_TYPE_CLOUD_PREMIUM,
	CBS_STORAGE_TYPE_CLOUD_SSD,
}

var CBS_STORAGE_USAGE = []string{
	CBS_STORAGE_USAGE_SYSTEM_DISK,
	CBS_STORAGE_USAGE_DATA_DISK,
}

const (
	CBS_PREPAID_RENEW_FLAG_NOTIFY_NOTIFY_AND_AUTO_RENEW    = "NOTIFY_AND_AUTO_RENEW"
	CBS_PREPAID_RENEW_FLAG_NOTIFY_AND_MANUAL_RENEW         = "NOTIFY_AND_MANUAL_RENEW"
	CBS_PREPAID_RENEW_FLAG_DISABLE_NOTIFY_AND_MANUAL_RENEW = "DISABLE_NOTIFY_AND_MANUAL_RENEW"
	CBS_CHARGE_TYPE_PREPAID                                = "PREPAID"
	CBS_CHARGE_TYPE_POSTPAID                               = "POSTPAID_BY_HOUR"
)

var CBS_PREPAID_RENEW_FLAG = []string{
	CBS_PREPAID_RENEW_FLAG_NOTIFY_NOTIFY_AND_AUTO_RENEW,
	CBS_PREPAID_RENEW_FLAG_NOTIFY_AND_MANUAL_RENEW,
	CBS_PREPAID_RENEW_FLAG_DISABLE_NOTIFY_AND_MANUAL_RENEW,
}

var CBS_CHARGE_TYPE = []string{
	CBS_CHARGE_TYPE_PREPAID,
	CBS_CHARGE_TYPE_POSTPAID,
}

var CBS_PREPAID_PERIOD = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36}