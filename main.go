package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contacto struct {
	nombre   string
	telefono string
	email    string
}

type Agenda struct {
	contacto []Contacto
}

func (a *Agenda) RegistrarContacto(nombre, telefono, email string) {
	nuevoContacto := Contacto{
		nombre:   nombre,
		telefono: telefono,
		email:    email,
	}

	a.contacto = append(a.contacto, nuevoContacto)
	fmt.Println(nuevoContacto.nombre, "ha sido registrado exitosamente en tus contactos")
}

func (a *Agenda) BuscarContacto(nombre string) {
	for _, c := range a.contacto {
		if nombre == c.nombre {
			fmt.Println("--- Contacto encontrado ---")
			fmt.Println("Nombre: ", c.nombre)
			fmt.Println("Teléfono: ", c.telefono)
			fmt.Println("Email: ", c.email)
			return
		}
	}

	fmt.Println("--- Contacto no encontrado ---")
}

func (a *Agenda) MostrarContactos() {
	fmt.Println("Lista de contactos")
	for i, c := range a.contacto {
		if len(a.contacto) > 0 {
			fmt.Println(i, "-", c.nombre)
		} else {
			fmt.Println("no hay contactos")
		}
	}
}

func (a *Agenda) BorrarContactos(nombre string) {
	encontrado := false
	var nuevoContacto []Contacto

	for _, c := range a.contacto {
		if c.nombre != nombre {
			nuevoContacto = append(nuevoContacto, c)
		} else {
			encontrado = true
		}
	}

	if encontrado {
		a.contacto = nuevoContacto
		fmt.Println("Contacto eliminado")
	} else {
		fmt.Println("No se encontro el contacto")
	}

}

func LeerEntrada(mensaje string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(mensaje)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	agenda := Agenda{}

	for {
		fmt.Println("-------- Agenda --------")
		fmt.Println("1. Registrar contacto")
		fmt.Println("2. Buscar contacto")
		fmt.Println("3. Ver contactos")
		fmt.Println("4. Borrar contactos")
		fmt.Println("5. Salir")

		opcion := LeerEntrada("Elija una opción: ")

		switch opcion {
		case "1":
			nombre := LeerEntrada("Ingrese nombre: ")
			telefono := LeerEntrada("Ingrese telefono: ")
			email := LeerEntrada("Ingrese email: ")

			agenda.RegistrarContacto(nombre, telefono, email)
		case "2":
			nombre := LeerEntrada("Ingrese nombre: ")

			agenda.BuscarContacto(nombre)
		case "3":
			agenda.MostrarContactos()
		case "4":
			nombre := LeerEntrada("Ingrese nombre: ")
			agenda.BorrarContactos(nombre)
		case "5":
			fmt.Println("¡Hasta la proxima!")
			return
		default:
			fmt.Println("Escoge una de las opciones del menu")
		}
	}

}
