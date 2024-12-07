package models

import (
	"errors"
	"strings"
)

// Usuario representa un usuario genérico del sistema
type Usuario struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
	Rol        string `json:"rol"` // "Cliente", "Administrador"
}

// Devolvemos y asignamos el nombre del usuario
func (u *Usuario) GetNombre() string {
	return u.Nombre
}

func (u *Usuario) SetNombre(nombre string) {
	u.Nombre = nombre
}

// GetRol y SetRol devuelve y asigna el rol del usuario
func (u *Usuario) GetRol() string {
	return u.Rol
}

func (u *Usuario) SetRol(rol string) {
	u.Rol = rol
}

// GetContrasena y SetContrasena asigna una nueva contraseña al usuario
func (u *Usuario) GetContrasena() string {
	return u.Contrasena
}

func (u *Usuario) SetContrasena(contrasena string) {
	u.Contrasena = contrasena
}

// Validación para que el correo no este vacío.
func (u *Usuario) ValidacionCorreoVacio(correo string) error {
	if correo == "" {
		return errors.New("El correo no puede estar vacío")
	}
	u.Correo = correo
	return nil
}

// Validación del @ en el correo
func (u *Usuario) validacionCorreo(correo string) error {
	if !strings.Contains(correo, "@") {
		return errors.New("El correo debe contener un @")
	}
	u.Correo = correo
	return nil
}

// Validamos que la contraseña tenga 6 carácteres como mínimo
func (u *Usuario) validarContrasena(contrasena string) error {
	if len(contrasena) < 6 {
		return errors.New("La contraseña debe tener al menos 6 carácteres")
	}
	u.Contrasena = contrasena
	return nil
}

// Cliente extiende Usuario
type Cliente struct {
	Usuario          // Composición
	Pedidos []Pedido `json:"pedidos"`
}

// Administrador extiende Usuario
type Administrador struct {
	Usuario // Composición
}
