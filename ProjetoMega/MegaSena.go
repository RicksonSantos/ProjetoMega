package main

import (
	f "fmt"
	r "math/rand"
	t "time"
)

func NumeroAleatorio() int {
	s := r.NewSource(t.Now().UnixNano())
	a := r.New(s)
	return a.Intn(60)
}

func main() {

	hora := t.Now()
	switch {
	case hora.Hour() < 12:
		f.Println("Bom dia, Jogador(a).")
	case hora.Hour() < 18:
		f.Println("Boa tarde, Jogador(a).")
	default:
		f.Println("Boa noite, Jogador(a).")
	}

	f.Println("Sugestão de números para jogar na Mega-Sena (entre os disponiveis no volante):")

	for c := 0; c < 15; c++ {
		i := NumeroAleatorio()
		f.Print(" ", i)
	}
}
