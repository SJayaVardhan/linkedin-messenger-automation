# LinkedIn Automation Proof-of-Concept (Go + Rod)

> ⚠️ **Educational Purpose Only**
>
> This project is a technical proof-of-concept created strictly for evaluation and learning purposes.  
> Automating LinkedIn violates LinkedIn’s Terms of Service.  
> **Do NOT use this project on real accounts or in production environments.**

---

## 📌 Overview

This repository demonstrates an advanced browser automation system built in **Go** using the **Rod** library.  
The goal is to showcase:

- Human-like browser interaction
- Anti-detection / stealth techniques
- Clean, modular Go architecture
- Ethical automation practices

The project focuses on **how** automation works, not on abusing platforms.

---

## ✨ Features

### 🔐 Authentication
- Login using environment variables
- Graceful handling of login failures
- Detection of security checkpoints (CAPTCHA / 2FA)
- Persistent session cookies for reuse

### 🔍 Search & Targeting
- Search LinkedIn users by keyword
- Collect and deduplicate profile URLs
- Handle pagination
- Human-like scrolling behavior

### 🤝 Connection Requests
- Visit profiles programmatically
- Detect and click **Connect** button safely
- Skip already connected / pending profiles
- Optional personalized note
- Enforced rate limits

### 🥷 Anti-Detection Techniques
Implemented multiple stealth mechanisms including:
- Randomized delays and cooldowns
- Mouse hovering and cursor movement
- Human-like typing simulation
- Scroll acceleration / deceleration
- Browser fingerprint masking
- Session reuse (reduces bot signals)

> ⚠️ This project **does NOT** bypass CAPTCHA, 2FA, or security mechanisms by design.

---

## 🧱 Architecture

cmd/
└── app/
└── main.go

internal/
├── auth/ # Login & checkpoint detection
├── browser/ # Browser initialization
├── connect/ # Connection request logic
├── search/ # Profile search & pagination
├── stealth/ # Anti-detection techniques
├── storage/ # Cookie/session persistence
└── logger/ # Structured logging

yaml
Copy code

The codebase follows:
- Clear separation of concerns
- Idiomatic Go practices
- Defensive automation patterns

---

## ⚙️ Setup Instructions

### 1️⃣ Clone Repository

git clone <your-repo-url>
cd linkedin-messenger-automation-poc
2️⃣ Install Dependencies
bash
Copy code
go mod tidy
3️⃣ Environment Configuration
Create a .env file (do NOT commit it):

env

LINKEDIN_EMAIL=your_email
LINKEDIN_PASSWORD=your_password
HEADLESS=false
DAILY_CONNECT_LIMIT=2
Refer to .env.example for guidance.

▶️ Running the Project
go run cmd/app/main.go

Expected Flow
Browser launches with stealth enabled
Existing session cookies loaded (if available)
Login attempted only if required
Search performed
Profiles collected
Connection requests sent (within limits)