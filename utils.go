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
