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

const monitamento = 5
const delay = 10

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

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
			fmt.Println("Não reconheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "laiza"
	versao := 1.1
	fmt.Println("Olá ", nome)
	fmt.Println("A versão utilizada é ", versao)
	fmt.Println(" ")
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)
	fmt.Println(" ")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println(" ")
	//sites := []string{"https://www.alura.com.br/", "https://www.alura.com.br/", "https://www.alura.com.br/"}

	sites := leSiteDeArquivo()

	for i := 0; i < monitamento; i++ {
		for i, site := range sites {
			fmt.Println("Estou passando por esse indice ", i, "E estou nesse site: ", site)
			testaSite(site)
			fmt.Println(" ")
			fmt.Println("----------------------------------------------------------------------------------------- ")
		}
		time.Sleep(delay * time.Second)

	}

	fmt.Println("  ")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "O monitaramento está tudo ok!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "Houve algum erro no monitoramento", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSiteDeArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
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

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("logs.txt")

	fmt.Println(string(arquivo))

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
}
