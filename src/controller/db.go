package controller

// Database operation return code
const (
	OperationSuccess  = 0
	OperationFailed   = -1
	CloseErr          = -11
	ConnectErr        = -10
	InsertFailed      = -100
	InsertSuccess     = 100
	InsertKeyExists   = -101
	QueryFailed       = -200
	GetKeyNotExist    = -201
	QueryRowNotExists = -202
	SessionExists     = -301
	TelExists         = -51
	EmailExists       = -52
	SQLiParameter     = -202
	PasswordInvalid   = -1000
)

// Databases
const (
	DbUsers    = 0
	DbEmail    = 1
	DbTel      = 2
	DbUserInfo = 3
	DbSession  = 4

	MysqlDB = "gingo:gingo@tcp(127.0.0.1:3306)/users?charset=utf8"
)
