
### DWRF -Debugging with arbitrary record format

Distributed Web Remote Function (DWRF) é um protocolo de comunicação que facilita a execução remota de funções em um ambiente distribuído. No entanto, não há uma implementação específica do DWRF em Go.
 

### Gerando a imagem

    $ docker build -t paulojr83/deploy-k8s:latest -f Dockerfile.prod .

### Docker push
    $ docker push paulojr83/deploy-k8s

### kind

[link](https://kind.sigs.k8s.io/)

    $ go install sigs.k8s.io/kind@v0.22.0 && kind create cluster
    $ kind create cluster --name=deploy-k8s

Not sure what to do next? 😅  [Check out](https://kind.sigs.k8s.io/docs/user/quick-start/)

    $ kubectl cluster-info --context kind-deploy-k8s

    $ kubectl apply -f k8s/deployment.yaml
    
    $ kubectl get pods
    NAME                      READY   STATUS    RESTARTS   AGE
    server-6644575555-bxt77   1/1     Running   0          35s
    server-6644575555-cn4mq   1/1     Running   0          35s
    server-6644575555-gd2mk   1/1     Running   0          35s
    server-6644575555-jhbcn   1/1     Running   0          35s
    server-6644575555-qz64m   1/1     Running   0          35s

    $ kubectl delete pod server-6644575555-bxt77

    $ kubectl apply -f k8s/service.yaml
    service/serversvc created

    $ kubectl get svc
    NAME         TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
    kubernetes   ClusterIP   10.96.0.1     <none>        443/TCP   18m
    serversvc    ClusterIP   10.96.93.97   <none>        80/TCP    38s

    $ kubectl port-forward svc/serversvc 8080:80