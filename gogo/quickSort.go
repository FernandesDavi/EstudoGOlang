package main
import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada{
		numero, err := strconv.Atoi(n)
		if err != nil{
			fmt.Printf("%s não é um numero valido \n",n)
			os.Exit(1)
		}
		numeros[i] = numero
	}
	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int{
	if len(numeros) <= 1{
		return numeros
	}
	//criar uma copia do vertor original
	n := make([]int, len(numeros))
	copy(n, numeros)

	//Criar o pivo da lista
	indPivo := len(n) / 2
	pivo := n[indPivo]
	n = append(n[:indPivo], n[indPivo+1:]...)
	menores, maiores := particionar(n, pivo)
	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
	
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {

	for _, n := range numeros{
		if n <= pivo{
			menores = append(menores, n)
		}else{
			maiores = append(maiores, n)
		}
	}
	return menores, maiores
}