<img src="asset/logo.png" alt="My Image" width="300" height="400">

# Authora - Telegram 2FA Bot Service

**Authora** is a Go-based Telegram bot that provides authentication (2FA) services for businesses. Instead of using email or SMS for OTP verification, businesses can integrate with this bot to send OTPs to users through Telegram.

---

## Features

- **Business Sign-up**: Businesses can register with the bot and receive a unique token for authentication.
- **OTP Generation**: Businesses can request the bot to send OTPs for user verification.
- **Secure**: Uses Telegram as a secure and free communication channel for OTP delivery.
- **PostgreSQL**: Stores business data in a PostgreSQL database.
- **Redis**: Optional cache layer for storing temporary data like OTPs.

---

## Prerequisites

To run this project, you need to have the following installed:

- **Go**: The programming language used to build the bot (https://golang.org/dl/)
- **PostgreSQL**: The database used for storing business information (https://www.postgresql.org/download/)
- **Redis** (Optional): Caching layer (https://redis.io/download)
- **Telegram Bot API Token**: You can get this from the BotFather on Telegram.

---

## Setup

### 1. Install Dependencies

Make sure you have the necessary Go modules installed. Run the following command to get all dependencies:

```bash
go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
go get -u golang.org/x/net/proxy
go get -u github.com/jinzhu/gorm
go get -u github.com/lib/pq
