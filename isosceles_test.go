//A Program to test and benchmark multiple values inserted into the program 
//@rachaelhogan

package main

//import "fmt"
import "testing"
//import "github.com/stretchr/testify/assert"
import "errors"
//TestIsosceles creates numerous test cases and tests each one for the expected ouput
var e =errors.New("All arguments must be above 0")
func TestIsosceles(t *testing.T) {  
	tests := []struct{
    		x int
   		y int
    		z int
    		b bool
    		err error
  	}{
		{1,2,3,false, nil },
		{3,3,3,false, nil },
   		{2,2,3,true, nil },
   		{1,3,3,true, nil},
		{3,1,3,true, nil},
		{-1,1,2,false,  &IsoscelesError{"All arguments must be above 0"}},
    		{1,0,1,false,  &IsoscelesError{"All arguments must be above 0"}},
    		{1,1,0,false,  &IsoscelesError{"All arguments must be above 0"}},
 	}
  	for _, test := range tests {
	   	x,v := Isosceles(test.x, test.y, test.z)
		 if(test.err == nil){
			if( x!=test.b || v!=test.err){
				t.Error(
        				"For", test.x,test.y,test.z,
           		   		"expected", test.b,test.err, 
                            		"got", x ,v,
         			 )	
     			 }
    		}
		if(test.err!=nil){
	  		if(x!=test.b){
				t.Error(
                            		"For", test.x,test.y,test.z,
                            		"expected", test.b,
	                           	"got", x, 
				)	
			}
		}
	}	
}

//benchmarkIsosceles runs through inputted test cases to calculate
//the benchmark value
func benchmarkIsosceles(x int, y int, z int, b *testing.B) {
        for n := 0; n < b.N; n++ {
                Isosceles(x,y,z)
        }
}

//these functions inputs multiple test cases into the benchmarkIsoscles function
func BenchmarkIsosceles1(b *testing.B) {benchmarkIsosceles(1,2,3,b) }
func BenchmarkIsosceles2(b *testing.B) {benchmarkIsosceles(3,3,3,b) }
func BenchmarkIsosceles3(b *testing.B) {benchmarkIsosceles(2,2,3,b) }
func BenchmarkIsosceles4(b *testing.B) {benchmarkIsosceles(1,2,2,b) }
func BenchmarkIsosceles5(b *testing.B) {benchmarkIsosceles(1,2,1,b) }
func BenchmarkIsosceles6(b *testing.B) {benchmarkIsosceles(0,2,3,b) }
func BenchmarkIsosceles7(b *testing.B) {benchmarkIsosceles(1,0,3,b) }
func BenchmarkIsosceles8(b *testing.B) {benchmarkIsosceles(1,2,0,b) } 

