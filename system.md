
# SonicStream Architecture Plan

SonicStream is a full-stack multimedia conversion platform. Below is the proposed architecture design, technology stack, and hosting strategy, optimized for minimal operational costs (under $20/month).

---

## Client

- **Framework:** React
- **Bundler:** Vite
- **Styling:** Tailwind CSS
- **Routing:** React Router (for multiple pages)
- **Hosting:** [AWS Amplify](https://aws.amazon.com/amplify/)
  - Supports CI/CD from GitHub
  - SSL, custom domain support
- **Optional:** [Route53](https://aws.amazon.com/route53/) for domain management

**Frontend Pages:**
- Homepage with file uploader
- File history/dashboard
- Subscription/pricing page
- Login & Signup pages
- User settings

---

## Backend

- **Language:** Golang
- **HTTP Server:** Go standard library (net/http)
- **Responsibilities:**
  - Handle multipart file uploads
  - Validate file types and convert using FFmpeg
  - Generate & return download links
  - Interact with DynamoDB & S3

### Hosting Suggestions:
- **Option 1:** [Fly.io](https://fly.io/) — Free tier + global deployment
- **Option 2:** [Railway.app](https://railway.app/) — Great for small projects
- **Option 3:** [Render](https://render.com/) — Free tier suitable for Go apps
- **Option 4 (cheap VPS):** [Linode](https://www.linode.com/), [Hetzner](https://www.hetzner.com/)

---
### StorageS3 Bucket (AWS)

- Store converted multimedia files
- Public download URLs with limited TTL
- Lifecycle policy to delete files older than X days to save costs

---
### DatabaseDynamoDB (AWS)
- Store file metadata, user data, and logs
- Lightweight and cost-efficient with on-demand billing
- Can implement rate limiting tracking via TTL attributes and counters

---
### Cost StrategyTarget: <$20/month total
- Steps to Control Costs:
- Use free tiers (AWS Free Tier, Render, Amplify, etc.) as much as possible
- Enable DynamoDB TTL to delete stale records
- Configure S3 lifecycle rules to auto-delete old files
- Avoid autoscaling on all services
- Gracefully fail requests when limits are reached (return 429/500 as fallback)

