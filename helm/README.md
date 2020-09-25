# Helm Chart for Kubernetes Deployment

This helm chart deploys the docker container into a Kubernetes cluster.

You will need to provide your own values.yml with all custom (secret) values for your deployment hosting.  Once you've got that, you should be able to use Helm to deploy:

```bash
helm install . -f /path/to/your/secret/values.yml -n postal-rest
```
