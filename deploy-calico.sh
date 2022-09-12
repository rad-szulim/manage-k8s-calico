docker pull calico/cni:v3.24.1
kind load docker-image calico/cni:v3.24.1
docker pull calico/pod2daemon-flexvol:v3.24.1
kind load docker-image calico/pod2daemon-flexvol:v3.24.1
docker pull calico/kube-controllers:v3.24.1
kind load docker-image calico/kube-controllers:v3.24.1
docker pull calico/node:v3.24.1
kind load docker-image calico/node:v3.24.1

# calico.yaml is downloaded from this link below since it can be moved or updated
# kubectl -n kube-system apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.24.1/manifests/calico.yaml
# 
kubectl -n kube-system apply -f ./calico.yaml