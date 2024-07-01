# NonHypervisor : A Mock Containerization Engine

This project is a basic implementation of a mock containerization engine designed to help understand the fundamentals of containerization and how Docker works. It is intended as an educational tool and should be used with individual discretion.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Configuration](#configuration)
- [Setup](#setup)
- [Usage](#usage)
- [Example](#example)
- [Limitations](#limitations)
- [Acknowledgements](#acknowledgements)

## Overview

This project simulates the behavior of containerization by isolating processes using Linux namespaces, setting up root filesystems, environment variables, and networking. It serves as a parody of Docker and demonstrates how container environments can be created and managed.

## Features

- Namespace isolation (PID, UTS, mount, network)
- Root filesystem setup using images from a container registry
- Environment variable configuration
- Command execution within the container
- Networking setup with virtual Ethernet pairs and port forwarding

## Configuration

Configuration for the container setup is provided via a text file (`config.txt` by default). The configuration file supports the following directives:

- `FROM <image>`: Specifies the base image to use for the root filesystem.
- `ENV <key>=<value>`: Sets environment variables inside the container.
- `RUN <command>`: Runs a command during the container setup.
- `CMD <command>`: Specifies the command to run inside the container.
- `EXPOSE <port>`: Exposes a port from the container.

## Setup

### Prerequisites

- Go (Golang) installed on your system.
- Basic knowledge of Linux commands and networking.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/mock-containerization-engine.git
   cd mock-containerization-engine
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Running the Mock Container Engine

1. Create a configuration file (`config.txt`) with the desired container setup.
2. Run the mock container engine:
   ```bash
   go run main.go config.txt
   ```

### Configuration File Format

Here is an example `config.txt`:

```txt
FROM ubuntu:latest
ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
RUN apt-get update
RUN apt-get install -y curl
CMD ["/bin/bash"]
EXPOSE 8080
```

## Example

1. Create a configuration file named `config.txt`:
   ```txt
   FROM ubuntu:latest
   ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
   RUN apt-get update
   RUN apt-get install -y curl
   CMD ["/bin/bash"]
   EXPOSE 8080
   ```

2. Run the mock container engine:
   ```bash
   go run main.go config.txt
   ```

## Limitations

- This project is a simplified and educational tool, not suitable for production use.
- Lacks advanced features and robustness of real container engines like Docker.
- Networking setup is basic and may not cover all use cases.
- Error handling is minimal and may need improvement for real-world applications.

## Acknowledgements

This project uses the following libraries:

- [go-containerregistry](https://github.com/google/go-containerregistry) by Google
- Various Go standard libraries

## Disclaimer

This project is a parody and should be used with individual discretion. It is designed for educational purposes to help understand containerization concepts and is not intended for production use.

---

Feel free to explore, learn, and modify the code to suit your educational needs. Happy containerizing!