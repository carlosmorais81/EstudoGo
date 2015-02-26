package main

import "fmt"

type ListaDeCompras []string

func imprimirSlice(slice [] string){
    fmt.Println("Slice: ",slice)
}

func main() {

    lista := ListaDeCompras{"Alface","Atum","Azeite"}

    slice := []string{"Alface","Atum","Azeite"}

    imprimirSlice([]string(lista))
    imprimirSlice(ListaDeCompras(slice))

}
