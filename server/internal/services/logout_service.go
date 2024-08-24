package services

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func LogoutService(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		// Clear session
		session.Values["user_id"] = 0
		session.Options.MaxAge = -1
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
