# HRDWSTATUS

A lightweight hardware monitoring utility written in Go that tracks CPU temperature and reports its status.

## Overview

HRDWSTATUS reads CPU temperature data from the Linux hwmon system interface and categorizes temperatures into three levels:
- **LOW**: Temperature ≤ 25°C
- **OK**: Temperature between 26°C and 59°C
- **HIGH**: Temperature ≥ 60°C

The application is designed to be simple and efficient, with minimal resource usage.

## Features

- Real-time CPU temperature monitoring
- Temperature status categorization
- Structured logging using Go's slog package
- Easy to run as a daemon service

## Installation

### Prerequisites
- Go 1.20 or later
- Linux-based operating system with access to `/sys/class/hwmon/`

### Building from source
```bash
git clone https://github.com/yourusername/hrdwstatus.git
cd hrdwstatus
go build
```

## Usage

### Running directly
```bash
./hrdwstatus
```

### Running as a service
The project includes a systemd service file (`hrdwstsd.service`) that can be used to run HRDWSTATUS as a daemon:

1. Copy the service file to the systemd directory:
   ```bash
   sudo cp hrdwstsd.service /etc/systemd/system/
   ```

2. Reload systemd:
   ```bash
   sudo systemctl daemon-reload
   ```

3. Enable and start the service:
   ```bash
   sudo systemctl enable hrdwstsd.service
   sudo systemctl start hrdwstsd.service
   ```

## Container Usage

HRDWSTATUS can be deployed in an Incus container for testing or production use:

1. Create an Incus container
2. Install the application inside the container
3. Set up the daemon using the provided hrdwstsd.service file

## Project Structure

```
HRDWSTATUS/
├── cmd/
│   └── main.go       # Main application entry point
├── deamon/
│   └── hrdwstsd.service # Systemd service file for running as daemon
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
└── readme.md         # This file
```

## Known Limitations

- Currently only detects AMD CPUs with the "k10temp" sensor
- Limited to Linux systems with access to `/sys/class/hwmon/`

## Future Improvements

- Add support for more CPU types
- Include GPU temperature monitoring
- Implement real-time alerts for prolonged high temperatures
- Add configuration file for customizing temperature thresholds

## License

[Insert your chosen license here]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.