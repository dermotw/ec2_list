package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
  // Print the first line
  fmt.Println( "EC2go" )
  fmt.Println( "---" )

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
      var instId = aws.StringValue(inst.Instances[0].InstanceId)
      var instName = aws.StringValue(inst.Instances[0].Tags[0].Value)
      var instState = aws.StringValue(inst.Instances[0].State.Name)

      var color string

      fmt.Printf("%s", instName)

      if instState == "running" {
        color = "green"
        fmt.Printf( " | color=%s\n", color )
        fmt.Printf( "-- Stop Instance | bash=/home/dermot/bin/stop_aws param1=%s terminal=false\n", instId )
      } else {
        color = "red"
        fmt.Printf( " | color=%s\n", color )
        fmt.Printf( "-- Start Instance | bash=/home/dermot/bin/start_aws param1=%s terminal=false\n", instId )
      }
    }
  }
}
