package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// a request tb recebe context
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // pega o context da requisicao
	log.Println("request iniciada")
	defer log.Println("request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// imprime na tela do servidor, no stdout
		log.Println("request processada com sucesso")
		// imprime no browser
		w.Write([]byte("request processada com sucesso"))
	case <-ctx.Done(): // ctrl + c ou esc do lado do client
		// imprime na tela do servidor, no stdout
		log.Println("request cancelada pelo cliente")
	}
}
