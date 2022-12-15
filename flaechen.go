package main

import "math"

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche des entsprechenden Rechtecks.
func AreaRectangle(a, b float64) float64 {
	return a * b
}

// Erwartet eine Seitenlänge a.
// Liefert die Fläche des entsprechenden Quadrats.
func AreaSquare(a float64) float64 {
	return AreaRectangle(a, a)
}

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func AreaRightTriangle(a, b float64) float64 {
	return AreaRectangle(a, b) / 2
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang des entsprechenden Rechtecks.
func PerimeterRectangle(a, b float64) float64 {
	return 2 * (a + b)
}

// Erwartet eine Seitenlänge a.
// Liefert den Umfang des entsprechenden Quadrats.
func PerimeterSquare(a float64) float64 {
	return PerimeterRectangle(a, a)
}

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	// Hinweis: Hierfür ist die Funktion
	//          math.Sqrt() nützlich.
	return math.Sqrt(a*a + b*b)
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	return a + b + Hypotenuse(a, b)
}
