func add2Numbers(a, b int) { 
    fmt.Println( a + b )
}

func main() {
 add2Numbers(1, 2) //a normal function call
 
 go add2Numbers(1, 2) //a normal function executed parallely as a goroutine
}
