alias k=kubectl

k apply -f ./deployment/k8/pv-db-init.yaml

k apply -f ./deployment/k8/pvc-db-init.yaml

k apply -f ./deployment/k8/pvc-db.yaml

k apply -f ./deployment/k8/pods/db-pod-preset.yaml

k apply -f ./deployment/k8/pods/app-pod-preset.yaml

#k apply -f ./deployment/k8/pods/db-pod.yaml

#k apply -f ./deployment/k8/pods/app-pod.yaml

#k apply -f ./deployment/k8/replica-set/db-replicaset.yaml

#k apply -f ./deployment/k8/replica-set/app-replicaset.yaml

k apply -f ./deployment/k8/db-deployment.yaml

k apply -f ./deployment/k8/app-deployment.yaml

k apply -f ./deployment/k8/db-service.yaml

k apply -f ./deployment/k8/app-service.yaml

k apply -f ./deployment/k8/ingress.yaml



