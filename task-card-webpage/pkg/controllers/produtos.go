package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/luuisavelino/github.com/luuisavelino/task-card-webpage/pkg/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var apiKey = os.Getenv("BACKENT_API_KEY")

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.New(apiKey)
	allProducts, err := produtos.SearchAll()
	if err != nil {
		fmt.Println(err.Error())
		temp.ExecuteTemplate(w, "Index", nil)
		return
	}

	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	produtos := models.New(apiKey)

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço!", err.Error())
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade!", err.Error())
		}

		produtos.Create(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	http.Redirect(w, r, "/api/v1/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	produtos := models.New(apiKey)

	idDoProduto := r.URL.Query().Get("id")

	produtos.Delete(idDoProduto)

	http.Redirect(w, r, "/api/v1/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	produtos := models.New(apiKey)

	idDoProduto := r.URL.Query().Get("id")

	produto, err := produtos.Read(idDoProduto)
	if err != nil {
		fmt.Println(err)
	}

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	produtos := models.New(apiKey)

	if r.Method == "PATCH" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco para float:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade para int:", err)
		}

		produtos.Update(id, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	http.Redirect(w, r, "/api/v1/", http.StatusPermanentRedirect)
}
