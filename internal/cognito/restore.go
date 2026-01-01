package cognito

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func RestoreUsers(poolId, inPath string) error {
	if poolId == "" {
		return fmt.Errorf("pool id is required")
	}
	if inPath == "" {
		return fmt.Errorf("input path is required")
	}

	client, err := CreateClient()
	if err != nil {
		return fmt.Errorf("failed to create cognito client: %w", err)
	}

	data, err := os.ReadFile(inPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var users []types.UserType
	if err := json.Unmarshal(data, &users); err != nil {
		return fmt.Errorf("failed to unmarshal users: %w", err)
	}

	slog.Info("starting restore", "pool_id", poolId, "user_count", len(users))

	var failedUsers []string
	for _, u := range users {
		params := &cognitoidentityprovider.AdminCreateUserInput{
			UserPoolId:     &poolId,
			Username:       u.Username,
			UserAttributes: u.Attributes,
		}
		_, err := client.AdminCreateUser(context.Background(), params)
		if err != nil {
			slog.Warn("failed to create user", "username", *u.Username, "error", err)
			failedUsers = append(failedUsers, *u.Username)
		}
	}

	if len(failedUsers) > 0 {
		return fmt.Errorf("failed to restore %d user(s): %v", len(failedUsers), failedUsers)
	}

	slog.Info("restore complete", "pool_id", poolId, "user_count", len(users))

	return nil
}
