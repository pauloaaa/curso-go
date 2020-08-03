package routes

import (
	"net/http"

	"github.com/paulo-arruda/controllers"
)

// CarregaRotas redirencionar
func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Cadastrar)
	http.HandleFunc("/delete", controllers.Excluir)
	http.HandleFunc("/edit", controllers.Carregar)
	http.HandleFunc("/update", controllers.Editar)

}
