# Cognito Backup Restore

## What

A CLI tool for backing up and restoring AWS Cognito user pool data.

## Why

AWS Cognito doesn't provide built-in backup/restore functionality. This tool
lets you:

- **Back up user data** from any Cognito user pool to a JSON file
- **Restore users** to a pool from a backup file
- **List user pools** in your AWS account for discovery

## Usage

### Installation

Download a pre-built binary from the
[releases page](https://github.com/cruxstack/cognito-backup-restore-go/releases),
or build from source:

```bash
go build -o cbr .
```

### Prerequisites

- Valid AWS credentials configured (via environment, shared credentials, or IAM role)
- IAM permissions: `cognito-idp:ListUserPools`, `cognito-idp:ListUsers`, `cognito-idp:AdminCreateUser`

### Commands

**List user pools**

```bash
cbr list
```

**Back up users**

```bash
cbr backup --pool-id <POOL-ID> --out <FILE-PATH>
```

| Flag        | Description                                              | Default       |
| ----------- | -------------------------------------------------------- | ------------- |
| `--pool-id` | The Cognito user pool ID                                 | **required**  |
| `--out`     | Output file path for the backup                          | `backup.json` |
| `--shard`   | Deterministic shard index/total (e.g., `1/3`, `2/3`)     | *(optional)*  |

**Sharded backups**

For large user pools, you can run backups in parallel across multiple CI runners
using the `--shard` flag.

```bash
# Run 3 parallel backup jobs
cbr backup --pool-id us-east-1_xxx --shard 1/3 --out backup.json  # outputs backup-1.json
cbr backup --pool-id us-east-1_xxx --shard 2/3 --out backup.json  # outputs backup-2.json
cbr backup --pool-id us-east-1_xxx --shard 3/3 --out backup.json  # outputs backup-3.json
```

Sharding uses the user's `sub` attribute (UUID) prefix to deterministically
partition users. Each shard only fetches its portion of users from AWS, enabling
true parallel execution without coordination between runners. The output
filename is automatically suffixed with the shard index.

**Restore users**

```bash
cbr restore --pool-id <POOL-ID> --in <FILE-PATH>
```

| Flag        | Description                      | Default       |
| ----------- | -------------------------------- | ------------- |
| `--pool-id` | The Cognito user pool ID         | **required**  |
| `--in`      | Input file path for the backup   | `backup.json` |

---

# Development

## Project Structure

```
├── cmd/                  # CLI command definitions
│   ├── backup.go
│   ├── list.go
│   ├── restore.go
│   └── root.go
├── internal/cognito/     # AWS Cognito client and operations
│   ├── backup.go
│   ├── client.go
│   ├── list.go
│   └── restore.go
├── build/                # Build scripts
└── main.go               # Entrypoint
```

## Building

```bash
# Local build
go build -o cbr .

# Cross-platform builds (for releases)
./build/build.sh v1.0.0
```

## Testing

```bash
go test ./...
```

## Linting

```bash
go vet ./...
gofmt -l .
```
