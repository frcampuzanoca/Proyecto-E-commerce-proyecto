// database/connect.go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importar el controlador de MySQL
)

// Connect establece la conexión con la base de datos MySQL usando variables de entorno
func Connect() (*sql.DB, error) {
	usuario := os.Getenv("DB_USER")
	contrasena := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	nombreBase := os.Getenv("DB_NAME")

	// Verificar si faltan variables
	var missingVars []string

	if usuario == "" {
		missingVars = append(missingVars, "DB_USER")
	}
	if contrasena == "" {
		missingVars = append(missingVars, "DB_PASSWORD")
	}
	if host == "" {
		missingVars = append(missingVars, "DB_HOST")
	}
	if port == "" {
		missingVars = append(missingVars, "DB_PORT")
	}
	if nombreBase == "" {
		missingVars = append(missingVars, "DB_NAME")
	}

	// Si falta alguna variable, se imprime cuáles son y se finaliza la app
	if len(missingVars) > 0 {
		log.Fatalf("Faltan variables de entorno: %v", missingVars)
	}

	// Construir la cadena de conexión (DSN)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		usuario, contrasena, host, port, nombreBase,
	)

	// Abrir la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error al abrir la base de datos: %v", err)
		return nil, err
	}

	// Verificar la conexión con un ping
	if err := db.Ping(); err != nil {
		log.Printf("Error al hacer ping a la base de datos: %v", err)
		return nil, err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}

// InsertarProductos inserta productos de ejemplo en la tabla "productos"
func InsertarProductos(db *sql.DB) error {
	// Definir productos de ejemplo
	productos := []struct {
		Nombre      string
		Categoria   string
		Precio      float64
		Descripcion string
		Stock       int
	}{
		// Lista de productos...
		// (Incluye todos los productos que deseas insertar)
	}

	// Preparar el statement para inserción
	stmt, err := db.Prepare("INSERT INTO productos (nombre, categoria, precio, descripcion, stock) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insertar cada producto
	for _, p := range productos {
		_, err := stmt.Exec(p.Nombre, p.Categoria, p.Precio, p.Descripcion, p.Stock)
		if err != nil {
			log.Printf("Error al insertar producto %s: %v", p.Nombre, err)
			// Puedes decidir si continuar o retornar el error
			return err
		}
	}

	log.Println("¡Productos insertados correctamente!")
	return nil
}
