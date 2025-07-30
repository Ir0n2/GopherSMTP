
# GopherSMTP

**GopherSMTP** is an SMTP-based phishing tool written in Go that automates email and SMS message delivery via SMTP gateways. It’s a terminal-driven app designed for sending spoofed or prank messages, including through SMS-to-email gateways for added realism.

> **⚠️ Disclaimer:** This tool is for **educational** and **ethical testing** purposes only. Do not use GopherSMTP to harass, threaten, or deceive individuals without their consent. Always comply with local laws and terms of service. Or don't. I can't stop you.

---

## 💡 What Can GopherSMTP Do?

GopherSMTP lets you send **emails and SMS messages** from the terminal, with the following features:

```
1. Send a single email
2. Send to multiple email addresses
3. Send SMS via SMTP (smsOverSmtp)
4. Send SMS via SMTP to multiple numbers
5. Set sender email (via AttackerEmails.txt)
6. Spam emails
7. Spam smsOverSmtp
```

---

## 🌐 Planned Feature

* 🔮 **AI-Powered Phishing**: Integrate **Google Gemini** to auto-generate convincing phishing messages.

---

## 🧠 Why Use This?

> Use it to scare your friends with creepy texts or fake emails. For example, I once used it to tell my friend (pretending to be his hospital) that his vaccine was contaminated with HIV. He didn’t sleep for two days. That’s the kind of prank power we’re talking about.

---

## 📦 How to Install

1. **Install Go** ([https://go.dev/doc/install](https://go.dev/doc/install))

1.5. Install figlet
```bash
sudo apt install figlet
```

2. Clone this repo:

   ```bash
   git clone https://github.com/Ir0n2/GopherSMTP
   cd GopherSMTP
   ```

3. Create a file called `AttackerEmails.txt` with your sender email and app password:

   ```
   your_email@gmail.com
   your_app_password
   ```

4. Run it:

   ```bash
   go run email.go
   ```

---

## ⚙️ Compatibility

You can use this on **any OS** with a Go compiler and internet access — Windows, macOS, Linux, WSL — it doesn’t matter.

---

## 👨‍💻 Who Made This?

Made by **Ir0n2**. Except for this read me. I got lazy and forced chatgpt to make it all pretty.

---

## 🛑 Legal Disclaimer

This tool is for **educational and research purposes only**. The developer is **not responsible** for any misuse of this software. Do not use this to impersonate institutions, harass people, or violate laws.

---
