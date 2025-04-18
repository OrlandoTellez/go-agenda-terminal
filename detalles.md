
## 📝 **Ejercicio: Sistema de Agenda de Contactos (en consola)**

### 🎯 Objetivo:
Crear un programa que permita:
- Registrar nuevos contactos
- Buscar contactos por nombre
- Mostrar todos los contactos
- Eliminar un contacto
- Salir

---

### 🔸 Estructura esperada

```go
type Contacto struct {
	nombre  string
	telefono string
	email   string
}

type Agenda struct {
	contactos []Contacto
}
```

---

### 🔹 Funcionalidades que debés implementar

1. `AgregarContacto(nombre, telefono, email string)`
   - Guarda un contacto nuevo en la agenda

2. `BuscarContacto(nombre string)`
   - Busca y muestra el contacto si existe

3. `MostrarContactos()`
   - Muestra todos los contactos almacenados

4. `EliminarContacto(nombre string)`
   - Elimina un contacto por nombre (si existe)

---

### 🔻 Interfaz en consola esperada (en `main()`)

```plaintext
--- Agenda de Contactos ---
1. Agregar contacto
2. Buscar contacto
3. Ver todos los contactos
4. Eliminar contacto
5. Salir
```
