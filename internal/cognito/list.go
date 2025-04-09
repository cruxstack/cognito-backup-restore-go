package cognito

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserPoolDescription = types.UserPoolDescriptionType

var maxResults int32 = 10

func ListUserpools() ([]UserPoolDescription, error) {
	client, err := CreateClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create aws client for cognito idp: %w", err)
	}

	var token *string
	var pools []UserPoolDescription

	for {
		params := &cognitoidentityprovider.ListUserPoolsInput{
			MaxResults: &maxResults,
			NextToken:  token,
		}
		results, err := client.ListUserPools(context.Background(), params)
		if err != nil {
			return nil, fmt.Errorf("could not list userpools: %w", err)
		}

		pools = append(pools, results.UserPools...)

		if results.NextToken == nil {
			break
		}
		token = results.NextToken
	}

	return pools, nil
}
