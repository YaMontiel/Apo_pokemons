package main

import (
	"context"
	"database/sql"
	"gos/handler"
	"gos/repository"
	"gos/service"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configurar conexión a la base de datos
	// usuario:contraseña@tcp(host:puerto)/nombre_base
	dsn := "root:1304@tcp(127.0.0.1:3306)/mi_primera_base"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la conexión: %v", err)
	}
	defer db.Close()

	// Verificar la conexión
	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("Error al conectarse a la base de datos: %v", err)
	}

	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	h := handler.NewHandler(srv)

	log.Printf("Servidor escuchando en :8080")
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}

type Persona struct {
	Nombre string
	Edad   int
}

func buscarP(personas []Persona, nombre string) Persona {
	var retorno Persona

	for _, v := range personas {
		if v.Nombre == nombre {
			retorno = v
		}
	}
	return retorno
}

func mayorP(personas []Persona) string {
	var mayor Persona = personas[0]
	for _, v := range personas {
		if v.Edad > mayor.Edad {
			mayor = v
		}
	}
	return mayor.Nombre
}

func filtrarPorEdad(personas []Persona, edadMin int) []Persona {

	//var mayorE Persona = personas[0] como si quisiera comparar contra la mayor, no me pide esto
	resultado := []Persona{} //necesito asignar un valor inicial, aunque sea vacio por eso las{}

	for _, p := range personas {
		if p.Edad >= edadMin {
			resultado = append(resultado, p)

		}

	}
	return resultado
}

func filtrarPorNombre(personas []Persona, nombre string) (Persona, bool) {

	for _, p := range personas {

		if p.Nombre == nombre {
			return p, true
		}
	}
	return Persona{}, false
}
