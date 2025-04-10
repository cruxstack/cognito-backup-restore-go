# cognito-backup-restore-go

A CLI tool for managing AWS Cognito user pools and user data. It allows you to
list Cognito user pools, back up users from a specific pool, and restore users
to a pool from a backup file.

## Prerequisites

- [Go 1.24+](https://go.dev/dl/) (or later) installed
- Valid AWS credentials configured
- Proper IAM permissions to list user pools, list pool users, and create users
  in Cognito

## Building

1. Navigate to the project root.
2. Run:
   ```bash
   go build -o cbr .
   ```
   This compiles the project and produces a `cbr` (or `cbr.exe` on Windows) binary.

## Usage

Once you have the binary, you can run the following commands:

- **List user pools**
  ```bash
  ./cbr list
  ```
  Lists up to 10 user pools at a time (it will iterate through all user pools until all are shown).

- **Back up users from a user pool**
  ```bash
  ./cbr backup --pool-id <POOL-ID> --out <FILE-PATH>
  ```
  - `--pool-id`: The user pool ID. Required.
  - `--out`: The path to the file where you want the backup written. Default: `backup.json`.

- **Restore users to a user pool**
  ```bash
  ./cbr restore --pool-id <POOL-ID> --in <FILE-PATH>
  ```
  - `--pool-id`: The user pool ID. Required.
  - `--in`: Path to the backup file. Default: `backup.json`.

