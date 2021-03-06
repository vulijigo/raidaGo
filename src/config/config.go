package config


const DEFAULT_TIMEOUT = 5
const DEFAULT_DOMAIN = "cloudcoin.global"


const LOG_LEVEL_DEBUG = 3
const LOG_LEVEL_INFO = 2
const LOG_LEVEL_ERROR = 1


const REMOTE_RESULT_ERROR_NONE = 0
const REMOTE_RESULT_ERROR_TIMEOUT = 1
const REMOTE_RESULT_ERROR_COMMON = 2
const REMOTE_RESULT_ERROR_SKIPPED = 3

const DEFAULT_NN = 1
const PUBLIC_CHANGE_MAKER_ID = 2

var CmdDebug bool
var CmdHelp bool
var CmdVersion bool
var CmdCommand string

const TOTAL_RAIDA_NUMBER = 25

const (
	RAIDA_STATUS_UNTRIED = 0
	RAIDA_STATUS_PASS = 1
	RAIDA_STATUS_ERROR = 2
	RAIDA_STATUS_FAIL = 3
	RAIDA_STATUS_NORESPONSE = 4
)


const TOPDIR = "CloudCoinStorage"

const DIR_ID = "ID"


const CHANGE_METHOD_250F = 1
const CHANGE_METHOD_100E = 2
const CHANGE_METHOD_25B = 3
const CHANGE_METHOD_5A = 4


const MAX_NOTES_TO_SEND = 3
