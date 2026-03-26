package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const BrasilAPIBaseUrl = "https://brasilapi.com.br/api/cep/v1"
const ViaCEPAPIBaseUrl = "https://viacep.com.br/ws"

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCEPAPIResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	arg := os.Args[1]

	chBrasil := make(chan BrasilAPIResponse)
	chViaCep := make(chan ViaCEPAPIResponse)

	go GetCepFromBrasilAPI(arg, chBrasil)
	go GetCepFromViaCEPAPI(arg, chViaCep)

	select {
	case result := <-chBrasil:
		fmt.Println("Resposta mais rapida: BrasilAPI")
		fmt.Printf("%+v\n", result)
	case result := <-chViaCep:
		fmt.Println("Resposta mais rapida: ViaCEP")
		fmt.Printf("%+v\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: nenhuma API respondeu em 1 segundo")
	}
}

func GetCepFromBrasilAPI(arg string, ch chan<- BrasilAPIResponse) {
	url := fmt.Sprintf("%v/%v", BrasilAPIBaseUrl, arg)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro na request BrasilAPI: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var brasilApiResponse BrasilAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&brasilApiResponse)
	if err != nil {
		fmt.Printf("Erro no decode BrasilAPI: %v\n", err)
		return
	}

	ch <- brasilApiResponse
}

func GetCepFromViaCEPAPI(arg string, ch chan<- ViaCEPAPIResponse) {
	url := fmt.Sprintf("%v/%v/json", ViaCEPAPIBaseUrl, arg)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro na request ViaCEP: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var viaCepApiResponse ViaCEPAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&viaCepApiResponse)
	if err != nil {
		fmt.Printf("Erro no decode ViaCEP: %v\n", err)
		return
	}

	ch <- viaCepApiResponse
}
