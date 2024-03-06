package servicos

import (
	"api_golang_ia/models" // Substitua pelo caminho correto do seu pacote models
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type IAServico struct{}

func (ia *IAServico) BuscaPalavras() []models.Palavra {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("A variável de ambiente OPENAI_API_KEY não está definida.")
		return nil
	}

	messages := []map[string]string{
		{
			"role":    "system",
			"content": "Sua missão é retornar para mim uma lista de palavras em ingles com 6 alternativas e uma correta, o formato retornado será assim: [{\"palavra\": \"Hello\", \"traducao\": \"Olá\", opcoes: [\"Boa\", \"Ok\", \"Olá\", \"Bacana\"] }] somente o json e nenhum outro texto",
		},
		{
			"role":    "user",
			"content": "Traga para mim a o campo 'palavra' em ingles, o campo 'traducao' em pt-br e o campo 'opcoes' uma lista de 6 itens em 'pt-br ",
		},
	}

	url := "https://api.openai.com/v1/chat/completions"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": messages,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, erro := io.ReadAll(resp.Body)

	if erro != nil {
		fmt.Println(err.Error())
		return nil
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(responseBody), &result)

	var palavras []models.Palavra

	fmt.Println("=============")
	fmt.Println(result)
	fmt.Println("=============")

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if firstChoice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := firstChoice["message"].(map[string]interface{}); ok {
				resposta := message["content"].(string)

				fmt.Println("=============")
				fmt.Println(resposta)
				fmt.Println("=============")

				err := json.Unmarshal([]byte(resposta), &palavras)
				if err != nil {
					fmt.Println("Erro ao deserializar resposta:", err)
					return nil
				}
				return palavras
			}
		}
	}

	fmt.Println("Não foi possível obter uma resposta válida.")
	return nil
}
