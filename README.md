# Log Processing with Golang, Logstash, and Elasticsearch

## Overview
This project sets up a logging pipeline using Logstash, Elasticsearch, and a Golang application. The Golang service sends logs to Logstash, which processes and forwards them to Elasticsearch for indexing and storage.

## Technologies Used
- **Golang**: The main application
- **Logstash**: Parses and forwards logs
- **Elasticsearch**: Stores and indexes logs
- **Docker**: Containerization and orchestration

## Project Structure
```
.
├── docker-compose.yml  # Docker services definition
├── logstash/
│   └── pipeline/       # Logstash pipeline configuration
├── main.go             # Golang application source
├── go.mod              # Go module file
└── Dockerfile          # Golang application container setup
```

## Logstash Configuration
Located in `logstash/pipeline/logstash.conf`, Logstash listens on port `5044` for logs and processes them using Grok patterns:

```json
input {
  beats {
    port => 5044
  }
}

filter {
  grok {
    match => { "message" => "%{TIMESTAMP_ISO8601:timestamp} %{LOGLEVEL:loglevel} %{GREEDYDATA:message}" }
  }
  date {
    match => [ "timestamp", "ISO8601" ]
  }
}

output {
  elasticsearch {
    hosts => ["http://elasticsearch:9200"]
    index => "logs-%{+YYYY.MM.dd}"
  }
  stdout {
    codec => rubydebug
  }
}
```

## Docker Setup
The project uses Docker Compose to orchestrate the services:

- **Golang App**
  - Built from `Dockerfile`
  - Exposes port `9090`
  - Depends on `logstash`
- **Elasticsearch**
  - Runs version `8.7.1`
  - Stores indexed logs
  - Exposes port `9200`
- **Logstash**
  - Runs version `8.7.1`
  - Listens on port `5044`
  - Processes and forwards logs to Elasticsearch


## Contributing
Feel free to fork the repository and submit pull requests for improvements!


