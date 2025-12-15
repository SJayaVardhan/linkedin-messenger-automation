🎯 Project Objective
Build a Go-based LinkedIn automation PoC that demonstrates:
Advanced browser automation using Rod
Human-like interaction behavior
Anti-bot and stealth techniques
Ethical handling of security checkpoints (CAPTCHA / 2FA)
Clean, modular, maintainable Go architecture

🧩 Core Features
✅ Authentication System
Login using environment variables
Graceful handling of login failures
Detection of security checkpoints (CAPTCHA / 2FA)
Human-in-the-loop login recovery
Persistent session cookies for reuse

🔍 Search & Targeting
Search users by keyword (e.g., job title)
DOM parsing to extract profile URLs
Pagination and scrolling support
Duplicate profile handling

🤝 Connection Requests
Programmatic navigation to profiles
Precise detection of the Connect action
Optional personalized notes
Rate limiting and daily caps
Safe skipping when Connect is unavailable (Follow / limits / UI changes)
Connection success is not guaranteed and not required.

🕵️ Anti-Bot & Stealth Techniques
Implemented stealth mechanisms include:
Browser fingerprint masking (navigator.webdriver)
Randomized delays and think times
Human-like scrolling behavior
Mouse hovering and movement simulation
Typing simulation with variable speed
Rate limiting and cooldown enforcement
Session reuse to avoid repeated logins
Ethical handling of CAPTCHA / 2FA (no bypassing)

🗂️ Project Structure
linkedin-messenger-automation-poc/
│
├── cmd/
│   └── app/
│       └── main.go          # Application entry point
│
├── internal/
│   ├── auth/                # Login & authentication logic
│   ├── browser/             # Browser initialization
│   ├── connect/             # Connection request workflow
│   ├── logger/              # Structured logging
│   ├── message/             # Messaging (PoC placeholder)
│   ├── search/              # Search & profile extraction
│   ├── stealth/             # Anti-detection techniques
│   └── storage/             # Session cookie persistence
│
├── configs/
│   └── config.yaml          # Optional configuration file
│
├── .env.example             # Environment variable template
├── .gitignore               # Git ignore rules
├── go.mod / go.sum          # Go module files
└── README.md


⚙️ Setup Instructions
1️⃣ Prerequisites
Go 1.20+
Google Chrome / Chromium
macOS / Linux / Windows

2️⃣ Environment Configuration
Create a .env file using the template: cp .env.example .env
Example .env.example:
LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_password
HEADLESS=false
DAILY_CONNECT_LIMIT=10

3️⃣ Install Dependencies
go mod tidy

4️⃣ Run the Application
go run cmd/app/main.go

🔐 CAPTCHA & Manual Login Handling
If LinkedIn presents a CAPTCHA or 2FA:
The program pauses automatically
Browser remains open
User completes login manually
Program detects successful login
Session cookies are saved
Execution resumes automatically
This behavior is intentional and ethical.

📊 Expected Runtime Flow
Launch Browser
→ Apply Stealth
→ Load Session Cookies
→ Detect Login State
→ Handle CAPTCHA if needed
→ Save Session
→ Search Profiles
→ Attempt Connect Requests (rate-limited)

🏁 Final Notes
This project demonstrates:
Real-world automation challenges
Responsible engineering practices
Clean Go architecture
Advanced browser automation concepts
Not intended for real-world use.
