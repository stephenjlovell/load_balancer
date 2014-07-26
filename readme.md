## Load Balancer

A simple load balancer based on a [presentation](https://www.youtube.com/watch?v=jgVhBThJdXc) by Rob Pike at Google I/O 2010

The main function creates 100 concurrent requester processes that send requests to the load balancer at pseudorandom intervals. Each request contains a channel back to the requesting process, and a closure specifying the work to be done.

The load balancer manages a set of worker objects organized into a priority queue based on how many tasks are assigned to each worker.  When a new request arrives, the load balancer assignes the task to the worker with the least items already assigned to it.  

When a worker completes a task, it sends the results back to process that requested it, and signals the load balancer via a channel to adjust the ordering of the worker priority queue.

The main function lets the requesters run for 10 seconds, periodically printing out the number of tasks assigned to each worker.

### Example output

    > go build
    > go install
    > load_balancer


    Running on 8 processors.

    15  16  15  15  17  16  | 94  
    15  16  14  16  15  17  | 93  
    15  15  16  16  16  15  | 93  
    15  16  15  16  16  15  | 93  
    15  15  15  16  16  15  | 92  
    15  15  14  15  18  16  | 93  
    13  16  17  16  14  17  | 93  
    17  17  17  16  15  18  | 100  
    12  14  12  17  18  14  | 87  
    17  17  17  17  17  | 85  
    15  16  16  16  17  16  | 96  
    15  15  16  16  16  16  | 94  
    16  17  17  16  17  17  | 100  
    15  16  17  17  17  15  | 97  
    14  15  15  16  16  16  | 92  
    15  16  15  16  15  15  | 92  
    13  16  14  18  18  17  | 96  
    16  15  16  16  16  16  | 95  
    15  18  16  18  15  16  | 98  
    15  15  15  15  16  15  | 91  
    15  15  15  15  15  15  | 90  
    13  16  16  15  17  16  | 93  
    14  14  14  14  15  15  | 86  
    17  17  16  18  18  16  | 102  
    16  16  16  16  17  16  | 97  
    14  16  15  15  16  15  | 91  
    15  15  17  15  15  17  | 94  
    15  15  15  15  15  16  | 91  
    14  15  15  16  15  15  | 90  
    15  16  16  16  16  16  | 95  
    17  15  15  17  15  17  | 96  
    15  15  15  16  16  16  | 93  
    16  17  16  17  16  16  | 98  
    15  15  12  16  16  12  | 86  
    15  18  14  17  16  17  | 97  
    15  17  15  17  17  16  | 97  
    15  15  15  16  15  16  | 92  
    14  14  15  15  15  15  | 88  
    14  14  15  15  15  15  | 88  

     3047249 jobs complete.