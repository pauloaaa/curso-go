package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/paulo-arruda/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index redirecionar para a tela a lista de produtos
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.ListarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)

}

// New redirencioar para a tela novo produto
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// Cadastrar novo produto
func Cadastrar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			fmt.Println("Erro ao converter o preço:", err)
		}

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			fmt.Println("Erro ao converter o quantidade:", err)
		}

		models.CadastrarProduto(nome, descricao, preco, quantidade)
	}

	http.Redirect(w, r, "/", 301)
}

// Excluir excluir produto
func Excluir(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.ExcluirProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

// Editar editar produto
func Carregar(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")

	produto := models.CarregarProduto(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

// Editar editar produto
func Editar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id, err := strconv.Atoi(r.FormValue("id"))

		if err != nil {
			fmt.Println("Erro ao converter o id:", err)
		}

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			fmt.Println("Erro ao converter o preço:", err)
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			fmt.Println("Erro ao converter o quantidade:", err)
		}

		models.EditarProduto(id, nome, descricao, preco, quantidade)
	}

	http.Redirect(w, r, "/", 301)
}
