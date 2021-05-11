# klocust
Klocust is a command-line tool for managing [Locust](https://locust.io/) distributed load testing on [Kubernetes](https://kubernetes.io/).  


## Installation
### Required
- [kubernetes](https://kubernetes.io/) version 1.16 or higher
- kubeconfig for connect to your k8s cluster

### Recommend
- [Ingress Controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/)   
   exposes your locust cluster to outside of the k8s cluster. (internal or internet-facing)
- [ExternalDNS](https://github.com/kubernetes-sigs/external-dns)   
   create(update) and delete domain for your locust cluster automatically.

### Install klocust binary
```bash
$ brew tap DevopsArtFactory/devopsart
$ brew update
$ brew install klocust 
$ klocust version
0.0.1
``` 

### Build from source
- See here on how to [build klocust CLI from source](./docs/build_from_source.md).

### shell autocompletion for bash/zsh
```bash
echo 'source <(kubectl completion bash)' >>~/.bashrc
or
echo 'source <(kubectl completion zsh)' >>~/.zsh
```

## Usages

### klocust list
- Display all of Locust clusters
```bash
$ klocust list

>>> 1 locust deployments in loadtest namespace. (PREFIX: locust-main-)
+-------+---------------------+-------+------------+-----------+------+
| NAME  | DEPLOYMENT          | READY | UP-TO-DATE | AVAILABLE | AGE  |
+-------+---------------------+-------+------------+-----------+------+
| hello | locust-main-hello   | 1/1   | 1          | 1         | 9m5s |
+-------+---------------------+-------+------------+-----------+------+
```

### klocust init
- Create config & locust files before applying. (ex name: hello)
```bash
$ klocust init hello
```

#### Update config & locust files
- Update config & locust files what you need. (ex name: hello)
```bash
$ vi hello-klocust.yaml
$ vi hello-locustfile.py
```

- If you want test locust in your local environments
```bash
$ docker run -p 8089:8089 -v $PWD:/mnt/locust locustio/locust -f /mnt/locust/hello-locustfile.py

or 

$ pip3 install locust
$ locust -f hello-locustfile.py
```

### klocust apply
- Create or Update locust cluster with config & locust files. (ex name: hello)
```bash
$ klocust apply hello
```

- **Connect to your locust cluster and do load testing.**  
(ex name: hello)  
```bash
$ open https://locust-hello.{your domain}
```

### klocust delete
- Delete locust cluster (ex name: hello)
```bash
$ klocust delete hello
```

## Contribution Guide
- Check [CONTRIBUTING.md](CONTRIBUTING.md) 
