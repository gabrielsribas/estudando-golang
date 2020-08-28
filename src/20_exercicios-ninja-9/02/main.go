package main

import (
   "fmt"
)

type pessoa struct {
   nome string
   sobrenome string
   idade int
   dizendocoisas []string
}

func (s *pessoa) falar() {
   fmt.Printf("Meu nome completo é %v %v e tenho %v anos\n- %v\n", s.nome, s.sobrenome, s.idade, s.dizendocoisas)   
}

type humanos interface {
   falar()
}

func dizerAlgumaCoisa(s humanos) {
   s.falar()
}

func main() {

   p1 := pessoa{
      nome: "Gabriel",
      sobrenome: "Ribas",
      idade: 32,
      dizendocoisas: []string{"Uma das coisas mais fantásticas que existem é comer pão com ovo e iogurte.", "Sem contar o famoso miojo com toddy.", "Minha vida é comer essas delícias..."},
   }

   p1.falar() // Aqui esta usando o metodo falar que faz um ponteiro. Entao esta implicito. (&p1).falar()
   //dizerAlgumaCoisa(p1) // Isso aqui nao funciona. porque é um ponteiro
   dizerAlgumaCoisa(&p1) // Isso aqui funciona porque está referenciando o ponteiro. Pois o receiver é um ponteiro.
   
}
