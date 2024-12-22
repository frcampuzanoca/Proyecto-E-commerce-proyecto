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

// ListarBasesDeDatos lista todas las bases de datos disponibles
func ListarBasesDeDatos(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		return nil, fmt.Errorf("error al obtener las bases de datos: %v", err)
	}
	defer rows.Close()

	var basesDeDatos []string
	for rows.Next() {
		var baseDeDatos string
		if err := rows.Scan(&baseDeDatos); err != nil {
			return nil, fmt.Errorf("error al leer las bases de datos: %v", err)
		}
		basesDeDatos = append(basesDeDatos, baseDeDatos)
	}

	return basesDeDatos, nil
}

// InsertarProductos inserta productos de ejemplo en la tabla "productos".
func InsertarProductos(db *sql.DB) error {
	// 10 productos de tecnología
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

	// 5 productos de zapatos
	shoeProducts := []struct {
		nombre    string
		categoria string
		precio    float64
	}{
		{"Nike Air Max 270", "Zapatos", 2900.00},
		{"Adidas Ultraboost 21", "Zapatos", 3200.00},
		{"Converse Chuck Taylor All Star", "Zapatos", 1200.00},
		{"Vans Old Skool", "Zapatos", 1300.00},
		{"Puma RS-X3", "Zapatos", 2200.00},
	}

	// 8 productos de cocina
	cocinaProducts := []struct {
		nombre    string
		categoria string
		precio    float64
	}{
		{"Cacerola T-fal 24 cm", "Cocina", 699.00},
		{"Juego de cuchillos Tramontina", "Cocina", 899.00},
		{"Set de sartenes antiadherentes", "Cocina", 1199.00},
		{"Batería de cocina de acero inoxidable", "Cocina", 2399.00},
		{"Procesador de alimentos Ninja", "Cocina", 2999.00},
		{"Licuadora Oster Pro", "Cocina", 1799.00},
		{"Horno de microondas Panasonic", "Cocina", 1599.00},
		{"Freidora de aire Philips", "Cocina", 2499.00},
	}

	// 8 productos de ropa
	ropaProducts := []struct {
		nombre    string
		categoria string
		precio    float64
	}{
		{"Camisa Polo Ralph Lauren", "Ropa", 899.00},
		{"Pantalón Levi's 501", "Ropa", 1199.00},
		{"Sudadera Adidas Originals", "Ropa", 1299.00},
		{"Chaqueta de mezclilla", "Ropa", 1499.00},
		{"Vestido Zara verano", "Ropa", 799.00},
		{"Camiseta básica H&M", "Ropa", 199.00},
		{"Chamarras The North Face", "Ropa", 2499.00},
		{"Blusa formal Calvin Klein", "Ropa", 1099.00},
	}

	// Insertar los productos de tecnología
	for _, p := range techProducts {
		if _, err := db.Exec(
			"INSERT INTO productos (nombre, categoria, precio) VALUES (?, ?, ?)",
			p.nombre, p.categoria, p.precio,
		); err != nil {
			return fmt.Errorf("error al insertar producto de tecnología (%s): %v", p.nombre, err)
		}
	}

	// Insertar los productos de zapatos
	for _, p := range shoeProducts {
		if _, err := db.Exec(
			"INSERT INTO productos (nombre, categoria, precio) VALUES (?, ?, ?)",
			p.nombre, p.categoria, p.precio,
		); err != nil {
			return fmt.Errorf("error al insertar producto de zapatos (%s): %v", p.nombre, err)
		}
	}

	// Insertar los productos de cocina
	for _, p := range cocinaProducts {
		if _, err := db.Exec(
			"INSERT INTO productos (nombre, categoria, precio) VALUES (?, ?, ?)",
			p.nombre, p.categoria, p.precio,
		); err != nil {
			return fmt.Errorf("error al insertar producto de cocina (%s): %v", p.nombre, err)
		}
	}

	// Insertar los productos de ropa
	for _, p := range ropaProducts {
		if _, err := db.Exec(
			"INSERT INTO productos (nombre, categoria, precio) VALUES (?, ?, ?)",
			p.nombre, p.categoria, p.precio,
		); err != nil {
			return fmt.Errorf("error al insertar producto de ropa (%s): %v", p.nombre, err)
		}
	}

	fmt.Println("¡Productos insertados correctamente!")
	return nil
}
