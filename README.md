Simple reverse proxy for AWS sigv4 termination.

### Usage:
```
go mod download
go build
./go-sigv4-proxy
```

Or with Docker:
```
docker-compose build
docker-compose up
```

### Configuration
Configure using environment variables or .env file:
|Variable|Description|
|---|---|
|GOSIGV4PROXY_PORT|Proxy will listen on this port|
|GOSIGV4PROXY_SERVICE|AWS service name, for example `execute-api`|
|GOSIGV4PROXY_ENDPOINT|Target URL|
|GOSIGV4PROXY_REGION|AWS region of the target|
|AWS_ACCESS_KEY_ID|AWS access key ID|
|AWS_SECRET_ACCESS_KEY|AWS access key|
|AWS_SESSION_TOKEN|AWS session token|
