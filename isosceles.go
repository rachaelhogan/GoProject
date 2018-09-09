//This program is designed to take in 3 inputs and determine if it is an isosceles triange
//@rachaelhogan

package main
import "github.com/fatih/color"

//The Isosceles function which determines if two integers are equal and all integers are above 0 
func Isosceles(x int,y int,z int) (bool,error) {
	err := &IsoscelesError{"All arguments must be above 0"}
	if (x<=0||z<=0||y<=0) {
		return false,err	
	}
	var found bool = false
	if (x==y){
		found=true	
	}
	if (y==z){
		return (!found),nil
	}
	if (x==z){
		return(!found),nil
	}
 return found, nil
}

//the creation of the IsoscelesError
type IsoscelesError struct{
        s string
}
func (e *IsoscelesError) Error() string {
    	return e.s 
}

//the main function, uses the color package and calls the Isosceles function
func main(){
	x:=3
	y:=4
	z:=3
 	if r, e := Isosceles(x,y,z); 
	e != nil {
	    red := color.New(color.FgRed)
            red.Println("Isosceles function failed:", e)
        } else {
		if(r){
			white:=color.New(color.FgWhite).Add(color.Bold)
			blueBackground:= white.Add(color.BgBlue)	
			blueBackground.Println("This is an Isosceles Triangle")
		}else{
	  		magenta:= color.New(color.FgMagenta).Add(color.Underline)	
	    		magenta.Println("Not an Isosceles Triangle")
		}
     	}
}
