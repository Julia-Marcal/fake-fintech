#!/bin/bash
# wait-for-it.sh
# Adapted from https://github.com/vishnubob/wait-for-it

# This script is used to wait for a service to become available before running a command.

# Usage:
#   ./wait-for-it.sh host:port [-s] [-t timeout] -- command args
#   -s: Use SSL/TLS (default: false)
#   -t: Timeout in seconds (default: 15 seconds)
#   host:port: The host and port of the service to wait for.
#   command args: The command with its arguments to run once the service is available.

# Parse command-line options
while [[ $# -gt 0 ]]
do
  case "$1" in
    *:* )
    hostport=(${1//:/ })
    host=${hostport[0]}
    port=${hostport[1]}
    shift 1
    ;;
    -s )
    scheme="https"
    shift 1
    ;;
    -t )
    timeout="$2"
    shift 2
    ;;
    -- )
    shift
    break
    ;;
    * )
    echo "Unknown argument: $1"
    exit 1
    ;;
  esac
done

# Set default values if not provided
timeout=${timeout:-15}
scheme=${scheme:-http}

# Function to check if the service is available
wait_for_service() {
  echo "Waiting for service to become available..."
  local start_time=$(date +%s)
  local end_time=$((start_time + timeout))

  while true; do
    if [ "$scheme" = "https" ]; then
      # Use openssl to check if the service is available via SSL/TLS
      if openssl s_client -connect "$host:$port" </dev/null >/dev/null 2>&1; then
        break
      fi
    else
      # Use netcat to check if the service is available via plain TCP
      if nc -z "$host" "$port"; then
        break
      fi
    fi

    sleep 1

    local current_time=$(date +%s)
    if [ "$current_time" -ge "$end_time" ]; then
      echo "Timeout reached. Service is not available."
      exit 1
    fi
  done

  local elapsed_time=$((current_time - start_time))
  echo "Service is available after $elapsed_time seconds."
}

# Wait for the service to become available
wait_for_service

# Run the specified command
exec "$@"
