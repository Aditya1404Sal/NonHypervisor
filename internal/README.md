# Mock Docker-Build

This is a mock Docker-like image building tool written in Go. It builds images from a configuration file, similar to `docker build`, with layer-by-layer building and caching.

## Usage

1. Create a configuration file like (`build.yaml`):

    ```yaml
    base_image: ubuntu:20.04
    layers:
      - run: apt-get update
      - run: apt-get install -y python3
      - copy:
          src: ./
          dest: /app
      - run: python3 /app/setup.py
    ```

2. Run the tool:

    ```sh
    go run cmd/mockdocker/main.go build.yaml
    ```


