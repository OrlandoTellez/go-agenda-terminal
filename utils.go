package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LeerEntrada(mensaje string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(mensaje)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ContactosIniciales() []Contacto {
	return []Contacto{
		{
			nombre:   "Orlando",
			telefono: "1234-1234",
			email:    "orlando@unext.com",
		},
		{
			nombre:   "Miguel",
			telefono: "4321-1234",
			email:    "miguel@unext.com",
		},
		{
			nombre:   "Rapha",
			telefono: "4321-1444",
			email:    "rapha@unext.com",
		},
		{
			nombre:   "Jorge",
			telefono: "4365-1234",
			email:    "jorge@unext.com",
		},
	}
}
