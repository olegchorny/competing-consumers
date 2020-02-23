## Install nginx Ingress controller via helm v3
helm install nginx-default stable/nginx-ingress --namespace default --set controller.replicaCount=2 --set controller.metrics.enabled=true --set controller.podAnnotations."prometheus\.io/scrape"="true" --set controller.podAnnotations."prometheus\.io/port"="10254"
