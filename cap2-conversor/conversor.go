package main

import(
       "fmt"
       "os"
       "strconv"
)

func main() {

    if len(os.Args) < 3{
        fmt.Println("Uso: conversor <valores> <unidade>")
        os.Exit(1)
    }

    unidadOrigem := os.Args[len(os.Args) -1] //entendi mais ou menos, pesquisar
    valoresOrigem  := os.Args[1 : len(os.Args) -1]

    var unidadeDestino string

    if unidadeOrigem == "celsius"{
        unidadeDestino = "fahrenheit"
    } else if unidadeOrigem == "quilometros" {
        unidadeDestino = "milhas"
    } else {
         fmt.Println("%s não é uma unidade conhecida!",unidadeDestino)
         os.Exit(1)
    }

    for i, v := range valoresOrigem {
        valorOrigem, err := strconv.ParseFloat(v,64) //nao entendi
        if err != nil {
            fmt.Println("O valor %s na posição %d " +
                        "não é um numero valido!\n",v,i)
            os.Exit(1)
        }

        var  valorDestino float64

        if unidadeOrigem == "celsius" {
            valorDestino = valorOrigem * 1.8 + 32
        } else {
            valorDestino = valorOrigem / 1.60934
        }

        fmt.Printf("%.2f %s = %.2f %s\n" ,valorOrigem , unidadeOrigem ,valorDestino , unidadeDestino)

    }

}
