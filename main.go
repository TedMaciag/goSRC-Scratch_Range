package main
  
import "fmt"


func main(){
	ranger(-3, -2)
	}

func ranger (low int, hi int){
	// This ensures order doesn't really matter 
	if low > hi{
		low, hi = hi, low
	}
	for i := low; i <= hi; i++ {
		fmt.Println(i)
		}
}
