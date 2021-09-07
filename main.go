package main
 
import (
    "fmt"
  )
var rows, menu int
var star string = "* "
 
func pyramid(rows int) string{
var count int = rows+1 //this variable is to help making the space
//Outer loop, this loop is for repeating the rows desired
  for i := 0; i <= rows; i++ {
  	//inner loop, this is when the count needed, to create space between star printed. Example, now we are in first inner loop, count = rows+1, this would print "  " by rows+1 
		for space := 0; space < count; space++ {
			fmt.Print("  ")
		}
    //the count is decremented so it would control the space printed. In the second inner loop, the space printed would be rows, third loop rows-1, and so on...
    count -= 1
		for j:=0;j<2*i-1;j++{
      //now this is the deepest loop, it uses limit as 2*i-1 to print stars unevenly.
			fmt.Print(star)
		}
		fmt.Println("")
  }
return "Program finished"
}

func inverted(rows int) string{
var count int = rows+1
  //this is the inverted version of pyramid. Outer loop starts with the biggest parameter and will stop in 0 rows
  for i := rows; i >= 0; i-- {
    //inner loop also starts with the biggest parameter but it would stop if space < count. For example we are in first inner loop, because space = rows and count = rows +1, then it would only print one "  "
		for space := rows; space >= count; space-- {
			fmt.Print("  ")
		}
    count -= 1
		for j:=0;j<2*i-1;j++{
			fmt.Print(star)
		}
		fmt.Println("")
  }
return "Program finished"
}

func main() {
  fmt.Print("Welcome to pyramid maker. Please select your menu (input number) :\n1. Normal Pyramid\n2. Inverted Pyramid\n")
  fmt.Scan(&menu)
  fmt.Println("How many rows do you want?")
  fmt.Scan(&rows)
  if menu == 1{
    fmt.Println(pyramid(rows))
  } else if menu == 2{
    fmt.Println(inverted(rows))
  } else {
    fmt.Println("Sorry, please input again")
  }  
}