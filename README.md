# Knative Dumper

Simple Cloud-events dumper service for debugging Knative calls.

## Deploy

Requires [ko](https://github.com/google/ko).

```
# eval $(minikube docker-env)
# export KO_DOCKER_REPO='ko.local'
ko apply -f config/
```

