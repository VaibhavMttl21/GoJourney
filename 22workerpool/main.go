// package main

// import (
//     "fmt"
//     "sync"
// )

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {

//     defer wg.Done()

//     for job := range jobs {
//         fmt.Println("Worker", id, "processing job", job)
//     }
// }

// func main() {

//     jobs := make(chan int)

//     var wg sync.WaitGroup

//     // Start 3 workers
//     wg.Add(3)

//     go worker(1, jobs, &wg)
//     go worker(2, jobs, &wg)
//     go worker(3, jobs, &wg)

//     // Send jobs
//     for i := 1; i <= 10; i++ {
//         jobs <- i
//     }

//     // Tell workers there are no more jobs
//     close(jobs)

//     // Wait for workers to finish
//     wg.Wait()
// }

    //                 jobs
    //                  │
    //       ┌──────────┼──────────┐
    //       ▼          ▼          ▼
    //    Worker 1   Worker 2   Worker 3
    //       │          │          │
    //       └──────────┼──────────┘
    //                  │
    //               Done()
    //                  │
    //                  ▼
    //               WaitGroup
    //                  │
    //               counter = 0
    //                  │
    //                  ▼
    //              wg.Wait()
    //               finishes
    //                  │
    //                  ▼
    //              main ends

	package main

import (
    "fmt"
    "sync"
)

func worker(
    id int,
    jobs <-chan int,
    results chan<- int,
    wg *sync.WaitGroup,
) {

    defer wg.Done()

    for job := range jobs {

        fmt.Println("Worker", id, "processing job", job)

        result := job * job

        results <- result
    }
}

func main() {

    jobs := make(chan int)
    results := make(chan int)

    var wg sync.WaitGroup

    // Start workers
    wg.Add(3)

    
    // Send jobs
    go func() {
        
        for i := 1; i <= 10; i++ {
            jobs <- i
        }
        
        close(jobs)
        
        }()
        
        go worker(1, jobs, results, &wg)
        go worker(2, jobs, results, &wg)
        go worker(3, jobs, results, &wg)
    // Receive results
    for i := 1; i <= 10; i++ {

        result := <-results

        fmt.Println("Result:", result)
    }

    wg.Wait()
}