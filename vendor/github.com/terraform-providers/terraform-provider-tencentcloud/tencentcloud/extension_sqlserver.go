package tencentcloud

const (
	SQLSERVER_CHARGE_TYPE_PREPAID  = "PREPAID"
	SQLSERVER_CHARGE_TYPE_POSTPAID = "POSTPAID_BY_HOUR"
)

var SQLSERVER_CHARGE_TYPE_NAME = map[string]string{
	"PRE":  SQLSERVER_CHARGE_TYPE_PREPAID,
	"POST": SQLSERVER_CHARGE_TYPE_POSTPAID,
	"ALL":  "ALL",
}
