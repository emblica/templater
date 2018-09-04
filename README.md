# templater
Combine template and JSON-formatted data

Uses Django-like templating library pongo2 (https://github.com/flosch/pongo2)

Use it as Docker container or deploy into Kubernetes as initContainer

## Usage
```
# Pull the image
docker pull emblica/templater
# Run it
docker run -it --rm -v $(pwd)/example:/mnt/example emblica/templater /mnt/example/data.json /mnt/example/template.md /mnt/example/output.md
```

## Usage with Kubernetes

```
# Apply pod:
kubectl apply -f k8s/example.yml
```

## Build
```
docker build -t templater .
```
