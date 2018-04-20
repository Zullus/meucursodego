package matematica


//Calculo calcula com uma função como parametro
func Calculo(funcao func(int, int) int, numero1 int, numero2 int) int{

    return funcao(numero1, numero2)

}

//Multiplicador multiplica um número x pelo y
func Multiplicador(x int, y int) int {

    return x * y

}

//Divisor divide dois número
func Divisor(a int, b int) (resultado int) {

    resultado = a/b

    return
}

//DivisorComResto divide dois número, com resto
func DivisorComResto(a int, b int) (resultado int, resto int) {

    resultado = a/b

    resto = a%b

    return
}