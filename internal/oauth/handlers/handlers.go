package OAuthHandlers

import (
	Env "codecompanion/timesheet/api/internal/env"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type authorizeParams struct {
	redirectUri string
	provider    string
}

func fillAuthorizeParams(r *http.Request) authorizeParams {
	return authorizeParams{
		redirectUri: r.URL.Query().Get("redirect_uri"),
		provider:    r.URL.Query().Get("provider"),
	}
}

func validateAuthorizeParams(params authorizeParams) error {
	if params.redirectUri == "" {
		return fmt.Errorf("redirect_uri is required")
	}

	if params.provider == "" {
		return fmt.Errorf("provider is required")
	}

	return nil
}

type state struct {
	RedirectUri string `json:"redirect_uri"`
}

func authorize(w http.ResponseWriter, r *http.Request) {
	params := fillAuthorizeParams(r)

	validationErr := validateAuthorizeParams(params)

	if validationErr != nil {
		http.Error(w, validationErr.Error(), http.StatusBadRequest)
		return
	}

	state, err := json.Marshal(state{RedirectUri: params.redirectUri})
	if err != nil {
		http.Error(w, "Error marshalling state", http.StatusInternalServerError)
		return
	}

	switch params.provider {
	case "google":
		callbackUrl := "http://localhost:3000/oauth/callback/google"
		urlString := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=profile email&state=%s", Env.Vars.GoogleClientId, callbackUrl, base64.StdEncoding.EncodeToString(state))
		_, err := url.Parse(urlString)

		if err != nil {
			http.Error(w, "Error parsing google url", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, urlString, http.StatusFound)
		return
	}

	http.Error(w, "Invalid provider given", http.StatusNotFound)
}

func getState(r *http.Request) (state, error) {
	stateString := r.URL.Query().Get("state")
	stateBytes, err := base64.StdEncoding.DecodeString(stateString)

	if err != nil {
		return state{}, err
	}

	var s state
	err = json.Unmarshal(stateBytes, &s)

	if err != nil {
		return state{}, err
	}

	return s, nil
}

func googleCallback(w http.ResponseWriter, r *http.Request) {
	state, err := getState(r)

	// TODO: Get code from query params, exchange for token, get user info and store and redirect to redirect_uri with generated token

	if err != nil {
		http.Error(w, "Error getting state", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, state.RedirectUri, http.StatusFound)
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/oauth/authorize", authorize)
	mux.HandleFunc("/oauth/callback/google", googleCallback)
}
