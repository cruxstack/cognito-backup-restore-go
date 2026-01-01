package cognito

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserPoolDescription = types.UserPoolDescriptionType

var maxResults int32 = 10

func ListUserpools() ([]UserPoolDescription, error) {
	client, err := CreateClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create cognito client: %w", err)
	}

	var token *string
	var pools []UserPoolDescription

	slog.Info("listing user pools")

	for {
		params := &cognitoidentityprovider.ListUserPoolsInput{
			MaxResults: &maxResults,
			NextToken:  token,
		}
		results, err := client.ListUserPools(context.Background(), params)
		if err != nil {
			return nil, fmt.Errorf("failed to list user pools: %w", err)
		}

		pools = append(pools, results.UserPools...)

		if results.NextToken == nil {
			break
		}
		token = results.NextToken
	}

	slog.Info("list complete", "pool_count", len(pools))

	return pools, nil
}
