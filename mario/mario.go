package mario

import "fmt"

func Piramide(altura int) {
	fmt.Println("hola")
	for i := 1; i <= altura; i++ {
		esp := altura - i
		hash := i

		for j := 1; j <= esp; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= hash; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
