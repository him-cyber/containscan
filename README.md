# ContainScan
ContainScan is a Go-based container security tool designed to detect vulnerabilities in Docker environments. It supports structured logging with optional AWS integration.

## Features
- Automated vulnerability scanning for containerized applications
- Structured logging in JSON format
- Local logging to a file (`containscan.log`)
- Optional AWS DynamoDB integration for cloud-based log storage

## OS Compatibility
ContainScan has been tested and works on:
- Linux (Ubuntu, Debian, Arch, Pop!_OS)
- Windows
- macOS (Testing planned)

By default, ContainScan runs in **local mode** and logs to `containscan.log`. If AWS credentials are configured, logs will also be stored in DynamoDB.

## Installation
1. **Clone the repository:**
   ```bash
   git clone https://github.com/him-cyber/containscan.git
   cd containscan

