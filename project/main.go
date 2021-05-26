// You can edit this code!
// Click here and start typing.
package main
//4444
import ("fmt" 
        "math"
		)

func main() {
     var a,b,c,d float64
	 	 fmt.Println(" Please input a=")
	 fmt.Scanf("%v \n", &a)
	 	 fmt.Println("Please input b=")
	 fmt.Scanf("%v \n", &b)
	 	 fmt.Println("Please input c=")
	 fmt.Scanf("%v \n", &c)
	  	 d=b*b-4*a*c
	 	 fmt.Println("\n")
	 if d>0 {
	   fmt.Println("root № 1=", (-1*b-math.Pow(d,0.5))/(2*a))
	   fmt.Println("root № 2=", (-1*b+math.Pow(d,0.5))/(2*a))
	 }else if d==0 {
	   fmt.Println("root =", (-1*b-math.Pow(d,0.5))/(2*a))
	 }else if d<0 {
	   fmt.Println("There are not any roots ")
	 }
	     
}