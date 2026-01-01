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

func BackupUsers(poolId, outPath string) error {
	if poolId == "" {
		return fmt.Errorf("pool id is required")
	}
	if outPath == "" {
		return fmt.Errorf("output path is required")
	}

	client, err := CreateClient()
	if err != nil {
		return fmt.Errorf("failed to create cognito client: %w", err)
	}

	var (
		token    *string
		allUsers []types.UserType
	)

	slog.Info("starting backup", "pool_id", poolId)

	for {
		results, err := client.ListUsers(context.Background(), &cognitoidentityprovider.ListUsersInput{
			UserPoolId:      &poolId,
			PaginationToken: token,
		})
		if err != nil {
			return fmt.Errorf("failed to list users: %w", err)
		}
		allUsers = append(allUsers, results.Users...)
		if results.PaginationToken == nil {
			break
		}
		token = results.PaginationToken
	}

	data, err := json.MarshalIndent(allUsers, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal users: %w", err)
	}

	if err := os.WriteFile(outPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	slog.Info("backup complete", "pool_id", poolId, "user_count", len(allUsers), "output", outPath)

	return nil
}
