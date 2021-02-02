package openHERTA

// NoSurf implements CSRF protection
/*func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   inProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}*/
