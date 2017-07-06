## Group2Users

List users in a AWS  group
=========================
  
Example Usage :

`package main

import "github.com/aws/aws-sdk-go/aws/session"



import (
	"github.com/prasanna-in/Group2Users"
	"fmt"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))
	fmt.Println(group2users.ListUsersForGroup(group2users.GetIAMService(sess),"Security"))

}
`




