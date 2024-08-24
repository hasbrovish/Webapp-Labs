package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func AuthMiddleware(store *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			session, _ := store.Get(r, "session-name")
			fmt.Println("get", session.Values)
			userID, ok := session.Values["user_id"].(uint)
			if !ok || userID == 0 {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
