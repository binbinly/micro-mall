package constvar

import "fmt"

const (
	// _onlinePrefix 在线key前缀
	_onlinePrefix = "user:online:%d"
	// _userPrefix 用户令牌前缀
	_userPrefix = "user:token:%d"
	//_historyPrefix 离线消息前缀
	_historyPrefix = "history:message:%d"
)

// BuildHistoryKey 历史消息键
func BuildHistoryKey(userID int64) string {
	return fmt.Sprintf(_historyPrefix, userID)
}

// BuildOnlineKey 用户在线键
func BuildOnlineKey(userID int64) string {
	return fmt.Sprintf(_onlinePrefix, userID)
}

// BuildUserTokenKey 用户令牌键
func BuildUserTokenKey(userID int64) string {
	return fmt.Sprintf(_userPrefix, userID)
}
