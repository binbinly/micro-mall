package constvar

import "fmt"

const (
	// _userPrefix 用户令牌前缀
	_userPrefix = "user:token:%d"
)

// BuildUserTokenKey 用户令牌键
func BuildUserTokenKey(userID int64) string {
	return fmt.Sprintf(_userPrefix, userID)
}
