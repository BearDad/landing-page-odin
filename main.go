package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))

	// Esta función evita que el navegador guarde caché
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		fs.ServeHTTP(w, r)
	}))

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	fmt.Println("✅ Caché deshabilitada - los cambios se verán inmediatamente")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
