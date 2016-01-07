package main

import(
        "fmt"
        "time"
        "math/rand"
        )

func main()  {
    average := 0
    for a := 0; a < 10; a++ {
        rand.Seed(time.Now().UnixNano())
        start := time.Now()
        fmt.Println(connectDB("Database 1"))
        fmt.Println(connectDB("Database 2"))
        average = average + (int(time.Since(start))/(1000*1000))
        fmt.Println("  Next try")
    }
    fmt.Print("Time in MS: ")
    fmt.Print(average/10)
}
    
func connectDB(db string) string{
    // DB access 
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    return "Response from "+ db
}
