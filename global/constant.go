package global

const (
	ERROR_OK     int = 0
	ERROR_FAILED int = 1

	ERROR_TOKEN_EMPTY   int = 11
	ERROR_TOKEN_EXPIRED int = 12
)

var ERROR_STRING = map[int]string{
	ERROR_OK:            "成功",
	ERROR_FAILED:        " 失败",
	ERROR_TOKEN_EMPTY:   "token无效",
	ERROR_TOKEN_EXPIRED: "token过期",
}
