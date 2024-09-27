package ycuser

import "sync"

type ycUser struct {
	mu     sync.RWMutex
	cookie string
}

func (u *ycUser) getCookie() string {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.cookie
}

func (u *ycUser) setCookie(cookie string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.cookie = cookie
}
