package oauth2

import (
	"html/template"
	"log"
	"net/http"
)

func Serve() {
	log.Fatal(start())
}

func start() error {
	fs := http.FileServer(http.Dir(config.FrontendStaticFiles))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", handler)
	log.Println("Listening on port", config.Port, "...")
	return http.ListenAndServe(":"+config.Port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := OAuthData{
		RefreshToken: getCookieValue(r, "RefreshToken"),
		BearerToken:  getCookieValue(r, "BearerToken"),
		OauthExpires: getCookieValue(r, "OauthExpires"),
		OauthHMAC:    getCookieValue(r, "OauthHMAC"),
	}

	tmpl, err := template.ParseFiles(config.FrontendStaticFiles + "/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
