package cognito

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// hexChars are the possible first characters of a UUID sub attribute
var hexChars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

// parseShard parses a shard string like "1/3" into index and total
func parseShard(shard string) (index, total int, err error) {
	parts := strings.Split(shard, "/")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("failed to parse shard: expected format 'index/total' (e.g., '1/3'), got '%s'", shard)
	}

	index, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse shard index: %w", err)
	}

	total, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse shard total: %w", err)
	}

	if total < 1 {
		return 0, 0, fmt.Errorf("failed to parse shard: total must be at least 1, got %d", total)
	}

	if index < 1 || index > total {
		return 0, 0, fmt.Errorf("failed to parse shard: index must be between 1 and %d, got %d", total, index)
	}

	return index, total, nil
}

// getHexPrefixesForShard returns the hex prefixes that belong to a given shard
// For example, with 3 shards:
//   - shard 1/3: ["0", "1", "2", "3", "4", "5"]
//   - shard 2/3: ["6", "7", "8", "9", "a"]
//   - shard 3/3: ["b", "c", "d", "e", "f"]
func getHexPrefixesForShard(shardIndex, totalShards int) []string {
	var prefixes []string
	for i, hex := range hexChars {
		// Assign hex char to shard using modulo
		if (i % totalShards) == (shardIndex - 1) {
			prefixes = append(prefixes, hex)
		}
	}
	return prefixes
}

// shardOutputPath modifies the output path to include the shard index
// e.g., "backup.json" with shard 2 becomes "backup-2.json"
func shardOutputPath(outPath string, shardIndex int) string {
	ext := filepath.Ext(outPath)
	base := strings.TrimSuffix(outPath, ext)
	return fmt.Sprintf("%s-%d%s", base, shardIndex, ext)
}

func BackupUsers(poolId, outPath, shard string) error {
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

	var allUsers []types.UserType

	// Determine if we're sharding
	if shard != "" {
		shardIndex, totalShards, err := parseShard(shard)
		if err != nil {
			return err
		}

		// Modify output path to include shard index
		outPath = shardOutputPath(outPath, shardIndex)

		prefixes := getHexPrefixesForShard(shardIndex, totalShards)
		slog.Info("starting backup", "pool_id", poolId, "shard", shard, "prefixes", prefixes)

		// Fetch users for each hex prefix sequentially
		for _, prefix := range prefixes {
			filter := fmt.Sprintf(`sub ^= "%s"`, prefix)
			slog.Debug("fetching users", "filter", filter)

			var token *string
			for {
				results, err := client.ListUsers(context.Background(), &cognitoidentityprovider.ListUsersInput{
					UserPoolId:      &poolId,
					Filter:          aws.String(filter),
					PaginationToken: token,
				})
				if err != nil {
					return fmt.Errorf("failed to list users with filter '%s': %w", filter, err)
				}
				allUsers = append(allUsers, results.Users...)
				if results.PaginationToken == nil {
					break
				}
				token = results.PaginationToken
			}
		}
	} else {
		// No sharding - fetch all users
		slog.Info("starting backup", "pool_id", poolId)

		var token *string
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
