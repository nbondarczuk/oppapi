#!/bin/bash

#
# Install NGINX ingress controllers on docker desktop and patch it to use port 8080
#

kubectl config current-context
kubectl config use-context docker-desktop
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.6.4/deploy/static/provider/cloud/deploy.yaml
kubectl patch svc -n ingress-nginx ingress-nginx-controller -p '{"spec": {"ports": [{"appProtocol": "http", "name": "http", "port": 8080, "protocol": "TCP", "targetPort": "http"}]}}' --type merge
kubectl -n ingress-nginx get pod -o wide
kubectl -n ingress-nginx get all -o wide

