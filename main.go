package main
import (
	"strconv"
	"fmt"
)

func main() {
	num, err := strconv.Atoi("42")

	if err != nil {
		fmt.Printf("l'erreur est : %s", err)
	}else{
		fmt.Printf("la valeur convertie est: %d", num)
	}
}

	
