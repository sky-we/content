package utils

import "fmt"

func GenSessionKey(sessionId string) string {
	return fmt.Sprintf("Session-ID:%s", sessionId)
}

func GenLoginKey(userId string) string {
	return fmt.Sprintf("user:login:%s", userId)
}
