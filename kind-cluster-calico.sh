kind create cluster --image kindest/node:v1.23.3 --config - <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 32000
    hostPort: 32000
  - containerPort: 32001
    hostPort: 32001
  - containerPort: 32002
    hostPort: 32002
  - containerPort: 32003
    hostPort: 32003
  - containerPort: 32004
    hostPort: 32004
  - containerPort: 32005
    hostPort: 32005
  - containerPort: 32006
    hostPort: 32006
  - containerPort: 32007
    hostPort: 32007
  - containerPort: 32008
    hostPort: 32008
  - containerPort: 32009
    hostPort: 32009
  - containerPort: 32010
    hostPort: 32010
  - containerPort: 32011
    hostPort: 32011
  - containerPort: 32012
    hostPort: 32012
  - containerPort: 30002
    hostPort: 30002
  - containerPort: 30003
    hostPort: 30003
  - containerPort: 32013
    hostPort: 32013
  - containerPort: 32014
    hostPort: 32014
  - containerPort: 32015
    hostPort: 32015
  - containerPort: 32016
    hostPort: 32016
  - containerPort: 32017
    hostPort: 32017
networking:
  podSubnet: 10.240.0.0/16
  serviceSubnet: 10.110.0.0/16
  disableDefaultCNI: true
EOF
