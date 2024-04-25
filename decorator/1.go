package decorator

import "fmt"

type CoffeeMaker interface {
	Cost() float32
	Description() string
}

type Coffee struct{}

func (c *Coffee) Cost() float32 {
	return 4.00
}

func (c *Coffee) Description() string {
	return "This is a basic coffee"
}

// Decorator 1
type Milk struct {
	coffee CoffeeMaker
}

func (m *Milk) Cost() float32 {
	return m.coffee.Cost() + 1.00
}

func (m *Milk) Description() string {
	return m.coffee.Description() + " with Milk"
}

// Sugar Decorator
type Sugar struct {
	coffee CoffeeMaker
}

func (s *Sugar) Cost() float32 {
	return s.coffee.Cost() + 2.00
}
func (s *Sugar) Description() string {
	return s.coffee.Description() + " with Sugar"
}

func Decorator1() {
	coffee := &Coffee{}
	coffeeWithMilk := &Milk{coffee: coffee}
	// with milk
	fmt.Println(coffeeWithMilk.Cost())
	fmt.Println(coffeeWithMilk.Description())

	// with sugar and milk
	coffeeWithMilkAndSugar := &Sugar{coffee: &Milk{coffee: coffee}}
	fmt.Println(coffeeWithMilkAndSugar.Description())
	fmt.Println(coffeeWithMilkAndSugar.Cost())
}
