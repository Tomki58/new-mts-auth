package v1

import (
	"mts/auth/cookies"
	"mts/auth/httpserver/serializer"
	"net/http"
)

// login is handler for processing basic auth
func (a *App) login(w http.ResponseWriter, r *http.Request) {
	// validating authentication data
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Basic auth header not found", http.StatusUnauthorized)
		return
	}
	if pass, ok := a.Config.Credentials[username]; ok && password == pass {
		cookies, err := cookies.New(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.SetCookie(w, &cookies[0])
		http.SetCookie(w, &cookies[1])

		response, err := serializer.SerializeResponseJson("Successfully logged in!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
		return
	}
	http.Error(w, "Invalid username or password!", http.StatusUnauthorized)

}
