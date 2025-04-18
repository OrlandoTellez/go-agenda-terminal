package main

import "fmt"

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
