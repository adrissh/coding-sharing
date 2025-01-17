package main

import (
	"fmt"
	"net/http"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name:     "access_token",
		Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.S5OGQJ7Lfj2tkoEHQRGv3K22pK5_fWzIVrGpTlKrJDE",
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	// send cookie via header
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie has been set!")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		w.Write([]byte("cookies not found"))
		return
	}
	fmt.Fprintf(w, "Cookie is Found: %s = %s", cookie.Name, cookie.Value)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {

	// Check if the cookie exists or not
	_, err := r.Cookie("access_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Write([]byte("cookie not found"))
			return
		}
		// another error
		http.Error(w, "failed to get cookie", http.StatusInternalServerError)
		return
	}
	// delete cookie
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	w.Write([]byte("Cookie successfully deleted!"))
}
func main() {
	http.HandleFunc("/set-cookie", SetCookie)
	http.HandleFunc("/get-cookie", GetCookie)
	http.HandleFunc("/del-cookie", DeleteCookie)
	fmt.Println("server runnig on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
