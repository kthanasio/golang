package main

func main() {

	a := 10
	println(a)

	ponteiro := &a // guarde o endereco de a
	println(ponteiro) // imprime o endereco de a
	*ponteiro = 20 // no espaco de memoria altere o valor para 20
	c := &a  // guarde o endereco de a
	println(*c) // apresente o valor no endereco de a
}