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
## AWS Integration (Optional)
ContainScan can run **without AWS**, logging data locally in `containscan.log`.  
If AWS credentials are available, logs are **automatically stored in DynamoDB**.

### **Running Without AWS (Local Mode)**
By default, if no AWS credentials are set, logs are stored locally:
```bash
go run main.go
   cd containscan

