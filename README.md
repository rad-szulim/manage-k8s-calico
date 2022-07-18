# manage-k8s-calico
Use Golang client to interact with Calico CNI deployed to k8s kind cluster.

## Delete kind k8s cluster

Delete prior installation of kind k8s cluster:
```sh
kind delete cluster
```

## Setup kind cluster

If you plan on installing Controller using `mage`, add latest port mappings to file based on virgo/controller/Contributing.md instructions for kind deployment.

Deploy kind cluster disabling the deault CNI:
```sh
sh kind-cluster-no-cni.sh
```

Verify that CoreDNS pods are in `Pending` state:
```sh
kubectl -n kube-system get pods
```

Deploy Calico CNI:
```sh
sh deploy-calico.sh
```

New Calico pods as well as CoreDNS pods should be in `Running` state:
```sh
kubectl -n kube-system get pods
```

## Preqs for using Calico API:

Issues discovered during investigation:
https://github.com/projectcalico/calico/issues/6108


Install API server (as demonstrated in https://projectcalico.docs.tigera.io/maintenance/install-apiserver using https://projectcalico.docs.tigera.io/manifests/apiserver.yaml with mods):
```sh
kubectl create -f apiserver.yaml
```

Note: please be aware that the prior script was modfied to remove a line with `-v=5` from the script provided by Calico.

The pod should be in `ContainerCreating` state:
```sh
kubectl -n calico-apiserver get pods
```

Use Docker container to create key and cert (works for mac too):
```sh
docker run -v $(pwd):/apps -w /apps --name alpine_openssl --rm -i -t alpine/openssl req -x509 -nodes -newkey rsa:4096 -keyout apiserver.key -out apiserver.crt -days 365 -subj "/" -addext "subjectAltName = DNS:calico-api.calico-apiserver.svc"
```

Create a secret from the cert generated in the previous step:
```sh
kubectl create secret -n calico-apiserver generic calico-apiserver-certs --from-file=apiserver.key --from-file=apiserver.crt
```

Load docker image to kind:
```sh
docker pull calico/apiserver:v3.23.2
kind load docker-image calico/apiserver:v3.23.2
```

Configure Calico API server with the CA bundle:
```sh
kubectl patch apiservice v3.projectcalico.org -p \
    "{\"spec\": {\"caBundle\": \"$(kubectl get secret -n calico-apiserver calico-apiserver-certs -o go-template='{{ index .data "apiserver.crt" }}')\"}}"
```

The pod should be in `Running` state:
```sh
kubectl -n calico-apiserver get pods
```

## Build Calico Docker container

Deploy Docker container
```sh
make namespace
make build
make run
```

Verify that there are no errors in the log of the container:
```sh
kubectl -n smartedge-system get pods
kubectl -n smartedge-system logs <pod-id>
```

The code running in the container is interacting with Calico using Golang client.

Check that BGP Config and Peer were setup:
```sh
calicoctl get bgpconfig
calicoctl get bgppeer
```

Remove deployment:
```sh
make clean
```

Remove BGP Config and Peer:
```sh
calicoctl delete bgpconfig default
calicoctl delete bgppeer my-peer-1
```
