# Strace Setup for Seccomp Profile Generation

This directory contains the necessary files to generate a seccomp profile using strace in a kind cluster.

## Prerequisites

- Podman
- kind (Kubernetes in Docker) with experimental Podman provider
- kubectl

## Setup Instructions

### 1. Set the Kind provider to Podman

```bash
export KIND_EXPERIMENTAL_PROVIDER=podman
```

### 2. Build the strace image with Podman

```bash
podman build -t go-snapnote-backend-strace:latest -f strace/Dockerfile .
```

### 3. Create the kind cluster

```bash
cd strace
kind create cluster --name snapnote-strace --config kind-config.yaml
```

### 4. Load the image into kind

Save the image as an OCI archive and load it into the kind cluster:

```bash
# Save the image as an OCI archive
podman save --format oci-archive -o go-snapnote-backend-strace.tar go-snapnote-backend-strace:latest

# Load the image archive into kind
kind load image-archive go-snapnote-backend-strace.tar --name snapnote-strace
```

### 5. Deploy MariaDB and the application with strace

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

Wait for both MariaDB and the application pods to be ready:

```bash
# Wait for MariaDB to be ready
kubectl wait --for=condition=ready pod -l app=mariadb --timeout=120s

# Wait for the application pod to be ready
kubectl wait --for=condition=ready pod -l app=go-snapnote-backend-strace --timeout=120s
```

### 6. Generate traffic to capture system calls

Once both pods are running, generate some traffic:

```bash
# Port forward to access the service
kubectl port-forward service/go-snapnote-backend-strace 8080:8080

# In another terminal, make some requests
# Create a note (returns a note ID)
curl -X POST http://localhost:8080/post \
  -H "Authorization: Bearer test-token" \
  -H "Content-Type: application/json" \
  -d '{"test": "data"}'

# Read the note using the returned note ID (replace {noteid} with the actual ID from the POST response)
curl -X GET http://localhost:8080/get/{noteid} \
  -H "Authorization: Bearer test-token"
```

### 7. Extract strace logs

```bash
# Get the pod name
POD_NAME=$(kubectl get pods -l app=go-snapnote-backend-strace -o jsonpath='{.items[0].metadata.name}')

# Copy strace logs from the pod
kubectl cp $POD_NAME:/tmp/strace.log ./strace.log
```

### 8. Analyze strace output and generate seccomp profile

Analyze the `strace.log` file to identify all system calls used by your application. You can use tools like:

- `strace -c` for summary statistics
- Manual analysis of the log file
- Tools like `strace2seccomp` or similar utilities

### 9. Create the final seccomp profile

Based on the analysis, create a `seccomp-profile.json` file that allows only the necessary system calls. The profile should follow the Kubernetes seccomp profile format.

## Example seccomp profile structure

```json
{
  "defaultAction": "SCMP_ACT_ERRNO",
  "architectures": [
    "SCMP_ARCH_X86_64",
    "SCMP_ARCH_X86",
    "SCMP_ARCH_X32"
  ],
  "syscalls": [
    {
      "names": [
        "read",
        "write",
        "open",
        "close",
        "mmap",
        "munmap",
        "mprotect",
        "brk",
        "rt_sigaction",
        "rt_sigprocmask",
        "clone",
        "execve",
        "exit_group",
        "futex",
        "getpid",
        "gettid",
        "socket",
        "connect",
        "accept",
        "bind",
        "listen",
        "setsockopt",
        "getsockopt",
        "epoll_create1",
        "epoll_ctl",
        "epoll_wait"
      ],
      "action": "SCMP_ACT_ALLOW"
    }
  ]
}
```

## Cleanup

```bash
kind delete cluster --name snapnote-strace
rm -f go-snapnote-backend-strace.tar
```

## Notes

- This setup uses the experimental Kind Podman provider (`KIND_EXPERIMENTAL_PROVIDER=podman`)
- The strace container requires `SYS_PTRACE` capability to trace system calls
- MariaDB is included in the deployment.yaml file and will be automatically deployed
- The MariaDB service is named `mysql` to match the DB_HOST environment variable
- MariaDB uses an emptyDir volume for data persistence (data will be lost when the pod restarts)
- Make sure to generate realistic traffic patterns to capture all system calls your application uses, including database operations
- The audit.json profile logs all system calls without blocking them, which is useful for initial analysis
- Once you have the seccomp profile, you can use it in your production Dockerfile by mounting it in kind or using it as a seccomp profile in Kubernetes
- The image archive file (`go-snapnote-backend-strace.tar`) can be deleted after loading into kind
