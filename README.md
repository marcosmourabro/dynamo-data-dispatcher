# AWS Lambda, API Gateway, and DynamoDB Integration

This project demonstrates a straightforward integration between AWS Lambda, API Gateway, and DynamoDB. The Lambda function is designed to receive a request via API Gateway and then insert the relevant data into DynamoDB.

## Architecture

1. **API Gateway**: Acts as an entry point, receiving HTTP requests and forwarding them to the Lambda function.
2. **AWS Lambda**: Processes the incoming request, extracts the necessary data, and performs operations on DynamoDB.
3. **DynamoDB**: Persistent storage where request data is inserted.

## Prerequisites

Before starting or deploying this project, ensure you have the following installed:

- AWS CLI
- Go Language (version 1.18)
- Serverless Framework

## Installation and Setup

1. **Configure AWS CLI**:
    ```bash
    aws configure
    ```

2. **Clone the Repository**:
    ```bash
    git clone git@github.com:marcosmourabro/dynamo-data-dispatcher.git
    cd dynamo-data-dispatcher
    ```

3. **Install Dependencies (if applicable)**:
    ```bash
    go mod tidy
    ```

4. **Deploy via Serverless Framework**:
    ```bash
    serverless deploy
    ```
