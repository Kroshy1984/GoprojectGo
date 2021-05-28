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
    var x float64
	fmt.Println(" Please input a="  )
	fmt.Scanf("%v \n", &x)
    (*ur1).a=x
	fmt.Println(" Please input b="  )
	fmt.Scanf("%v \n", &x)
    (*ur1).b=x
	fmt.Println(" Please input c="  )
	fmt.Scanf("%v \n", &x)
    (*ur1).c=x
	}

func (ur1 *uravnen) calcdiscr() {
    var x float64
	x=(*ur1).b*(*ur1).b-4*(*ur1).a*(*ur1).c
	(*ur1).d=x
	}

func (ur1 *uravnen) calcroot() {
    if (*ur1).d>0 {
	   fmt.Println("root № 1=", (-1*(*ur1).b-math.Pow((*ur1).d,0.5))/(2*(*ur1).a))
	   fmt.Println("root № 2=", (-1*(*ur1).b+math.Pow((*ur1).d,0.5))/(2*(*ur1).a))
	 }else if (*ur1).d==0 {
	   fmt.Println("root =", (-1*(*ur1).b-math.Pow((*ur1).d,0.5))/(2*(*ur1).a))
	 }else if (*ur1).d<0 {
	   fmt.Println("There are not any roots ")
	 }
	}
		
func main() {
     uobj:= uravnen{}
	 uobj.vvod()
	 uobj.calcdiscr()
	 uobj.calcroot()
	}

