package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	} 
	tamanho, err := f.Write([]byte("Escrevendo bytes no arquivo"))
	//tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso - Tamaho %d\n", tamanho)
	f.Close()

	/// leitura total
	arquivo,err := os.ReadFile("nomes.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	/// leitura linha a linha
	arquivo2,err := os.Open("nomes.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n,err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Printf(string(buffer[:n]))
	}

	// remove
	erro := os.Remove("arquivo.txt")
	if erro != nil {
		panic(erro)
	}

}