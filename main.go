package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error cargando el archivo .env")
		os.Exit(1)
	}

	// Obtener las variables de entorno USER y PASSWORD
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	// Validar si las variables de entorno están presentes
	if user == "" || password == "" {
		fmt.Println("Las variables de entorno USER y PASSWORD no están configuradas correctamente")
		os.Exit(1)
	}

	// Imprimir las variables de entorno
	fmt.Println("USER:", user)
	fmt.Println("PASSWORD:", password)
}
