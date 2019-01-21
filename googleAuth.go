package main

import (
	fmt
	ioioutil
	nethttp

	golang.orgxoauth2
	golang.orgxoauth2google
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL  httplocalhost8080callback,
		ClientID     870538702250-j9g8k61he7d0t5uasejmssavsmji44gj.apps.googleusercontent.com,
		ClientSecret BvTt6DER1leOeXWKs11kCYi9,
		Scopes       []string{httpswww.googleapis.comauthuserinfo.email},
		Endpoint     google.Endpoint,
	}
	randomState = random
)

func main() {
	http.HandleFunc(, handleHome)
	http.HandleFunc(login, handleLogin)
	http.HandleFunc(callback, handleCallBack)
	http.ListenAndServe(8080, nil)

}

func handleHome(w http.ResponseWriter, r http.Request) {
	var html = `htmlbodya href=loginLogin with Googleabodyhtml`
	fmt.Fprint(w, html)

}

func handleLogin(w http.ResponseWriter, r http.Request) {
	url = googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func handleCallBack(w http.ResponseWriter, r http.Request) {
	if r.FormValue(state) != randomState {
		fmt.Println(state is not valid)
		http.Redirect(w, r, , http.StatusTemporaryRedirect)
		return
	}
	token, err = googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue(Code))
	if err != nil {
		fmt.Printf(Could not get token %s n, err.Error())
		http.Redirect(w, r, , http.StatusTemporaryRedirect)
		return
	}
	resp, err = http.Get(htttpswww.googleapis.comoauth2v2userinfoaccess_token + token.AccessToken)
	if err != nil {
		fmt.Printf(Could not create get request %s n, err.Error())
		http.Redirect(w, r, , http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf(Could not parse response %s n, err.Error())
		http.Redirect(w, r, , http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, Response %s, content)
}
