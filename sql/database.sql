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
    imagen VARCHAR(500) NOT NULL -- URL de la imagen
    );

CREATE TABLE compras (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         nombre VARCHAR(255) NOT NULL,
                         correo VARCHAR(255) NOT NULL,
                         direccion TEXT NOT NULL,
                         total DECIMAL(10, 2) NOT NULL,
                         estado VARCHAR(50) DEFAULT 'pendiente',
                         fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     cantidad_items INT NOT NULL
);

CREATE TABLE productos_comprados (
                                     id INT AUTO_INCREMENT PRIMARY KEY,
                                     compra_id INT NOT NULL,
                                     nombre VARCHAR(255) NOT NULL,
                                     precio DECIMAL(10, 2) NOT NULL,
                                     cantidad INT NOT NULL,
                                     subtotal DECIMAL(10, 2) NOT NULL,
                                     FOREIGN KEY (compra_id) REFERENCES compras(id)
);

-- Insertar productos de ejemplo
INSERT INTO productos (nombre, categoria, precio, descripcion, stock, imagen) VALUES
                                                                                  ('Laptop Lenovo IdeaPad 3', 'Tecnología', 13500.00, 'Laptop con procesador Intel i5 y 8GB de RAM.', 10, 'https://comandato.vteximg.com.br/arquivos/ids/231412-1000-1000/100061973.jpg?v=638221839192130000'),
                                                                                  ('Laptop Dell XPS 13', 'Tecnología', 32000.00, 'Laptop ultraligera con pantalla 4K.', 5, 'https://i.dell.com/is/image/DellContent/content/dam/ss2/product-images/dell-client-products/notebooks/xps-notebooks/9345/media-gallery/touch/gray/notebook-xps-13-9345-t-gray-gallery-2.psd?fmt=pjpg&pscan=auto&scl=1&wid=3988&hei=2361&qlt=100,1&resMode=sharp2&size=3988,2361&chrss=full&imwidth=5000'),
                                                                                  ('Smartphone Xiaomi Redmi Note 10', 'Tecnología', 6800.00, 'Smartphone con cámara de 48MP y batería de 5000mAh.', 25, 'https://happy.ec/wp-content/uploads/2024/05/XIAOMI-REDMI-NOTE-10-NEGRO-1.png'),
                                                                                  ('Smartphone iPhone 13', 'Tecnología', 22000.00, 'iPhone con chip A15 Bionic y cámara avanzada.', 15, 'https://i5.walmartimages.com/seo/Straight-Talk-Apple-iPhone-13-128GB-Pink-Prepaid-Smartphone-Locked-to-Straight-Talk_96348bb0-5338-4dbc-9bdc-6b5f1bcdaf5a.8e6c4ae8acb4cb933ece20fd4317bb58.jpeg'),
                                                                                  ('Tablet Samsung Galaxy Tab S7', 'Tecnología', 11500.00, 'Tablet con pantalla AMOLED y S Pen.', 8, 'https://novicompu.vtexassets.com/arquivos/ids/162867-800-auto?v=637823656959470000&width=800&height=auto&aspect=true'),
                                                                                  ('Monitor LG UltraWide 29"', 'Tecnología', 7500.00, 'Monitor UltraWide de 29 pulgadas con resolución 2560x1080.', 12, 'https://www.lg.com/ec/images/monitores/md05856675/gallery/Global_29UM69G_Gallery_large01.jpg'),
                                                                                  ('Audífonos Sony WH-1000XM4', 'Tecnología', 6500.00, 'Audífonos con cancelación de ruido líder en la industria.', 20, 'https://www.sony.com.ec/image/5d02da5df552836db894cead8a68f5f3?fmt=pjpeg&wid=330&bgcolor=FFFFFF&bgc=FFFFFF'),
                                                                                  ('Bocina JBL Flip 6', 'Tecnología', 3000.00, 'Bocina portátil resistente al agua con sonido potente.', 30, 'https://maxitec.vteximg.com.br/arquivos/ids/191370-1000-1000/maxitec-jbl-parlante-inalambrico-jbl-flip-6-color-negro-jblflip6blkam-1.jpg?v=638304422419130000'),
                                                                                  ('Smartwatch Apple Watch SE', 'Tecnología', 7200.00, 'Smartwatch con múltiples funciones de salud y conectividad.', 18, 'https://example.com/images/https://m.media-amazon.com/images/I/61wjDV8p6PL._AC_SL1500_.jpg'),
                                                                                  ('Disco Duro Externo WD 1TB', 'Tecnología', 1500.00, 'Disco duro externo portátil con 1TB de almacenamiento.', 25, 'https://www.tecnosmart.com.ec/wp-content/uploads/2022/03/HDD0101.jpg')
