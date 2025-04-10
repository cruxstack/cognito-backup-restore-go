package cognito

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func BackupUsers(poolId, outPath string) error {
	if poolId == "" {
		return fmt.Errorf("userpool id is required")
	}
	if outPath == "" {
		return fmt.Errorf("output path is required")
	}

	client, err := CreateClient()
	if err != nil {
		return fmt.Errorf("could not load aws config: %w", err)
	}

	var (
		token    *string
		allUsers []types.UserType
	)

	for {
		results, err := client.ListUsers(context.Background(), &cognitoidentityprovider.ListUsersInput{
			UserPoolId:      &poolId,
			PaginationToken: token,
		})
		if err != nil {
			return fmt.Errorf("could not list users: %w", err)
		}
		allUsers = append(allUsers, results.Users...)
		if results.PaginationToken == nil {
			break
		}
		token = results.PaginationToken
	}

	data, err := json.MarshalIndent(allUsers, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal users: %w", err)
	}

	if err := os.WriteFile(outPath, data, 0644); err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	return nil
}
