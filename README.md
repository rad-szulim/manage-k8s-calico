# manage-k8s-calico
Use Golang client to interact with Calico CNI deployed to k8s cluster

## Preques:

Issues discovered during investigation
https://github.com/projectcalico/calico/issues/6108

Install API server:
https://projectcalico.docs.tigera.io/maintenance/install-apiserver

## Setup kind cluster

Add latest port mappings to file based on virgo/controller/Contributing.md instructions for kind deployment.

Deploy kind cluster with Calico as the deault CNI:
```sh
sh kind-cluster-calico.sh
```

Verify that CoreDNS pods are pending:
```sh
kubectl -n kube-system get pods
```

Deploy Calico CNI (this file is created from https://projectcalico.docs.tigera.io/v3.19/manifests/calico.yaml by referencing amr Intel registry instead of docker.io):
```sh
sh deploy-calico.sh
```

New Calico pods as well as CoreDNS pods should be running:
```sh
kubectl -n kube-system get pods
```

## Build calico Docker container

```sh
make namespace
make build
make run
```
