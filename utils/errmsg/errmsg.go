package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USERNAME_NOFOUND = 1001
	ERROR_PASSWORD_WERONG  = 1002
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_NOFOUND: "No Found User",
	ERROR_PASSWORD_WERONG:  "Password Wrong",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
