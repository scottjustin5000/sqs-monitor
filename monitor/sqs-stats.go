package monitor

import (
  "fmt"
  "strings"
  "sync"
  "strconv"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/sqs"
)

var sqssvc *sqs.SQS

func getSvc()(*sqs.SQS, error){
  if(sqssvc != nil) {
     return sqssvc, nil
  }
  sqssvc, err := NewSQSClient("","","")
  if err != nil {
    return nil, err
  } 
  return sqssvc, nil
}

type QueueStatus struct {
    Name    string 
    Attributes  map[string]int
}

func ListQueues() map[string]string {
  svc, _ := getSvc()
  params := &sqs.ListQueuesInput{}
  sqs_resp, err := svc.ListQueues(params)
  if err != nil {
    fmt.Println(err.Error())
    return nil
  }
  qs := make(map[string]string)
  for _, url := range sqs_resp.QueueUrls {
    var u = *url
    index := strings.LastIndex(u, "/")
    if index > -1 {
      substring := u[index+1:len(u)]
      qs[substring] = u
    }
  }
  return qs
}

func GetQueueStatus(qs []string) []QueueStatus  {
  numQueues := len(qs)
  var wg sync.WaitGroup
  wg.Add(numQueues)
  var results []QueueStatus
  for _, q := range qs {
    go func(queue string) {
      var qs QueueStatus
      qs.Name = queue 
      defer wg.Done()
      qs.Attributes = getAttributes(queue)
      results = append(results, qs)
    }(q)
  }
  wg.Wait()

  return results

}

func getQueueUrl(name string) (string, error) {
  params := &sqs.GetQueueUrlInput{
    QueueName: aws.String(name), 
  }
  svc, _ := getSvc()
  resp, err := svc.GetQueueUrl(params)
  if err != nil {
    return "", err
  }
  return aws.StringValue(resp.QueueUrl), nil
}

func getAttributes(name string) map[string]int {
  x := make(map[string]int)
  url, err := getQueueUrl(name)
  if err != nil {
    fmt.Println("failed to get queue url,", err)
    return nil
  }
  params := &sqs.GetQueueAttributesInput{
    QueueUrl: aws.String(url),
    AttributeNames: []*string{
      aws.String("ApproximateNumberOfMessages"),
      aws.String("ApproximateNumberOfMessagesDelayed"),
      aws.String("ApproximateNumberOfMessagesNotVisible"),
    },
  }
  svc, _ := getSvc()
  resp, _ := svc.GetQueueAttributes(params)
  for attrib, _ := range resp.Attributes {
    prop := resp.Attributes[attrib]
    i, _ := strconv.Atoi(*prop)
    x[attrib] = i
    fmt.Println(attrib, i)
  }
  return x
}