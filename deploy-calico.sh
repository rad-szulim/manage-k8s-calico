mage -v gen:registryCacheCert
docker pull amr-cache-registry.caas.intel.com/cache/calico/cni:v3.19.4
kind load docker-image amr-cache-registry.caas.intel.com/cache/calico/cni:v3.19.4
docker pull amr-cache-registry.caas.intel.com/cache/calico/pod2daemon-flexvol:v3.19.4
kind load docker-image amr-cache-registry.caas.intel.com/cache/calico/pod2daemon-flexvol:v3.19.4
docker pull amr-cache-registry.caas.intel.com/cache/calico/kube-controllers:v3.19.4
kind load docker-image amr-cache-registry.caas.intel.com/cache/calico/kube-controllers:v3.19.4
docker pull amr-cache-registry.caas.intel.com/cache/calico/node:v3.19.4
kind load docker-image amr-cache-registry.caas.intel.com/cache/calico/node:v3.19.4
kubectl -n kube-system apply -f calico2.yaml