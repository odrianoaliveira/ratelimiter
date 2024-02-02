# Go Rate Limiter using Token Bucket Algorithm
This project implements a rate limiter in Go, utilizing the token bucket algorithm.
The token bucket algorithm allows for flexible rate limiting, accommodating bursts of requests up to the capacity of the bucket, with tokens replenishing at a steady rate.
This approach is handy for APIs and web services, ensuring fair usage and preventing abuse.

## Features
- Token bucket rate limiting strategy.
- Flexible token replenishment rates and burst capacities.
- In-memory and Redis token storage options for scalability.

## Getting Started
You can follow these instructions to get the project up and running on your local machine for development and testing.

### Prerequisites
- Go (version 1.18 or later recommended).

### Installing
1. Clone the repository:

```bash
git clone https://github.com/yourusername/ratelimiter.git
cd ratelimiter
