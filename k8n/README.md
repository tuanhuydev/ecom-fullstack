# Kubernetes Manifests

This directory contains Kubernetes manifests organized by resource type.

## Directory Structure

```
k8n/
├── deployments/     # Application deployments
├── services/        # Service definitions
├── config/          # ConfigMaps for non-sensitive configuration
├── secrets/         # Secret templates (samples only)
├── storage/         # PersistentVolumeClaims
└── README.md        # This file
```

## Usage

### 1. Configure Secrets

Before deploying, create actual secret files from the samples:

```bash
# Copy sample files and add your actual values
cp secrets/postgres-secret.sample.yml secrets/postgres-secret.yml
cp secrets/go-api-secret.sample.yml secrets/go-api-secret.yml

# Edit the secret files with your actual base64-encoded passwords
# Example: echo -n "your_password" | base64
```

### 2. Deploy Resources

Deploy in the following order:

```bash
# 1. Storage
kubectl apply -f storage/

# 2. ConfigMaps
kubectl apply -f config/

# 3. Secrets (your actual secret files, not samples)
kubectl apply -f secrets/postgres-secret.yml
kubectl apply -f secrets/go-api-secret.yml

# 4. Deployments
kubectl apply -f deployments/

# 5. Services
kubectl apply -f services/
```

### 3. One-command deployment

```bash
# Deploy everything (make sure secrets are configured first)
kubectl apply -f storage/ -f config/ -f secrets/ -f deployments/ -f services/
```

## Security Notes

- Never commit actual secret files to version control
- Only commit `.sample.yml` files in the secrets directory
- Add `secrets/*.yml` to `.gitignore` (but allow `secrets/*.sample.yml`)
- Use tools like sealed-secrets or external secret managers for production

## Files

### Deployments
- `go-api-deployment.yml` - Go API application
- `nuxt-app-deployment.yml` - Nuxt.js frontend
- `postgres-deployment.yml` - PostgreSQL database
- `redis-deployment.yml` - Redis cache

### Services
- `go-api-service.yml` - Go API service
- `nuxt-app-service.yml` - Nuxt.js service
- `postgres-service.yml` - PostgreSQL service
- `redis-service.yml` - Redis service

### Storage
- `postgres-pvc.yml` - PostgreSQL persistent volume
- `redis-pvc.yml` - Redis persistent volume

### Config
- `go-api-configmap.yml` - Go API non-sensitive configuration
- `postgres-configmap.yml` - PostgreSQL non-sensitive configuration

### Secrets (samples)
- `postgres-secret.sample.yml` - PostgreSQL password template
- `go-api-secret.sample.yml` - Go API secrets template
