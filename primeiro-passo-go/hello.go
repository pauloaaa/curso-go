package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 5
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		// nome, idade := nomeIdade()
		// fmt.Println(nome, "tenho", idade, "anos")

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

}

func nomeIdade() (string, int) {
	nome := "Paulo"
	idade := 35
	return nome, idade
}

func exibeIntroducao() {
	nome := "Paulo"
	versao := 1.1
	fmt.Println("Oĺá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	// fmt.Println("O tipo da variavael nome é", reflect.TypeOf(nome))
	// fmt.Println("O tipo da variavael versao é", reflect.TypeOf(versao))
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	// fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br", "https://www.google.com.br"}

	sites := leSiteArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func exibeNomes() {
	nomes := []string{"Paulo", "Bebeto", "Arruda", "Ingrides", "Heloisa"}
	nomes = append(nomes, "Jose")
	fmt.Println(nomes)
}

func leSiteArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
