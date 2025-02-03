# ContainScan
**A Go-based container security tool for detecting vulnerabilities in Docker environments.**  

## Features
- Automated vulnerability scanning for containerized applications  
- Structured logging in JSON format with AWS DynamoDB integration  
- Real-time security alerts (Planned for future versions)  

## Installation
1. **Clone the repository:**  
   ```bash
   git clone https://github.com/himaneesh/containscan.git

## OS Compatibility
ContainScan has been tested and works on:
- ✅ **Linux (Ubuntu, Debian, Arch, Pop!_OS)**
- ✅ **Windows (Fully compatible)**
- ✅ **macOS (Planned Testing)**

By default, ContainScan runs in **local mode** (logs to `containscan.log`). If AWS credentials are present, logs are also stored in DynamoDB.

## AWS Integration (Optional)
ContainScan can run **without AWS**, logging data locally in `containscan.log`.  
If AWS credentials are available, logs are **automatically stored in DynamoDB**.

### **Running Without AWS (Local Mode)**
By default, if no AWS credentials are set, logs are stored locally:
```bash
go run main.go
   cd containscan

