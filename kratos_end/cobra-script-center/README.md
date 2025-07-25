# Cobra Script Center

A powerful command-line script management and execution center built with Go and Cobra.

## Features

- 📝 **Script Management**: Create, edit, delete, and organize scripts
- ⚡ **Script Execution**: Execute scripts with parameters and monitoring
- ⏰ **Task Scheduling**: Schedule scripts with cron expressions
- 👥 **User Management**: User authentication and role-based permissions
- 📊 **Monitoring**: Execution logs, status tracking, and notifications
- 🔍 **Search & Filter**: Find scripts by name, tags, or description
- 🔒 **Security**: Secure script execution with sandboxing and permissions

## Quick Start

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd cobra-script-center

# Install dependencies
make deps

# Build the application
make build

# Install globally (optional)
make install
```

### Configuration

Copy the example configuration file and customize it:

```bash
cp .script-center.yaml.example .script-center.yaml
```

Edit `.script-center.yaml` to configure database, security, and logging settings.

### Initialize Database

```bash
./bin/script-center migrate
```

### Create First User

```bash
./bin/script-center user create --username admin --role admin
```

## Usage

### Script Management

```bash
# Create a new script
./bin/script-center script create --name "backup" --language bash

# List all scripts
./bin/script-center script list

# Edit a script
./bin/script-center script edit backup

# Delete a script
./bin/script-center script delete backup
```

### Script Execution

```bash
# Execute a script
./bin/script-center script run backup

# Execute with parameters
./bin/script-center script run backup --param key=value

# Schedule a script
./bin/script-center script schedule backup --cron "0 2 * * *"
```

### User Management

```bash
# Create a user
./bin/script-center user create --username john --role user

# Login
./bin/script-center user login --username john

# List users (admin only)
./bin/script-center user list
```

### Daemon Mode

```bash
# Start the daemon for scheduled tasks
./bin/script-center daemon start

# Stop the daemon
./bin/script-center daemon stop

# Check daemon status
./bin/script-center daemon status
```

## Development

### Prerequisites

- Go 1.21 or higher
- SQLite3
- Make

### Setup Development Environment

```bash
make dev-setup
```

### Running Tests

```bash
make test
```

### Code Formatting

```bash
make fmt
```

### Building

```bash
make build
```

## Configuration

The application uses a YAML configuration file. Here's an example:

```yaml
database:
  driver: sqlite3
  dsn: ./script-center.db

server:
  host: localhost
  port: 8080

security:
  jwt_secret: "your-secret-key"
  password_salt: "your-salt"
  max_executions: 10

logging:
  level: info
  format: json
  file: ""
```

## Architecture

```
├── cmd/                 # CLI commands
├── internal/
│   ├── app/            # Application initialization
│   ├── config/         # Configuration management
│   ├── database/       # Database connection and migrations
│   ├── logger/         # Logging configuration
│   ├── models/         # Data models
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic layer
│   ├── executor/       # Script execution engine
│   └── scheduler/      # Task scheduling
├── migrations/         # Database migrations
└── docs/              # Documentation
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run `make test` and `make lint`
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.