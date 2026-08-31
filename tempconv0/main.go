package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// String Method
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {

	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f) isso irá gerar um erro de compilação por conta dos tipos
	fmt.Println(c == Celsius(f))
	fmt.Println(c.String())
	fmt.Println(c)
}
