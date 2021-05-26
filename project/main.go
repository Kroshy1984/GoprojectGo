// You can edit this code!
// Click here and start typing.
package main
import ("fmt" 
        "math"
		)

func main() {
     var a,b,c,d float64
	 	 a=vvod("a")
		 b=vvod("b")
		 c=vvod("c")
		 d=calcdiscr(a,b,c)
	 	 fmt.Println("\n")
		 calcroot(a,b,d)
		}


func vvod(pole string) float64 {
    var x float64
	fmt.Println(" Please input ",pole,"="  )
	 fmt.Scanf("%v \n", &x)
	 return x
}

func calcdiscr(a,b,c float64) float64 {
    var x float64
	x=b*b-4*a*c
	 return x
}

func calcroot(a,b,d float64)  {
    if d>0 {
	   fmt.Println("root № 1=", (-1*b-math.Pow(d,0.5))/(2*a))
	   fmt.Println("root № 2=", (-1*b+math.Pow(d,0.5))/(2*a))
	 }else if d==0 {
	   fmt.Println("root =", (-1*b-math.Pow(d,0.5))/(2*a))
	 }else if d<0 {
	   fmt.Println("There are not any roots ")
	 }
	}
