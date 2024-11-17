#!/bin/bash

# Set the target gRPC server address
target_addr="localhost:8080"

# Define the number of concurrent requests and total requests
concurrency=50
total_requests=1000

# Define the gRPC method to test
method="log.LogService/CreateLog"

# Create a file with sample log requests
cat > log_requests.json <<EOF
[
  {
    "id": "1",
    "message": "This is a sample log message",
    "level": 1,
    "createdAt": "2023-04-01T12:00:00Z",
    "updatedAt": "2023-04-01T12:00:00Z"
  },
  {
    "id": "2",
    "message": "Another sample log message",
    "level": 2,
    "createdAt": "2023-04-01T12:00:01Z",
    "updatedAt": "2023-04-01T12:00:01Z"
  }
]
EOF

# Run the load test using grpcurl
grpcurl -d @log_requests.json -c $concurrency -n $total_requests $target_addr $method

# Output the results
echo "Load test results:"
grpcurl -d @log_requests.json -c $concurrency -n $total_requests $target_addr $method | wc -l