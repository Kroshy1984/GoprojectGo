// You can edit this code!
// Click here and start typing.
package main
import ("fmt" 
        "math"
		)

type uravnen struct {
    a,b,c,d float64 
    }
	
func (ur1 *uravnen) vvod() {
	fmt.Println(" Please input a="  )
	fmt.Scanf("%v \n", &ur1.a)
	fmt.Println(" Please input b="  )
	fmt.Scanf("%v \n", &ur1.b)
	fmt.Println(" Please input c="  )
	fmt.Scanf("%v \n", &ur1.c)
	}

func (ur1 *uravnen) calcdiscr() {
	ur1.d=ur1.b*ur1.b-4*ur1.a*ur1.c
	}

func (ur1 *uravnen) calcroot() {
    if ur1.d>0 {
	   fmt.Println("root № 1=", (-1*ur1.b-math.Pow(ur1.d,0.5))/(2*ur1.a))
	   fmt.Println("root № 2=", (-1*ur1.b+math.Pow(ur1.d,0.5))/(2*ur1.a))
	 }else if ur1.d==0 {
	   fmt.Println("root =", (-1*ur1.b-math.Pow(ur1.d,0.5))/(2*ur1.a))
	 }else if ur1.d<0 {
	   fmt.Println("There are not any roots ")
	 }
	}
		
func main() {
     uobj:= uravnen{}
	 uobj.vvod()
	 uobj.calcdiscr()
	 uobj.calcroot()
	}

