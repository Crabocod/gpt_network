package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func InitSessionStore() {
	// Можно настроить параметры сессий, например, срок действия
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "session")
}
