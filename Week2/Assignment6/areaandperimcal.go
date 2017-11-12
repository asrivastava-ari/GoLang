/*
By using structs and methods in Go implement a circle, rectangle and calculate their areas.
Extend the above by introducing an “interface” having methods “area” (which calculates area)
and “perimeter” (which calculates perimeter)

*/

package main

import (
	"fmt"
	"math"
)


type GetSize interface{
	area() float64
	perimeter() float64
}

type Rectangle struct{
	len float64
	wid float64
}

type Circle struct{
	rad float64
}

func(circle Circle) area() float64{
	return math.Pi * circle.rad * circle.rad
} 

func(rectangle Rectangle) area() float64{
	return rectangle.len*rectangle.wid
} 

func (circle Circle) perimeter() float64{
	return 2*math.Pi*circle.rad
}

func (rectangle Rectangle) perimeter() float64{
	return 2 * (rectangle.len+ rectangle.wid)
}

func getArea(properties GetSize) float64{
	return properties.area()
}

func getPerimeter(properties GetSize) float64{
	return properties.perimeter()
}


func main(){
	circle := Circle{rad:5}
	rectangle := Rectangle{len:10, wid:20}
	
	fmt.Printf("\n Circle :: Area:%f,Perimiter:%f\n",getArea(circle),getPerimeter(circle))
	fmt.Printf("\n Rectangle :: Area:%f,Perimiter:%f\n",getArea(rectangle),getPerimeter(rectangle))

}