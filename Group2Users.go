package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"fmt"
)

func main() {
	lookupGroup, users := getUsers(os.Args[1])
	fmt.Println("Users in Group "+lookupGroup+" : ", users)
}

func getUsers(GroupName string) (string, []string) {
	lookupGroup := GroupName
	var users []string
	iamService := getIAMService()
	
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
func getIAMService() *iam.IAM {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	iamService := iam.New(sess)
	return iamService
}
