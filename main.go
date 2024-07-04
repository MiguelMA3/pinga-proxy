package main

import (
	"fmt"
	"net/url"
	"time"

	"pinga-proxy/pkg/httpClient"
	"pinga-proxy/pkg/testaUrls"
)

func main() {
	// Defina o endereço do proxy
	proxyURL, err := url.Parse("http://proxy.exemplo:8080")
	if err != nil {
		fmt.Println("Erro ao analisar o URL do proxy:", err)
		return
	}

	// Defina o nome de usuário e a senha do proxy
	proxyUsername := "usuario"
	proxyPassword := "senha"

	// Crie um cliente HTTP com o proxy configurado e a autenticação
	client := httpClient.CreateHTTPClient(proxyURL, proxyUsername, proxyPassword)

	// Defina o intervalo de tempo entre os testes (a cada 5 minutos)
	testInterval := 5 * time.Second

	// Crie um ticker para gerar sinais a cada intervalo de teste
	ticker := time.NewTicker(testInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Teste a URL interna
			err := testaUrls.TestURL(client, "https://exemplo.com.br/")
			if err != nil {
				fmt.Printf("Erro ao testar a URL interna: %v\n", err)
				// Aqui você pode adicionar código para log ou notificação
			} else {
				fmt.Println("URL interna testada com sucesso")
			}

			// Teste a URL externa
			err = testaUrls.TestURL(client, "https://www.google.com/")
			if err != nil {
				fmt.Printf("Erro ao testar a URL externa: %v\n", err)
				// Aqui você pode adicionar código para log ou notificação
			} else {
				fmt.Println("URL externa testada com sucesso")
			}
		}
	}
}
