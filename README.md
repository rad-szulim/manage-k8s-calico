# manage-k8s-calico
Use Golang client to interact with Calico CNI deployed to k8s cluster

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

Deploy Calico CNI:
```sh
sh deploy-calico.sh
```

New Calico pods as well as CoreDNS pods should be running:
```sh
kubectl -n kube-system get pods
```
