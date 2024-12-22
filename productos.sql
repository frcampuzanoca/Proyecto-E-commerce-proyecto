-- Crear la base de datos mi_tienda si no existe
CREATE DATABASE IF NOT EXISTS mi_tienda
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_general_ci;

-- Seleccionar la base de datos mi_tienda
USE mi_tienda;

-- Crear la tabla productos si no existe
CREATE TABLE IF NOT EXISTS productos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    categoria VARCHAR(100) NOT NULL,
    precio DECIMAL(10, 2) NOT NULL,
    descripcion TEXT,
    stock INT DEFAULT 0,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insertar productos de ejemplo
INSERT INTO productos (nombre, categoria, precio, descripcion, stock) VALUES
('Laptop Lenovo IdeaPad 3', 'Tecnología', 13500.00, 'Laptop con procesador Intel i5 y 8GB de RAM.', 10),
('Laptop Dell XPS 13', 'Tecnología', 32000.00, 'Laptop ultraligera con pantalla 4K.', 5),
('Smartphone Xiaomi Redmi Note 10', 'Tecnología', 6800.00, 'Smartphone con cámara de 48MP y batería de 5000mAh.', 25),
('Smartphone iPhone 13', 'Tecnología', 22000.00, 'iPhone con chip A15 Bionic y cámara avanzada.', 15),
('Tablet Samsung Galaxy Tab S7', 'Tecnología', 11500.00, 'Tablet con pantalla AMOLED y S Pen.', 8),
('Monitor LG UltraWide 29"', 'Tecnología', 7500.00, 'Monitor UltraWide de 29 pulgadas con resolución 2560x1080.', 12),
('Audífonos Sony WH-1000XM4', 'Tecnología', 6500.00, 'Audífonos con cancelación de ruido líder en la industria.', 20),
('Bocina JBL Flip 6', 'Tecnología', 3000.00, 'Bocina portátil resistente al agua con sonido potente.', 30),
('Smartwatch Apple Watch SE', 'Tecnología', 7200.00, 'Smartwatch con múltiples funciones de salud y conectividad.', 18),
('Disco Duro Externo WD 1TB', 'Tecnología', 1500.00, 'Disco duro externo portátil con 1TB de almacenamiento.', 25),

('Nike Air Max 270', 'Zapatos', 2900.00, 'Zapatos deportivos con excelente amortiguación.', 50),
('Adidas Ultraboost 21', 'Zapatos', 3200.00, 'Zapatos running con máxima comodidad y soporte.', 40),
('Converse Chuck Taylor All Star', 'Zapatos', 1200.00, 'Clásicos zapatos de lona para estilo casual.', 60),
('Vans Old Skool', 'Zapatos', 1300.00, 'Zapatos de skate con diseño icónico.', 45),
('Puma RS-X3', 'Zapatos', 2200.00, 'Zapatos deportivos con diseño moderno y colores vibrantes.', 35),

('Cacerola T-fal 24 cm', 'Cocina', 699.00, 'Cacerola antiadherente con tapa de vidrio.', 30),
('Juego de cuchillos Tramontina', 'Cocina', 899.00, 'Juego de cuchillos de acero inoxidable de alta calidad.', 20),
('Set de sartenes antiadherentes', 'Cocina', 1199.00, 'Set completo de sartenes antiadherentes para todo tipo de cocina.', 15),
('Batería de cocina de acero inoxidable', 'Cocina', 2399.00, 'Batería de cocina duradera y resistente al desgaste.', 10),
('Procesador de alimentos Ninja', 'Cocina', 2999.00, 'Procesador multifunción para preparar alimentos rápidamente.', 12),
('Licuadora Oster Pro', 'Cocina', 1799.00, 'Licuadora potente ideal para batidos y sopas.', 25),
('Horno de microondas Panasonic', 'Cocina', 1599.00, 'Microondas con múltiples funciones de cocción.', 18),
('Freidora de aire Philips', 'Cocina', 2499.00, 'Freidora de aire para cocinar de manera saludable sin aceite.', 22),

('Camisa Polo Ralph Lauren', 'Ropa', 899.00, 'Camisa de algodón de alta calidad.', 40),
('Pantalón Levi\'s 501', 'Ropa', 1199.00, 'Pantalón clásico de mezclilla con corte cómodo.', 35),
('Sudadera Adidas Originals', 'Ropa', 1299.00, 'Sudadera deportiva con el logo de Adidas.', 30),
('Chaqueta de mezclilla', 'Ropa', 1499.00, 'Chaqueta de denim resistente y estilizada.', 25),
('Vestido Zara verano', 'Ropa', 799.00, 'Vestido ligero y fresco ideal para el verano.', 20),
('Camiseta básica H&M', 'Ropa', 199.00, 'Camiseta de algodón suave y cómoda.', 50),
('Chamarras The North Face', 'Ropa', 2499.00, 'Chamarra impermeable y resistente para actividades al aire libre.', 15),
('Blusa formal Calvin Klein', 'Ropa', 1099.00, 'Blusa elegante ideal para ocasiones formales.', 10);
