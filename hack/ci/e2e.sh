#!/bin/bash

cd /vagrant

export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin
export GOPROXY=proxy.golang.org,direct

# Deploy flotta operator
sed -i 's/LOG_LEVEL=info/LOG_LEVEL=debug/g' config/manager/controller_manager_config.properties
make build
IMG=flotta-operator:latest make docker-build
kind load docker-image flotta-operator:latest
timeout 2m make deploy IMG=flotta-operator TARGET=kind || retVal=$?
if [[ -n "$retVal" && $retVal -ne 124 ]]; then
  echo "Make run failed"; exit 1
fi
kubectl wait --timeout=120s --for=condition=Ready pods --all -n flotta
# TODO: add ingress resource to config/kind/kustomization.yaml instead of port forwarding
function port_forward() {
  while true; do
   kubectl port-forward deploy/flotta-operator-controller-manager -n flotta --address 0.0.0.0 8043:8043;
  done
}
port_forward &

# Run test
INPUT_IMAGE="${{ github.event.inputs.image }}"
TEST_IMAGE="${INPUT_IMAGE:-quay.io/project-flotta/edgedevice:latest}" make integration-test
