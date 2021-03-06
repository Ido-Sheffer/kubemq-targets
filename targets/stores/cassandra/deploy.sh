helm repo add bitnami https://charts.bitnami.com/bitnami
helm install cassandra bitnami/cassandra --set dbUser.password=cassandra --set dbUser.user=cassandra
kubectl port-forward --namespace default svc/cassandra 9042:9042 & cqlsh -u cassandra -p $CASSANDRA_PASSWORD 127.0.0.1 9042
helm delete cassandra

