# Go Learning Lab: SRE & DevOps

A collection of hands-on projects focused on mastering **Golang** through the lens of **Site Reliability Engineering (SRE)** and **DevOps** principles. This repository serves as a sandbox for exploring automation, high-concurrency patterns, and Kubernetes internals.

## 🛠 Project Categories

### 1. Cloud-Native Orchestration (Kubernetes)
*Focus: Understanding the K8s Control Plane and Operator patterns.*
* **`replica-watchdog`**: An automated enforcer that monitors Pod health and maintains desired state using **Shared Informers**.
* **`k8s-pod-log-aggregator`**: A high-concurrency tool for streaming logs across distributed systems using **Semaphores** and **errgroups** to manage resource limits.

### 2. High-Concurrency & Reliability
*Focus: Building resilient systems that handle scale efficiently.*
* **`health-checker`**: A distributed URL prober that utilizes **Goroutines** and **Channels** to execute parallel health checks with minimal latency.
* **`step1-cli-worker`**: A task-scheduling engine simulating a Cloud VM provisioning worker pool to handle background jobs.

### 3. Backend Foundations & Tooling
*Focus: Service architecture and developer productivity.*
* **`http-server`**: A modular boilerplate for building high-performance Go microservices.
* **`todo`**: A CLI utility focused on state management, file I/O, and building fast developer productivity tools.

## 🧠 Learning Objectives

Through these projects, I am exploring the core pillars of SRE:
* **Automation:** Replacing manual toil with software-defined controllers (e.g., `replica-watchdog`).
* **Observability:** Aggregating logs and metrics from distributed environments (e.g., `log-aggregator`).
* **Performance:** Leveraging Go's concurrency primitives to build lightweight, high-speed infrastructure tools.

## 🚀 How to Use

Each directory is a self-contained Go module. To explore a project:

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/TharunNB/go.git](https://github.com/TharunNB/go.git)
   cd go
