package controller

// Operation return code
const (
	// Basic Operation |err| < 10
	OperationSuccess = 0
	OperationFailed  = -1
	// DB Connection  10 < |err| < 20
	ConnectErr = -11
	CloseErr   = -12
	// Users Basic Info 50 < |err| < 60
	TelExists       = -51
	EmailExists     = -52
	PasswordInvalid = -59
	PasswordCorrect = 59
	// User Advance Info 60 < |err| < 70
	LackOfAdvInfo = 61
	// Session 70 < |err| < 80
	SessionExpired = 71
	// DB Insertion 90 < |err| < 100
	InsertFailed  = -99
	InsertSuccess = 99
	// DB Query 100 < |err| < 110
	QueryFailed       = -101
	QueryRowNotExists = -102
	// DB Security 900 < |err| < 1000
	SQLiParameter = -901
)

// Databases
const (
	DbSession = 0
	MysqlDB   = "gingo:gingo@tcp(127.0.0.1:3306)/users?charset=utf8"
)
