package main

func main() {
	// number := soma(5, 2)

	// parsedNumber := fmt.Sprint(number)
	// value1, value2 := calcular(number)

	// fmt.Printf("%T %v\n", parsedNumber, parsedNumber)
	// fmt.Printf("Valor 1: %v / Valor 2: %v\n", value1, value2)

	// if n := number + 1; n < 10 {
	// 	fmt.Printf("Number + 1 = %v\n", n)
	// }

	// var soma int = 2

	// for soma <= 100 {
	// 	soma += soma
	// }

	// fmt.Println(soma)

	// switch 5 {
	// case 1:
	// case 2:
	// case 3:
	// case 4:
	// case 5:
	// 	fmt.Println("Está entre 1 e 5")
	// case 7:
	// 	fmt.Println("Número: 7")
	// default:
	// 	fmt.Println("É qualquer número")
	// }

	// nota := 9

	// switch {
	// case nota > 9:
	// 	fmt.Println("Ótima nota")
	// case nota > 7:
	// 	fmt.Println("Boa Nota")
	// case nota > 5:
	// 	fmt.Println("Aprovado")
	// default:
	// 	fmt.Println("Reprovado")
	// }

	// var numeros = [7]int{1, 2, 3, 4, 5, 6, 7}

	// somaN := 0

	// for i := 0; i < len(numeros); i++ {
	// 	somaN += numeros[i]
	// }

	// fmt.Println(somaN)

	// // Slice
	// var novosNumeros = numeros[2:5]

	// fmt.Println(novosNumeros) // [3, 4, 5]

	// var animes [3]string

	// animes = [3]string{"Dragon Ball", "Sailor Moon", "Yuyu Hakusho"}

	// sliceAnimes := animes[1:]

	// fmt.Println(animes)
	// fmt.Println(sliceAnimes)
	// fmt.Println("Capacidade:", cap(sliceAnimes), len(sliceAnimes))

	// var makedArray = make([]int, 0, len(animes))

	// fmt.Println(makedArray)
	// fmt.Println("Gustavo tem:", len("Gustavo"), "letras")

	// var number []int

	// fmt.Println(number == nil)

	// var names = [][]string{
	// 	[]string{"Gustavo", "Davi", "Guedes"},
	// 	[]string{"Herta", "Milene", "França"},
	// 	[]string{"Matheus", "Vieira", "Nascimento"},
	// }

	// for i := 0; i < len(names); i++ {
	// 	for j := 0; j < len(names[i]); j++ {
	// 		println(names[j][i])
	// 	}
	// }
	// defer println("Oi!")
	// println("Opa")

	a := 32

	p := &a
	// *p = 18

	println(p)
}

func soma(a, b int) int {
	return a + b
}

func calcular(a int) (int, int) {
	quadrado := a * a
	cubo := a * a * a

	return quadrado, cubo
}
