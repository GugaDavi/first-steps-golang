# Anotações de estudos em Go

## Tipos de variáves:

Os tipos de variáveis são:

**string** - Texto;

**bool** - Boolean

**int** - Número inteiro:

Pode ter diversos formatos:

- int8;
- int16;
- int32;
- int64;

Por padrão qualquer valor inteiro é: `int64`

**uint** - Número inteiro sem sinal, maiores que 0:

Pode ter diversos formatos:

- uint8;
- uint16;
- uint32;
- uint64;

Por padrão qualquer valor inteiro é: `uint64`

**float** - Número com ponto flutuante:

Pode ter os seguintes formatos:

- float32;
- float64;

Padrão: `float64`

> Existe outros tipos númerios, como: `byte` e `rune`. Verificar em: https://www.tutorialspoint.com/go/go_data_types.htm

## Declaração de variável:

1. `var name string = "Gustavo"`;
2. `var name = "Gustavo"`;
3. `var name, lastName = "Gustavo", "Guedes"`;
4. `var name, lastName string`;
5. `name := "Gustavo"`;
6.

```go
var (
  name string
  idade int
  age = 27
)
```

## Parseando dados:

Basta utilizar o constructor do tipo para isso:

Exempo:

```go
var number = 27

float64(number)
```

**De tipo número para string**: `fmt.Sprint(number)`

## Importação de pacotes:

`import "fmt"`

```go
import (
  "fmt"
  "math"
)
```

## Funções:

Funções com um tipo de retorno:

```go
func soma(a int, b int) int {
  return a + b
}

// OR

func soma(a, b int) int {
  return a + b
}
```

Com dois ou mais tipos re retorno:

```go
func calcular(a int) (int, int) {
  quadrado := a * a
  cubo := a * a * a

  return quadrado, cubo
}
```

## Condicionais:

```go
if a > 10 {
  fmt.Println("a maior que 10")
} else if a > 5 {
  fmt.Println("a está entre 6 e 10")
} else {
  fmt.Println("a é menor ou igual a 5")
}
```

Sem utilização de parenteses nas condicionais.

É possivel settar variaveis para o escopo das condicionais:

```go
if a := 11; a > 10 {
  fmt.Println("a maior que 10")
} else if a > 5 {
  fmt.Println("a está entre 6 e 10")
} else {
  fmt.Println("a é menor ou igual a 5")
}
```

A variavel `a` só poderá ser acessada no escopo das condicionais.

## Loops:

For:

```go
for i := 1; i <= 10; i++ {
  fmt.Println(i)
}
```

While:

Utiliza a mesma syntax do for, mas só passando a condicional:

```go
var soma int = 2

for soma <= 100 {
  soma += soma
}
```

## Switch

Por padrão o Go já faz a implementação do break no final dos cases

```go
switch number {
  case 1:
    fmt.Println("Número: 1")
  case 2:
    fmt.Println("Número: 2")
  default:
    fmt.Println("É qualquer número")
}
```

Devido break automatico do go, não é possivel case encadeados:

```go
switch name {
  case "Gustavo":
  case "Matheus":
    fmt.Println("Gustavo e Matheus")
  default:
    fmt.Println("É qualquer número")
}
```

No caso acima, se o valor for `"Gustavo"`, não será retornado nada.

Switch `true`:

```go
nota := 7

switch {
  case nota > 9:
    fmt.Println("Ótima nota")
  case nota > 7:
    fmt.Println("Boa Nota")
  case nota > 5:
    fmt.Println("Aprovado")
  default:
    fmt.Println("Reprovado")
}
```

Usado como uma alternativa para o uso de `if`, `else if` e `else` encadeados.

## Arrays

```go
var numeros [7]int
```

Iniciamos um array vazio, mas que possuirá 7 "elementos" dentro da coleção.

```go
var numeros = [7]int{1, 2, 3, 4, 5, 6, 7}
```

Nesse caso inicializamos o array com todas as 7 posições preenchidas.

É possivel iniciar um array limitando o número de elementos `[7]` ou não `[]`.
Caso seja pre-settado, o array não poderá ser maior do que o settado inicialmente.
Quando é passado um número maximo de elementos, mesmo que seja passado menos do indicado, tipo:

`[7]int{1, 2, 3, 4}`

As outras posições serão ocupadas pelo valor padrão do tipo da variavel, nesse caso o resultado será:

`[1, 2, 3, 4, 0, 0, 0]`

Interando um array:

```go
var numeros = []int{1, 2, 3, 4, 5, 6, 7}

for i := 0; i < len(numeros); i++ {
  fmt.Println(numeros[i])
}
```

O método `len`(length) retorna o tamanho do parametro passado.

### Slice

Retorna um array baseado no original, baseado nos index passados.

```go
var numeros = [7]int{1, 2, 3, 4, 5, 6, 7}

// Slice
var novosNumeros = numeros[2:5]

fmt.Println(novosNumeros) // [3, 4, 5]
```

O primeiro número é o index inicial e o segundo é posterior ao ultimo desejado.

Os parametros são opcionais:

```go
numeros[2:5] // [3, 4, 5]
numeros[2:] // [3, 4, 5, 6, 7]
numeros[:5] // [1, 2, 3, 4, 5]
numeros[:] // [1, 2, 3, 4, 5, 6, 7]
```

Quando é feito essa segmentação, a referencia do array original continua a mesma.
Se algum dado for alterado no slice, será alterando também no array original.

### Slice 0

Quando um slice é inicializado sem nenhum parametro de cap ou length, ele é settado como:

- `[]` conjunto vazio;
- `0` zero;
- `nil` Valor nil;

### SliceOfSlice:

É usado para tipagem de array dentro de arrays:

```go
var numeros = [][]string{
  []string{"Gustavo", "Davi", "Guedes"},
  []string{"Herta", "Milene", "França"},
  []string{"Matheus", "Vieira", "Nascimento"},
}
```

### Pontieros:

Usado para localizar valores armazenados na memoria:

```go
var a int = 32

var p *int = &a
*p = 18

println(a)
```

Nesse exemplo, o `p` irá armazenar o local da memoria de `a`, tipo: `0xc00004c770`.
E utilizando o `*` antes de `p` conseguimos o valor de `a`: `32`.
Podendo até sobscrever o valor original.

## Métodos nativos:

- `cap` - Retorna a capacidade maxima de um slice;
- `len` - Retorna o número de elementos de um elemento;
- `make` - Constroe um array recebendo como parametro: `([]type, lenghtInicial, Capacidade)`
- `append` - Adiciona um novo elemento a um array: `append(array, element1, element2, ...)`
- `strings.Join` - Concatenar strings: `strings.Join(variable, ", ")` (Necessário importação da lib: `strings`)
- `defer` - Joga a execução de uma linha para o final da execução. É feito um esquema de pilha, onde o ultimo a entrar é o primeiro a sair
  - Caso a linha com defer possua a execução de uma função, o que será amazenado para ser executado no final do codigo será o retorno dela
