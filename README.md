# 🚀 Go Cloud Backend Systems

A focused collection of Go projects built to strengthen my cloud-native backend skills.  
This repository tracks my fast-tracked upskilling (12 hrs/day) after 1.5 years of experience as Cloud Platform Engineer at **E2E Networks Ltd.** (VM provisioning, scaling, load balancing & Django APIs).

**Goal**: Turn the employment gap into a strength by building production-style Go systems that directly extend my real cloud experience — ready for Golang Backend / Cloud Engineer roles.

---
## 📌 Projects Roadmap

### ✅ 1. CLI Cloud VM Task Scheduler + Worker Pool
**Status:** Completed (March 2026)  
**Folder:** [`step1-cli-worker`](./step1-cli-worker)

A command-line system that simulates real cloud VM provisioning using Go concurrency.  
**Real-world utility**: Mirrors exactly how AWS/E2E Networks handle background VM commissioning and scaling without blocking the main system.

**Learnings:**
- Mastered goroutines, channels, buffered queues, and worker pools
- Implemented context-based graceful shutdown and live task lifecycle logging
- Built background job processing patterns used in production cloud platforms
- Understood why Go concurrency is preferred for high-scale backend systems

**Link to project:** [step1-cli-worker](./step1-cli-worker)

---
### ⏳ 2. REST API (Gin Framework)
**Status:** To be done  
A high-performance RESTful API layer on top of the scheduler for handling VM operations over HTTP.

---
### ⏳ 3. API + Database (PostgreSQL + GORM + JWT)
**Status:** To be done  
Full backend with persistent storage, ORM, authentication, and secure endpoints.

---
### ⏳ 4. Worker Pool + Redis Queue
**Status:** To be done  
Production-grade asynchronous job queue with Redis for reliable background VM provisioning.

---
### ⏳ 5. Mini System (Full Integration)
**Status:** To be done  
Combine API + DB + Redis queue + worker pool into one cohesive backend system.

---
### 🔥 6. Main Project: Cloud VM / Inventory Management API
**Status:** To be done  
Production-style cloud infrastructure backend (Gin + GORM + Redis + Docker + AWS deployment).  
The flagship project that will close my gap and make me immediately interview-ready.

---
## ⚠️ Notes
- Learnings are added **only after** each project is fully completed and pushed.
- Every project is designed for **maximum employability** — concurrency, scalability, and direct connection to E2E Networks cloud work.
- Code quality, structure, and documentation improve with every step.
- All projects are built “learn-by-doing” style with 12-hour daily focus.

---

**Built during focused Golang upskilling (Sep 2024 – Present)** to return to the job market stronger in cloud-native backend development.
