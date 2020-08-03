package models

import (
	"fmt"

	"github.com/paulo-arruda/db"
)

// Produto struct com todos os parametros de produto
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// ListarProdutos lista todos os produtos da base
func ListarProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	listaProdutos, err := db.Query("select * from produto order by id")

	if err != nil {
		panic(err.Error())
	}

	prod := Produto{}
	produtos := []Produto{}

	for listaProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = listaProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		prod.Id = id
		prod.Nome = nome
		prod.Descricao = descricao
		prod.Preco = preco
		prod.Quantidade = quantidade

		produtos = append(produtos, prod)
	}

	defer db.Close()
	return produtos
}

// CadastrarProduto cadastrar produto no banco
func CadastrarProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	novoProduto, err := db.Prepare("insert into produto(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		fmt.Println("Erro ao salvar produto:", err)
	}

	novoProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

//ExcluirProduto excluir produto do banco
func ExcluirProduto(id string) {
	db := db.ConectaComBancoDeDados()

	excluir, err := db.Prepare("delete from produto where id = $1")

	if err != nil {
		fmt.Println("Erro ao salvar produto:", err)
	}

	excluir.Exec(id)

	defer db.Close()
}

//CarregarProduto carregar produto pelo id
func CarregarProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	prod, err := db.Query("select * from produto where id = $1", id)

	if err != nil {
		fmt.Println("Erro ao carregar produto:", err)
	}

	editarProduto := Produto{}

	for prod.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = prod.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			fmt.Println("Erro ao carregar produto:", err)
		}
		editarProduto.Id = id
		editarProduto.Nome = nome
		editarProduto.Descricao = descricao
		editarProduto.Preco = preco
		editarProduto.Quantidade = quantidade
	}

	defer db.Close()

	return editarProduto
}

//EditarProduto editar produto no banco
func EditarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	novoProduto, err := db.Prepare("update produto set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")

	if err != nil {
		fmt.Println("Erro ao salvar produto:", err)
	}

	novoProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
