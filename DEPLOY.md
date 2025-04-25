# Cloud Deployment Guide

This guide covers deploying the Go URL Shortener to popular cloud providers.

## Docker (Any Cloud or VPS)
1. Build the Docker image:
   ```bash
   docker build -t shorten-it .
   ```
2. Run the container:
   ```bash
   docker run -e STORE_TYPE=boltdb -e API_KEY=yourkey -p 8080:8080 shorten-it
   ```

## Heroku
1. Add a `Procfile` with:
   ```
   web: ./shorten-it
   ```
2. Deploy using Heroku CLI:
   ```bash
   heroku create
   heroku stack:set container
   heroku container:push web
   heroku container:release web
   heroku open
   ```

## Google Cloud Run
1. Build and push the image:
   ```bash
   gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/shorten-it
   ```
2. Deploy to Cloud Run:
   ```bash
   gcloud run deploy shorten-it --image gcr.io/YOUR_PROJECT_ID/shorten-it --platform managed --region YOUR_REGION --allow-unauthenticated --set-env-vars STORE_TYPE=boltdb,API_KEY=yourkey
   ```

## AWS Elastic Beanstalk (Docker)
1. Initialize Elastic Beanstalk:
   ```bash
   eb init -p docker shorten-it
   eb create shorten-it-env
   eb open
   ```

## Netlify (Static Frontend Only)
For a Go backend, use Netlify Functions or deploy the Docker image elsewhere and connect via API.

---
For all cloud deployments, set environment variables for `STORE_TYPE`, `API_KEY`, and any backend storage config (`DATABASE_URL`, `REDIS_ADDR`, etc).

For more advanced setups (Postgres, Redis, TLS, etc.), refer to your cloud provider's documentation.
