// Avance del Autonomo#3
//Nombres: Cristhian Sare, Freddy Campuzano, Iriany
/*Descripción: Avance pequeño donde podemos observar la estructura de los
productos y del usuario junto con la función getProduct*/
//Materia: Programación orientada a objetos.

package main

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
	missingVars := []string{}
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

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error al abrir la base de datos: %v", err)
		return nil, err
	}

	// Verificar la conexión con la base de datos
	if err := db.Ping(); err != nil {
		log.Printf("Error al hacer ping a la base de datos: %v", err)
		return nil, err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}

// InsertarProductos inserta productos de ejemplo en la tabla "productos".
func InsertarProductos(db *sql.DB) error {
	// Productos de ejemplo
	techProducts := []struct {
		nombre    string
		categoria string
		precio    float64
	}{
		{"Laptop Lenovo IdeaPad 3", "Tecnologia", 13500.00},
		{"Laptop Dell XPS 13", "Tecnologia", 32000.00},
		{"Smartphone Xiaomi Redmi Note 10", "Tecnologia", 6800.00},
		{"Smartphone iPhone 13", "Tecnologia", 22000.00},
		{"Tablet Samsung Galaxy Tab S7", "Tecnologia", 11500.00},
		{"Monitor LG UltraWide 29\"", "Tecnologia", 7500.00},
		{"Audífonos Sony WH-1000XM4", "Tecnologia", 6500.00},
		{"Bocina JBL Flip 6", "Tecnologia", 3000.00},
		{"Smartwatch Apple Watch SE", "Tecnologia", 7200.00},
		{"Disco Duro Externo WD 1TB", "Tecnologia", 1500.00},
	}

	// Insertar productos de ejemplo
	for _, p := range techProducts {
		if _, err := db.Exec(
			"INSERT INTO productos (nombre, categoria, precio) VALUES (?, ?, ?)",
			p.nombre, p.categoria, p.precio,
		); err != nil {
			return fmt.Errorf("error al insertar producto de tecnología (%s): %v", p.nombre, err)
		}
	}

	fmt.Println("¡Productos insertados correctamente!")
	return nil
}

func main() {
	// Configurar las variables de entorno (en producción puedes definirlas en tu sistema)
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "tu_contraseña")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "mi_tienda")

	// Conectar a la base de datos
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Insertar productos
	err = InsertarProductos(db)
	if err != nil {
		log.Fatal(err)
	}
}
