alias k=kubectl

k delete ing todolist-ingress

k delete svc todolist-app-svc todolist-db-svc

k delete deploy todolist-app-deploy todolist-db-deploy

k delete PodPreset todolist-app-preset todolist-db-preset

k delete PersistentVolumeClaim todolist-db-init-pvc

k delete PersistentVolume todolist-db-init-pv

k delete PersistentVolumeClaim todolist-db-pvc
