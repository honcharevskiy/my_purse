package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const baseAddress string = "https://api.monobank.ua"

type monobankProvider struct {
	apiToken string

	client http.Client
}

type clientInfo struct {
	Name     string        `json:"name"`
	Accounts []accountInfo `json:"accounts"`
}

type accountInfo struct {
	Type         string `json:"type"`
	Balance      int    `json:"balance"`
	Currency     string `json:"currency"`
	CurrencyCode int    `json:"currencyCode,omitempty"`
}

func (provider monobankProvider) clientInfo() (clientInfo, error) {
	log.Println("hitting monobank")
	const endpointPath string = "/personal/client-info"

	req, err := http.NewRequest("GET", baseAddress+endpointPath, nil)
	if err != nil {
		return clientInfo{}, err
	}

	req.Header.Set("X-Token", provider.apiToken)
	resp, err := provider.client.Do(req)

	if err != nil {
		log.Fatalf("Error when making request to monobank API, %v", err)
		return clientInfo{}, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyString := readBody(*resp)
		log.Fatal(bodyString)
		return clientInfo{}, nil
	}

	defer resp.Body.Close()

	var personalInforamtion clientInfo

	if err := json.NewDecoder(resp.Body).Decode(&personalInforamtion); err != nil {
		log.Fatalf("Error when parsing response, %v", err)
		return clientInfo{}, err
	}

	for index, accaunt := range personalInforamtion.Accounts {
		accountInArray := &personalInforamtion.Accounts[index]
		accountInArray.Currency = currencyDecoder(accaunt.CurrencyCode)
		accountInArray.Balance = balanceDecoder(accaunt.Balance)
	}

	return personalInforamtion, nil

}

func readBody(resp http.Response) string {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyBytes)
}

func currencyDecoder(currencyCode int) string {
	switch currencyCode {
	case 980:
		return "UAH"
	case 840:
		return "USD"
	case 978:
		return "EUR"
	default:
		return ""
	}
}

func balanceDecoder(balance int) int {
	if balance < 100 {
		return 0
	}
	return balance / 100
}
