package dingTalk

const (
	StandardDateFormat = "2006-01-02 15:04:05"
)

const (
	Format                 = "json"
	AlphabetsPool          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	DefaultRandomNum       = 16
	PKCS7EncoderBlockSize  = 32
	OK                     = 0
	ValidateSignatureError = 900005
	ComputeSignatureError  = 900006
	EncryptAESError        = 900007
	DecryptAESError        = 900008
	ValidateSuiteKeyError  = 900010
)

// Http头标签
const (
	HTTPHeaderContentLength             = "Content-Length"
	HTTPHeaderContentType               = "Content-Type"
	HTTPHeaderDate                      = "Date"
)

const (
	ContentTypeJson = "application/json"
)

// 审批常量
const (
	CC_Position_Start = "START"
	CC_Position_Finish = "FINISH"
	CC_Position_Start_Finish = "START_FINISH"
)