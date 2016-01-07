package main

import(
        "fmt"
        "time"
        "math/rand"
        )

func main()  {
    
    defer fmt.Println("Ready!!")
    average := 0
    fmt.Println("Creat channel")
    dbconnection := make(chan string)
      
    fmt.Println("Start goroutine")
    for a := 0; a < 10; a++ {
        rand.Seed(time.Now().UnixNano())
        start := time.Now()
        go func() { dbconnection <- connectDB("Database 1") } ()
        go func() { dbconnection <- connectDB("Database 2") } ()

        for i := 0; i < 2; i++ {
            select {
                case data := <-dbconnection:
                    fmt.Println(data)
                case <-time.After(500 * time.Millisecond):
                    fmt.Println("Timeout!")
                    break
                }
        }
        average = average + (int(time.Since(start))/(1000*1000))
        fmt.Println("  Next try")
    }
    fmt.Print("Time in MS: ")
    fmt.Print(average/10)
    fmt.Println("")
}

func connectDB(db string) string{
    // DB Zugriff
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    return "Antwort von "+ db
}
