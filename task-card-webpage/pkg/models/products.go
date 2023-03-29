package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const apiVersion = "/api/v1"

var BaseURL = os.Getenv("BASE_URL")

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

type Actions interface {
	SearchAll()
	Read()
	Delete()
	Create()
	Update()
}

func ApiConsume(method, apiKey, endpoint string, reqBody io.Reader) (io.Reader, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		method, fmt.Sprintf(BaseURL+apiVersion+endpoint), reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res.Body, nil
}

func (p *Products) SearchAll() ([]Product, error) {
	reqBody := bytes.NewBufferString(``)

	Body, err := ApiConsume(http.MethodGet, p.apiKey, "/products", reqBody)
	if err != nil {
		return nil, err
	}

	var produtos []Product
	if err := json.NewDecoder(Body).Decode(&produtos); err != nil {
		return nil, err
	}

	return produtos, nil
}

func (p *Products) Read(id string) (Product, error) {
	Body, err := ApiConsume(http.MethodGet, p.apiKey, "/products/"+id, nil)
	if err != nil {
		return Product{}, err
	}

	var produto Product
	if err := json.NewDecoder(Body).Decode(&produto); err != nil {
		return Product{}, err
	}

	return produto, nil
}

func (p *Products) Delete(id string) {
	body, err := ApiConsume(http.MethodDelete, p.apiKey, "/products/"+id, nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}

func (p *Products) Create(nome, descricao string, preco float64, quantidade int) {

	req, err := json.Marshal(Product{
		Name:        nome,
		Description: descricao,
		Price:       preco,
		Amount:      quantidade,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	reqBody := bytes.NewBufferString(string(req))

	body, err := ApiConsume(http.MethodPost, p.apiKey, "/products", reqBody)

	if err != nil {
		fmt.Println(err)
	}
	log.Println(body)
}

func (p *Products) Update(id, nome, descricao string, preco float64, quantidade int) {
	req, err := json.Marshal(Product{
		Name:        nome,
		Description: descricao,
		Price:       preco,
		Amount:      quantidade,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	reqBody := bytes.NewBufferString(string(req))

	body, err := ApiConsume(http.MethodPatch, p.apiKey, "/products"+id, reqBody)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}
