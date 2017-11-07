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
      var netIfs = inst.Instances[0].NetworkInterfaces

      var color string

      fmt.Printf("%s", instName)

      if instState == "running" {
        color = "green"
        fmt.Printf( " | color=%s\n", color )
        fmt.Printf( "-- Stop Instance | bash=ec2_control param1=stop param2=%s terminal=false\n", instId )
      } else {
        color = "red"
        fmt.Printf( " | color=%s\n", color )
        fmt.Printf( "-- Start Instance | bash=ec2_control param1=start param2=%s terminal=false\n", instId )
      }

      fmt.Printf( "-- State: %s\n", instState )

      // iterate over network interfaces
      for j := 0; j < len( netIfs ); j++ {
        fmt.Printf( "-- IP: %s", aws.StringValue(netIfs[j].PrivateIpAddress) )
        Assoc := netIfs[j].Association
        if Assoc != nil {
          var PublicIp = aws.StringValue(Assoc.PublicIp)
          fmt.Printf( ", %s", PublicIp)
        } else {
          // do nothing
        }
        fmt.Printf("\n")
      }
    }
  }
}
