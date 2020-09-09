package  main

import (
	"fmt"
	"strings"
)

func main(){

	
	var x string
	fmt.Scan(&x)
	x=strings.ToLower(x)
	n :=len(x)
	if n<3{
		fmt.Println("Not Found")
	}else{
		flag:=0
		if x[0]=='i' && x[n-1]=='n'{
			for  i:=1; i<n; i++{
				if x[i]=='a'{
					flag=1
					break
				}
			}
		}
		if flag==0{
			fmt.Println("Not Found")
		}else{
			fmt.Println("Found")
		}
	}

}
