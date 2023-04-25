package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func homePortfolio(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmplPortfolio/accueil.html")
}

func myWork(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmplPortfolio/myWork.html")
}

func discoverMe(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmplPortfolio/discoverMe.html")
}

func contact(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmplPortfolio/contact.html")
}

func main() {
	http.HandleFunc("/home", homePortfolio)
	http.HandleFunc("/myWork", myWork)
	http.HandleFunc("/discoverMe", discoverMe)
	http.HandleFunc("/contactUs", contact)

	http.Handle("/portfolio/", http.StripPrefix("/portfolio", http.FileServer(http.Dir("./portfolio"))))
	http.Handle("/ProjectsImg/", http.StripPrefix("/ProjectsImg/", http.FileServer(http.Dir("./ProjectsImg/"))))
	http.Handle("/iconsImg/", http.StripPrefix("/iconsImg/", http.FileServer(http.Dir("./iconsImg/"))))
	http.Handle("/Pics/", http.StripPrefix("/Pics/", http.FileServer(http.Dir("./Pics/"))))
	http.Handle("/backgroundImg/", http.StripPrefix("/backgroundImg/", http.FileServer(http.Dir("./backgroundImg/"))))

	fmt.Println("\n(http://localhost:8080/home) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
