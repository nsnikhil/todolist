helm install --name-template todolist-prerequisite ./deployment/helm/todolist_prerequisite

echo "wait for volumes to start..."
sleep 120

helm install --name-template todolist ./deployment/helm/todolist