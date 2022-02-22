package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

var cmd = []string{
	"createUser",
	"listAccessKey",
	"listUsers",
	"updateUserName",
	"deleteAccessKeyID",
}

func Operation(input string) string {

	for _, val := range cmd {
		if val == input {
			return val
		}
	}
	return "invalid"
}

// Usage:
// set the following environment variables in the terminal

// export AWS_ACCESS_KEY_ID=<<Your-api-access-key>>
// export AWS_SECRET_ACCESS_KEY=<<your-secret-access-key>>

func main() {
	//creating a session for asia-pacific(mumbai) region
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	},
	)

	// creating a new iam service client
	svc := iam.New(sess)

	command := os.Args[1]

	switch Operation(command) {
	case "createUser":
		_, err = svc.GetUser(&iam.GetUserInput{
			UserName: &os.Args[2],
		})
		// create a new user in case of  NoSuchEntity error
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == iam.ErrCodeNoSuchEntityException {
			result, err := svc.CreateUser(&iam.CreateUserInput{
				UserName: &os.Args[2],
			})

			if err != nil {
				fmt.Println("Create user error: ", err)
				return
			}

			fmt.Println("User Created: ", result)

		} else if err != nil {
			fmt.Println("Get User Error: ", err)

		} else {
			fmt.Println("User already present")

		}
	case "listAccessKey":
		result, err := svc.ListAccessKeys(&iam.ListAccessKeysInput{
			MaxItems: aws.Int64(5),
			UserName: &os.Args[2],
		})

		if err != nil {
			fmt.Println("List Access Keys Error: ", err)
			return
		}
		// print the list of access keys
		fmt.Println("Success:\n", result)
	case "listUsers":
		result, err := svc.ListUsers(&iam.ListUsersInput{
			MaxItems: aws.Int64(50),
		})

		if err != nil {
			fmt.Println("Error", err)
			return
		}

		for i, user := range result.Users {
			if user == nil {
				continue
			}
			fmt.Printf("%d user %s created %v\n", i, *user.UserName, user.CreateDate)
		}
	case "updateUserName":
		result, err := svc.UpdateUser(&iam.UpdateUserInput{
			UserName:    &os.Args[2],
			NewUserName: &os.Args[3],
		})

		if err != nil {
			fmt.Println("update name error", err)
		}
		fmt.Println("Success:", os.Args[2], "changed to", os.Args[3], "\n", result)
	case "deleteAccessKeyID":
		result, err := svc.DeleteAccessKey(&iam.DeleteAccessKeyInput{
			AccessKeyId: &os.Args[3],
			UserName:    &os.Args[2],
		})

		if err != nil {
			fmt.Println("Error", err)
			return
		}

		fmt.Println("Success", result)
	default:
		fmt.Println("Please enter a valid command")
	}

}
