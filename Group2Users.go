//This Package returns the list of users for a given group
//The Package takes a AWS session and groupname as input
//The Package returns the groupname and a list of users
package group2users

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"log"
	"github.com/aws/aws-sdk-go/aws"

)

func ListUsersForGroup(iamService *iam.IAM,groupName string)  (string, []string) {
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
		return lookupGroup, users
	
}

func GetIAMService(session *session.Session) (*iam.IAM) {
	return iam.New(session)
}
