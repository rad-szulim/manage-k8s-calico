docker pull calico/cni:v3.23.2
kind load docker-image calico/cni:v3.23.2
docker pull calico/pod2daemon-flexvol:v3.23.2
kind load docker-image calico/pod2daemon-flexvol:v3.23.2
docker pull calico/kube-controllers:v3.23.2
kind load docker-image calico/kube-controllers:v3.23.2
docker pull calico/node:v3.23.2
kind load docker-image calico/node:v3.23.2

kubectl -n kube-system apply -f https://projectcalico.docs.tigera.io/v3.23/manifests/calico.yaml