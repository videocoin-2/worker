
# Blockchain Worker Service

The Blockchain Worker Service is a Go-based service designed to process distributed tasks, such as video/audio transcoding, and interact with a blockchain network. This service integrates with smart contracts, enabling secure task tracking, resource utilization, and result verification on the blockchain.

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Architecture](#architecture)
4. [Components](#components)
5. [Workflow](#workflow)
6. [Installation](#installation)
7. [Usage](#usage)
8. [Configuration](#configuration)
9. [Development](#development)
10. [Contributing](#contributing)
11. [License](#license)

---

## Overview

The Blockchain Worker Service serves as a task executor in a decentralized network. It is designed to:
- Process computationally intensive tasks such as media transcoding.
- Report task progress and results back to the blockchain.
- Ensure transparent and verifiable task execution.

---

## Features

- **Task Management**: Handle distributed tasks efficiently using a queue-based system.
- **Media Processing**: Perform video and audio transcoding using FFmpeg and generate HLS streams.
- **Blockchain Integration**: Communicate with blockchain nodes for smart contract interactions.
- **System Health Monitoring**: Regular health checks to monitor service and node availability.
- **Capacity Management**: Dynamically check and manage available system resources for optimal task execution.
- **Scalability**: Designed to scale across multiple workers in a decentralized environment.
- **System Monitoring**: Fetch and report system resource utilization.

---

## Architecture

The service is modular and includes the following key components:
- **Task Management**: Handles task creation, assignment, and status updates.
- **Transcoding Engine**: Uses FFmpeg for video/audio processing.
- **Blockchain Client**: Interfaces with blockchain nodes and smart contracts.
- **Utility Layer**: Provides helper functions and error handling mechanisms.
- **Health Monitoring**: Ensures system and blockchain node availability.
- **Resource Management**: Manages system capacity and dynamic task allocation.

---

## Components

### File Descriptions

- **`error.go`**: Defines custom error types and error-handling utilities.
- **`ffmpeg.go`**: Implements FFmpeg transcoding operations.
- **`hls.go`**: Handles HLS (HTTP Live Streaming) generation.
- **`task.go`**: Manages task lifecycle, including task creation, updates, and completion tracking.
- **`transcoder.go`**: Coordinates and executes transcoding tasks.
- **`util.go`**: Provides helper functions for common operations.
- **`sysinfo.go`**: Fetches system resource information (CPU, memory, disk).
- **`client.go`**: Manages interactions with blockchain nodes via RPC.
- **`contract.go`**: Interfaces with smart contracts to store and retrieve task data.
- **`main.go`**: The entry point for the service, initializing all components.
- **`caller.go`**: Facilitates blockchain method calls and event subscriptions.
- **`capacity.go`**: Manages system capacity by assessing available resources for task execution.
- **`health.go`**: Implements health-check functionality for the service and blockchain nodes.
- **`pinger.go`**: Sends periodic pings to ensure connectivity with the blockchain network.
- **`Dockerfile`**: Defines the containerization process for the service.
- **`Makefile`**: Provides build automation and task execution commands.

---

## Workflow

1. **Task Initialization**: Tasks are either pulled from the blockchain or received from an external source.
2. **Processing**:
   - If media processing is required, FFmpeg performs the transcoding.
   - HLS streams are generated if necessary.
3. **Health Monitoring**:
   - The service periodically checks its own health and blockchain node status.
   - Connectivity issues trigger alerts or retries.
4. **Capacity Management**:
   - Before processing, the service checks available resources.
   - Ensures tasks are executed within optimal conditions.
5. **Blockchain Interaction**:
   - Task progress and results are sent to the blockchain via smart contract methods.
   - Results are logged for verification.
6. **Completion**: The service marks tasks as completed and ready for retrieval.

---

## Installation

### Prerequisites

- [Go 1.19+](https://golang.org/doc/install)
- [FFmpeg](https://ffmpeg.org/download.html)
- Access to a blockchain node with RPC capabilities.
- Docker (for containerized deployments).

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/blockchain-worker-service.git
   cd blockchain-worker-service
   ```

   Install dependencies & compile the service:
    ```
   go mod tidy
   go build -o blockchain-worker-service
    ```

    Build the Docker image (optional):

    docker build -t blockchain-worker-service .

Usage

Run the service using the following command:

```
./blockchain-worker-service
```

Alternatively, run with Docker:

```
docker run --rm -v $(pwd)/config:/app/config blockchain-worker-service
```

Command-Line Options

- -config: Specify the configuration file path.
- -verbose: Enable verbose logging.

Configuration

The service requires a configuration file (e.g., config.json) to define essential parameters:

- Blockchain RPC URL: URL for the blockchain node.
- FFmpeg Path: Path to the FFmpeg executable.
- Task Queue Settings: Configuration for task queuing and processing.

Example configuration:
```
{
  "rpc_url": "http://localhost:8545",
  "ffmpeg_path": "/usr/bin/ffmpeg",
  "task_queue": {
    "max_retries": 3,
    "retry_interval": "10s"
  }
}
```



