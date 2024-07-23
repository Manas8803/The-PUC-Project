# PUC Project Backend

This repository contains the backend code for the PUC (Pollution Under Control) Project, implemented using Golang and AWS services.

## Architecture Overview

The backend is built on a microservices architecture deployed on AWS, utilizing various AWS services for scalability, performance, and cost-effectiveness.

Key components:

- **Golang**: Main programming language for backend logic 
- **AWS Lambda**: Serverless compute for running backend functions
- **AWS DynamoDB**: NoSQL database for storing PUC certificate data
- **AWS Textract**: OCR service for extracting vehicle registration numbers from images
- **AWS Scheduler**: For scheduling periodic tasks
- **Twilio**: Integrated for sending notifications to users

## Microservices

Our backend consists of several microservices:

   1. Auth Service
   2. OCR Service
   3. VRC Service
   4. Registration Renewal Reminder Service
   5. Fetch Vehicle Service
   6. Registration Expiration Job Service

## Setup and Configuration

### Prerequisites

- Go 1.x or later
- AWS CLI
- AWS CDK
- AWS account with appropriate permissions
- Make (for using the provided Makefile)

### Environment Setup

1. Clone this repository:
   ```
   git clone [repository-url]
   cd [repository-name]
   ```

2. Configure AWS credentials:
   ```
   aws configure
   ```
   Enter your AWS Access Key ID and Secret Access Key when prompted.

3. Install dependencies:
   ```
   go mod tidy
   ```

4. Set up AWS CDK (if not already done):
   ```
   npm install -g aws-cdk
   cdk bootstrap
   ```

## Build and Deployment

We use a Makefile to simplify the build and deployment process. Here are the main commands:

### Build

To build all services:

```
make build
```

This command compiles each service for Linux AMD64 architecture.

### Deploy

To deploy using AWS CDK:

```
make deploy
```

For a faster deployment using CDK's hotswap feature:

```
make deploy-swap
```

### Clean

To remove all compiled binaries:

```
make clean
```

### Build and Deploy

To clean, build, and deploy in one command:

```
make all
```

Or with hotswap deployment:

```
make all-swap
```

## Key Features

1. **OCR Processing**: Utilizes AWS Lambda and Textract to extract vehicle registration numbers from images.
2. **Periodic Scanning**: AWS Scheduler triggers a Lambda function to scan DynamoDB for expired PUC certificates.
3. **Notification System**: Integrated with Twilio to send reminders for expiring PUC certificates.