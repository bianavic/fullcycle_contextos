package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// preparar a request que chama o servidor
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	// fazer a request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	// fechar conexao
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
