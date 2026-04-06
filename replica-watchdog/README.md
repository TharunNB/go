# K8s Replica Enforcer

A lightweight Kubernetes controller built in Go that monitors cluster events and enforces a strict scaling policy: **No deployment shall exceed 3 replicas.** ## Features
* **Event-Driven:** Uses `SharedInformerFactory` for low-latency, resource-efficient monitoring.
* **Auto-Remediation:** Automatically patches non-compliant Deployments back to 3 replicas using `StrategicMergePatch`.
* **Dual-Informer:** Watches both **Pods** (for status logging) and **Deployments** (for policy enforcement).
* **Resync Logic:** Periodically verifies state to ensure eventual consistency even if events are missed.

## Prerequisites
* Go 1.21+
* A running Kubernetes cluster (Minikube, Kind, or Cloud)
* `kubectl` configured with a valid context

## Getting Started

1. **Initialize the module:**
   ```bash
   go mod init k8s-enforcer
   go mod tidy

2. **Run the controller:**
    ```bash 
    # It will use your default KUBECONFIG (~/.kube/config)
    go run main.go

3. **Test the Enforcement:**
    ```bash
    kubectl create deployment test-app --image=nginx
    kubectl scale deployment test-app --replicas=10

4. **Observe Logs:**
    The controller will log the violation and immediately trigger a patch to scale it back down to 3.