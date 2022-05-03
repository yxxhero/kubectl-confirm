# kubectl-confirm
Confirm before run kubectl.   
This not a kubectl plugin.   
Add the logic of cluster information confirmation in kubectl default behavior.  


# disable
```bash
# it same as the Official kubectl
export KUBECTL_CONFIRM_DISABLE=1
```
# example
```bash
+----------+----------+-----------+-------------------------+-----------------------+----------+-------------+
| CONTEXT  | CLUSTER  | NAMESPACE | API SERVER              | INSECURESKIPTLSVERIFY | AUTHINFO | CLUSTERNAME |
+----------+----------+-----------+-------------------------+-----------------------+----------+-------------+
| kind-k8s | kind-k8s | Unset     | https://127.0.0.1:46335 | false                 | kind-k8s | kind-k8s    |
+----------+----------+-----------+-------------------------+-----------------------+----------+-------------+
Continue (y/n):y
NAME                                        READY   STATUS    RESTARTS   AGE
coredns-74ff55c5b-d5pwr                     1/1     Running   1          6d8h
coredns-74ff55c5b-rv26v                     1/1     Running   1          6d8h
etcd-k8s-control-plane                      1/1     Running   1          6d8h
kindnet-5988d                               1/1     Running   1          6d8h
kindnet-kbqlz                               1/1     Running   1          6d8h
kube-apiserver-k8s-control-plane            1/1     Running   1          6d8h
kube-controller-manager-k8s-control-plane   1/1     Running   1          6d8h
kube-proxy-7vtrh                            1/1     Running   1          6d8h
kube-proxy-lhcg8                            1/1     Running   1          6d8h
kube-scheduler-k8s-control-plane            1/1     Running   1          6d8h
```