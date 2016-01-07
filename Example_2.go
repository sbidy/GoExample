package main

import(
        "fmt"
        "time"
        "math/rand"
        )

func main()  {
    average := 0
    fmt.Println("Create channel")
    dbconnection := make(chan string)
      
    fmt.Println("Starten der goroutine")
    for a := 0; a < 10; a++ {
        rand.Seed(time.Now().UnixNano())
        start := time.Now()
        go func() { dbconnection <- connectDB("Database 1") } ()
        go func() { dbconnection <- connectDB("Database 2") } ()
        
        for i := 0; i < 2; i++ {
            fmt.Println(<- dbconnection)
        }
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
