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

| Flag        | Description                      | Default       |
| ----------- | -------------------------------- | ------------- |
| `--pool-id` | The Cognito user pool ID         | **required**  |
| `--out`     | Output file path for the backup  | `backup.json` |

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
