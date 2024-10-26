package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	// opcao de abortar a operacao caso chegue ate 3 seg
	ctx, cancel := context.WithTimeout(ctx, time.Second*3) // apos 3 seg é cancelado, ele executa o DONE
	defer cancel()
	bookHotel(ctx)
}

// receber contexto
func bookHotel(ctx context.Context) {

	select {
	case <-ctx.Done(): // se contexto foi finalizado, excedeu o tempo dos 3 seg
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	// abaixo pode ser qq outra operacao (chamar api, por ex)
	case <-time.After(1 * time.Second): //se apos 1 seg o contexto NAO foi cancelado
		fmt.Println("Hotel booked.")
	}

}

// async
