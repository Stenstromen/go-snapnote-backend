#!/bin/bash
set -e

# Get the pod name
POD_NAME=$(kubectl get pods -l app=go-snapnote-backend-strace -o jsonpath='{.items[0].metadata.name}')

if [ -z "$POD_NAME" ]; then
    echo "Error: No pod found with label app=go-snapnote-backend-strace"
    exit 1
fi

echo "Found pod: $POD_NAME"
echo "Extracting strace logs..."

# Copy strace logs from the pod
kubectl cp "$POD_NAME:/tmp/strace.log" ./strace.log

echo "Strace logs extracted to ./strace.log"
echo ""
echo "To analyze system calls, you can run:"
echo "  grep -oP '^[0-9]+\s+\K\w+' strace.log | sort | uniq -c | sort -rn"
echo ""
echo "Or use strace summary format:"
echo "  strace -c -r -f -e trace=all -o strace-summary.log <your-command>"

