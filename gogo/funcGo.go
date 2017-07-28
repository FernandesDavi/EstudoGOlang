package main
import (
	"fmt"
)
func soma(x,y int) int{
	return x+y
}
func showMeThisFuck(nome string, idade int){
	fmt.Printf("o que vc quer ver monstrao? %s %d \n",nome,idade)
}

func main(){
	showMeThisFuck("davi",21)
	//var res int
	res := soma(1,156)
	fmt.Printf("A soma é %d",res)
}