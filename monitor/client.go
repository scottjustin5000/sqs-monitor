package monitor

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/sqs"
)

func NewSQSClient(key string, secret string, region string) (*sqs.SQS, error) {

  var svc *sqs.SQS
  var regionVal = ""
  if(region !=""){
    regionVal =  "us-west-2"
  }

  if key !="" && secret !="" {
    awsConfig := &aws.Config {
      Credentials: credentials.NewStaticCredentials(key, secret, ""),
      Region: aws.String(regionVal),
    }
    sess, err := session.NewSession()
    if err != nil {
      fmt.Println("failed to create session,", err)
      return nil, err
    }
    svc = sqs.New(sess, awsConfig)
  }else {
    svc = sqs.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})
  }

  return svc, nil
}