# CropTrace Telemetry

A robust, multi-tier microservice for ingesting agricultural telemetry data. 

## Architecture
- **Stateless API Tier**: Built in Go using the Gin web framework and GORM for database operations. It implements graceful shutdown and structured JSON logging.
- **Data Buffer Tier**: A lightweight PostgreSQL database deployed as a StatefulSet with persistent storage to buffer the telemetry data.
- **Orchestration**: Kubernetes manifests (Namespace, ConfigMap, Secrets, StatefulSet, Deployment, HPA, Services).

## API Specifications
- `POST /api/v1/telemetry`: Accepts JSON data (`facility_id`, `timestamp`, `crop_type`, `weight_kg`, `quality_rating`).
- `GET /api/v1/health`: Readiness/liveness probe checking DB connectivity.

## CI/CD Pipeline
GitHub Actions automatically builds the Docker image, pushes it to DockerHub, and deploys it to a `kind` cluster to verify integration and configuration.

## Deployment Instructions

### Local Testing with Docker/Kind
1. Configure your DockerHub credentials in GitHub Secrets (`DOCKER_USERNAME`, `DOCKER_PASSWORD`) to enable the CI/CD pipeline.
2. If you are running on Kind, install the `local-path` provisioner before applying manifests:
   ```bash
   kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/master/deploy/local-path-storage.yaml
   ```
3. To test locally on your own machine using Kind:
   ```bash
   # Load k8s manifests
   kubectl apply -f k8s/
   
   # Verify deployment
   kubectl get all -n croptrace
   ```

### Security Best Practices Implemented
- **Multi-stage Docker Builds**: Small, secure production images using `alpine`.
- **Non-root Container User**: Security best-practice.
- **Headless Service DNS**: For internal database cluster communication.
- **Secrets Management**: Configuration decoupled from secrets. 
- **Resource Constraints**: CPU/Memory limits and Horizontal Pod Autoscaler (HPA) configured.