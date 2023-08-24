package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Zapato struct {
	Marca  string
	Precio int
	Talla  int
}

type Inventario struct {
	Zapato Zapato
	Stock  int
}

type Tienda struct {
	Inventario []Inventario
}

func (t *Tienda) AgregarZapato(inventario Inventario) {
	t.Inventario = append(t.Inventario, inventario)
}

func (t *Tienda) VenderZapato(marca string, talla int) error {
	for i, inv := range t.Inventario {
		if inv.Zapato.Marca == marca && inv.Zapato.Talla == talla {
			if inv.Stock > 0 {
				t.Inventario[i].Stock--
				return nil
			} else {
				return errors.New("Sin stock")
			}
		}
	}
	return errors.New("Zapato no encontrado")
}

func (t *Tienda) MostrarInventario() {
	for _, inv := range t.Inventario {
		fmt.Printf("Marca: %s\n", inv.Zapato.Marca)
		fmt.Printf("Precio: $%d\n", inv.Zapato.Precio)
		fmt.Printf("Talla: %d\n", inv.Zapato.Talla)
		fmt.Printf("Stock: %d\n", inv.Stock)
		fmt.Println()
	}
}

func main() {
	aleatorio := rand.New(rand.NewSource(time.Now().UnixNano()))

	marcas := []string{"Nike", "Adidas", "Puma", "Reebok", "Converse", "Vans", "New Balance", "Fila", "Under Armour", "Skechers"}

	tienda := Tienda{}

	for _, marca := range marcas {
		zapato := Zapato{
			Marca:  marca,
			Precio: aleatorio.Intn(10000),
			Talla:  aleatorio.Intn(11) + 34,
		}

		inventario := Inventario{
			Zapato: zapato,
			Stock:  aleatorio.Intn(50),
		}

		tienda.AgregarZapato(inventario)
	}

	tienda.MostrarInventario()

	fmt.Println()

	err := tienda.VenderZapato("Converse", 40)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("¡Vendido!")
	}

	err = tienda.VenderZapato("Nike", 42)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("¡Vendido!")
	}

	err = tienda.VenderZapato("Puma", 37)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("¡Vendido!")
	}

	err = tienda.VenderZapato("Under Armour", 38)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("¡Vendido!")
	}

	fmt.Println()

	tienda.MostrarInventario()

	fmt.Println()

	err = tienda.VenderZapato("Puma", 37)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("¡Vendido!")
	}
}
