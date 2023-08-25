package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
}

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) {
	var err = -1
	var p *producto = nil
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			err = 0
			p = &((*l)[i])
		}
	}
	return p, err
}

func (l listaProductos) venderProducto(nombre string, cant int) {
	var prod, err = l.buscarProducto(nombre)

	if err != -1 {
		(*prod).cantidad -= cant
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 13, 2000)
	lProductos.agregarProducto("leche", 10, 1200)
	lProductos.agregarProducto("café", 16, 4500)
}

func (l *listaProductos) listarProductosMinimos() listaProductos {
	newSlice := make(listaProductos, 0)

	for _, p := range *l {
		if p.cantidad <= existenciaMinima {
			newSlice = append(newSlice, p)
		}
	}
	return newSlice
}

func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for i, p := range listaMinimos {
		neededQuantity := existenciaMinima - p.cantidad
		if neededQuantity > 0 {
			(*l)[i].cantidad += neededQuantity
		}
	}
}

func (l *listaProductos) ordenarPorPrecioAscendente() {
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].precio < (*l)[j].precio
	})
}

func main() {
	llenarDatos()
	fmt.Println("Inventario Inicial:")
	fmt.Println(lProductos)

	lProductos.venderProducto("arroz", 11)
	lProductos.venderProducto("frijoles", 10)
	lProductos.venderProducto("leche", 8)
	lProductos.venderProducto("café", 12)
	fmt.Println("\nInventario después de ventas:")
	fmt.Println(lProductos)

	pminimos := lProductos.listarProductosMinimos()
	fmt.Println("\nProductos con existencia mínima:")
	fmt.Println(pminimos)

	lProductos.aumentarInventarioDeMinimos(pminimos)
	fmt.Println("\nInventario después de aumentar existencias:")
	fmt.Println(lProductos)

	lProductos.ordenarPorPrecioAscendente()
	fmt.Println("\nInventario ordenado por precio ascendente:")
	fmt.Println(lProductos)
}
