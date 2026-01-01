package cognito

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func RestoreUsers(poolId, inPath string) error {
	if poolId == "" {
		return fmt.Errorf("userpool id is required")
	}
	if inPath == "" {
		return fmt.Errorf("input path is required")
	}

	client, err := CreateClient()
	if err != nil {
		return fmt.Errorf("failed to create aws client for cognito idp: %w", err)
	}

	data, err := os.ReadFile(inPath)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	var users []types.UserType
	if err := json.Unmarshal(data, &users); err != nil {
		return fmt.Errorf("could not unmarshal users: %w", err)
	}

	var failedUsers []string
	for _, u := range users {
		params := &cognitoidentityprovider.AdminCreateUserInput{
			UserPoolId:     &poolId,
			Username:       u.Username,
			UserAttributes: u.Attributes,
		}
		_, err := client.AdminCreateUser(context.Background(), params)
		if err != nil {
			fmt.Printf("could not create user %s: %v\n", *u.Username, err)
			failedUsers = append(failedUsers, *u.Username)
		}
	}

	if len(failedUsers) > 0 {
		return fmt.Errorf("failed to restore %d user(s): %v", len(failedUsers), failedUsers)
	}

	return nil
}
