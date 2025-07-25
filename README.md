# GopherSMTP

GopherSMTP is an ethical phishing simulation tool written in Go that lets you craft and send emails and SMS messages via SMTP. It's designed for red teamers, security awareness trainers, and IT admins who want to test how users respond to common social engineering techniques.
🚨 Ethical Disclaimer

    GopherSMTP is intended for educational and authorized security testing only. Do not use this tool on networks, people, or systems you do not own or have explicit permission to test.

## ✉️ Features

    Send custom or templated phishing emails to single or multiple recipients

    SMS-over-SMTP support: send texts using carrier gateways

    Supports injection into templates for dynamic customization

    Minimal, terminal-based UI

    Compatible with Gmail, Outlook, and other SMTP servers

## 🖥️ Menu Overview
 ` ` `
   ______            __              _____ __  _____________ 
  / ____/___  ____  / /_  ___  _____/ ___//  |/  /_  __/ __ \
 / / __/ __ \/ __ \/ __ \/ _ \/ ___/\__ \/ /|_/ / / / / /_/ /
/ /_/ / /_/ / /_/ / / / /  __/ /   ___/ / /  / / / / / ____/ 
\____/\____/ .___/_/ /_/\___/_/   /____/_/  /_/ /_/ /_/      
          /_/                                                
 ` ` `
SMTP attack tool
List:
0: quit 
1: send single email 
2: Phish from a template 
3: send to multiple emails 
4: smsOverSmtp
5: Send smsOverSmtp to multiple numbers

🛠️ Usage
🔐 Requirements

    Go 1.18+

    Access to an SMTP server (e.g. Gmail SMTP)

    Allow less secure apps or app passwords (for Gmail)

📦 Installation

git clone https://github.com/yourusername/gophersmtp.git
cd gophersmtp
go build -o gophersmtp
./gophersmtp

📧 Templated Email Example

Your email template (e.g. email_template.txt) can use:

Hello {{.A}}, please reset your password at {{.B}}.

Then GopherSMTP injects values into {{.A}} and {{.B}}.
🧪 Examples
Send a Single Email

Choose option 1
Enter recipient email
Enter subject and message body

Phish from Template

Choose option 2
Paste email, subject, template file
Inject name and link variables

SMS Over SMTP

Choose option 4 or 5
Enter phone number(s)
Use known gateway or try guessing
Send via email-to-text

📁 File Format
Email List File (for option 3):

alice@example.com
bob@example.com

SMS List File (for option 5):

4165551234
5146669876

⚖️ Legal Notice

This tool must only be used in environments where you have explicit permission. Misuse can lead to legal consequences.
💡 Why GopherSMTP?

    Lightweight and fast

    Go-native and cross-platform

    Extensible and easy to modify

    Great for phishing simulation and training

📃 License

Licensed software is fucking gay
