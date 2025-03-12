package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"name"`
	Email  string `json:"email"`
}

var usuarios []Usuario

// Usuarios

func RetornarJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("test-header", "header")
	json.NewEncoder(w).Encode(data)
}

func LeerUsuario(w http.ResponseWriter, r *http.Request) Usuario {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
	}
	var user Usuario
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Error parseando el JSON", http.StatusBadRequest)
	}
	return user
}

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RetornarJson(w, usuarios)
	case http.MethodPost:
		user := LeerUsuario(w, r)
		user.ID = len(usuarios) + 1
		usuarios = append(usuarios, user)

		RetornarJson(w, user)
	case http.MethodDelete:
		user := LeerUsuario(w, r)
		for index := range usuarios {
			if usuarios[index].ID == user.ID {
				usuarios = append(usuarios[:index], usuarios[index+1:]...)
				fmt.Fprintln(w, "Usuario eliminado")
				return
			}
		}
	case http.MethodPut:
		user := LeerUsuario(w, r)
		for index := range usuarios {
			if usuarios[index].ID == user.ID {
				usuarios[index] = user
				fmt.Fprintln(w, "Usuario actualizado")
				return
			}
		}
	default:
		http.Error(w, "Método no permitido", 405)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "pong")
	default:
		http.Error(w, "Método no permitido", 405)
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("./public/index.html")
	if err != nil {
		fmt.Fprintln(w, "error leyendo el html")
		return
	}
	fmt.Fprintln(w, string(content))
}
func main() {
	usuarios = append(usuarios, Usuario{
		ID:     1,
		Nombre: "Alfredo",
		Email:  "Alfredo@mail.com",
	})
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/v1/users", Users)
	http.HandleFunc("/", Index)

	fmt.Println("Servidor escuchando en el puerto 3000")
	fmt.Println("http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
