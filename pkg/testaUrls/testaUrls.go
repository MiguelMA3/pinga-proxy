package testaUrls

import (
	"fmt"
	"net/http"
)

func TestURL(client *http.Client, url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Erro ao fazer a solicitação HTTP para a URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("URL %s acessada com sucesso (status %d)\n", url, resp.StatusCode)
		return nil
	} else {
		fmt.Printf("Erro ao acessar a URL %s (status %d)\n", url, resp.StatusCode)
		return fmt.Errorf("Erro ao acessar a URL %s", url)
	}
}
