package model

import "sync"

type YCUser struct {
	mu     sync.RWMutex
	cookie string
}

func (u *YCUser) GetCookie() string {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.cookie
}

func (u *YCUser) SetCookie(cookie string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.cookie = cookie
}
