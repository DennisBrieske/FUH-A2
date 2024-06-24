package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64

	fmt.Println("Bitte geben Sie die drei Seitenlängen des Quaders ein:")

	// Werte abfragen
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		println("Fehler:", err)
		return
	}

	// Berechnungen
	volumen := a * b * c
	kantensumme := (4 * a) + (4 * b) + (4 * c)
	oberflaeche := (2 * (a * b)) + 2*(a*c) + (2 * (b * c))
	raumdiagonale := math.Sqrt(a*a + b*b + c*c)
	umkugelradius := 0.5 * raumdiagonale

	// Format setzten
	format := "%.3g"

	// Ausgabe
	fmt.Println("Ein", a, "x", b, "x", c, "Quader hat die geometrischen Größen:")
	fmt.Printf("Volumen: "+format+" \n", volumen)
	fmt.Printf("Kantensumme: "+format+" \n", kantensumme)
	fmt.Printf("Oberfläche: "+format+" \n", oberflaeche)
	fmt.Printf("Umkugelradius: "+format+" \n", umkugelradius)
	fmt.Printf("Raumdiagonale: "+format+" \n", raumdiagonale)
}
