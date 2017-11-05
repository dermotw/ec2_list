package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
  // Load the session
  sess := session.Must(session.NewSessionWithOptions( session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))

  // create an EC2 Client
  ec2Svc := ec2.New( sess )
  result, err := ec2Svc.DescribeInstances(nil)
  if err != nil {
    fmt.Println( "Error: ", err )
  } else {
    var l = len(result.Reservations)
    for i := 0; i < l; i++ {
      var inst = result.Reservations[i]
      fmt.Println(aws.StringValue(inst.Instances[0].Tags[0].Value))
    }
  }
}
