package main

import "fmt"

func main() {
	agenda := Agenda{ContactosIniciales()}

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
