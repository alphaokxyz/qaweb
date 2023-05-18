package errormessage

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIST     = 1003
	ERROR_TOKEN_NOTEXIST     = 1004
	ERROR_TOKEN_RUNTIME      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_TYPE_WRONG   = 1007
	ERROR_USER_NO_RIGHT      = 1008
	ERROR_QUESTION_NOT_EXIST = 2001
	ERROR_ANSWER_NOT_EXIST   = 2002
)

var codemessage = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已存在！",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_TOKEN_NOTEXIST:     "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:      "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:        "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:      "该用户无权限",
	ERROR_QUESTION_NOT_EXIST: "问题不存在",
	ERROR_ANSWER_NOT_EXIST:   "回答不存在",
}

func Geterrormessage(code int) string {
	return codemessage[code]
}
