## Group2Users

List users in a AWS  group
=========================
  
Example Usage :

package main


import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/prasanna-in/Group2Users"
	"fmt"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))
	//We first create aws session. We pass aws session object and groupname to get the users 
	// Output is string groupname and list users in the group
	
	fmt.Println(group2users.ListUsersForGroup(group2users.GetIAMService(sess),"GroupName"))
}




