package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler devuelve un mensaje de bienvenida
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "¡Hola desde el sistema de gestión de e-commerce!")
}
