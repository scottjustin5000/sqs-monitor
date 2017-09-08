# SQS-Monitor

simple wrapper for monitoring SQS.

### Usage


* List Queues

```go
import (
  "github.com/scottjustin5000/sqs-monitor/monitor"
)


func main() {
  queues := monitor.ListQueues()

  for k, v := range queues {
   fmt.Println(k)
   fmt.Println(v)
  }
}

```

* Queue Statuses

```go
import (
  "github.com/scottjustin5000/sqs-monitor/monitor"
)


func main() {
  qs := []string{ "queue1","queue2", }
  statuses := monitor.GetQueueStatus(qs)

   for _, k := range statuses {
     fmt.Println("name:", k.Name)
     for kv, vv := range k.Attributes {
       fmt.Println("key:", kv, "val:", vv )
     }
   }
}

```

Assumes ~/.aws profile is present
