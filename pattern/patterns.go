package pattern

import "fmt"

func PrintWorld(){
	fmt.Println("hello world")
}

func OnetoTen(){
	for i:=0; i<=10; i++{
		fmt.Println(i)
   
	}
}

func Pattern1(){
	for i:=1; i<=6; i++{
		for j:=1; j<i; j++{
			fmt.Printf("%d ",j)
		}
		fmt.Println("")
	}
}

func Pattern2(){
	for i:=5; i>=1; i--{
		for j:=1; j<=i; j++{
			fmt.Printf("%d ",j)
		}
		fmt.Println("") 
	}
}

 func Pattern3(){
 	for i:=1; i<=5; i++{
 		for j:=1;j<=5; j++{
			fmt.Printf(" %d ",i*j)
			}
			fmt.Println()
 		}

 	}


func Pat4() {
	var  c, count1, k int = 0, 0, 0

	for i := 1; i <= 5; i++ {
		k = 0
		for j := 1; j <= 5-i; j++ {
			fmt.Print("  ")
			c++
		}
		for {
			if k == 2*i-1 {
				break
			}
			if c <= 4 {
				fmt.Printf("%d ", i+k)
				c++
			} else {
				count1++
				fmt.Printf("%d ", (i + k - 2*count1))
			}
			k++
		}
		count1, k, c = 0, 0, 0
		fmt.Println("")
	}
}

func Pat5(){
	var temp int = 1

	for i := 0; i < 7; i++ {	

		for j := 0; j <= i; j++ {

			if (j==0 || i==0) {
					temp = 1
				}else{
					temp = temp*(i-j+1)/j 
				}

			fmt.Printf(" %d",temp)				
		}
		fmt.Println("")
		
	}

}
