//This Package returns the list of users in a group. This is very helpful to find who are part of a group in AWS Group. 
//The benefit of this library is they can use it to find users in a group or script it like a security IDS (Intrusion Detection System) which can be used to detect any changes to groups. 
package group2users

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"log"
	"github.com/aws/aws-sdk-go/aws"

)

func ListUsersForGroup(iamService *iam.IAM,groupName string)  ([]string) {
	lookupGroup := groupName
		var users []string
		resultUsers, err := iamService.ListUsers(&iam.ListUsersInput{})
		for _, value := range resultUsers.Users {
			groupsOutput, _ := iamService.ListGroupsForUser(&iam.ListGroupsForUserInput{UserName: value.UserName})
			for _, val := range groupsOutput.Groups {
				if aws.StringValue(val.GroupName) == lookupGroup {
					users = append(users, aws.StringValue(value.UserName))
				}
			}
		}
		if err != nil {
			log.Fatal("error", err.Error())
		}
		return  users
	
}

func GetIAMService(session *session.Session) (*iam.IAM) {
	return iam.New(session)
}
