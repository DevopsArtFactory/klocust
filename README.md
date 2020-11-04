# klocust
Klocust is a command-line tool for managing [Locust](https://locust.io/) distributed load testing on [Kubernetes](https://kubernetes.io/).  

## Installation
### Install klocust binary (TODO)
```bash
$ brew tap DevopsArtFactory/klocust
$ brew update
$ brew install klocust 
``` 

### Build from source
See here on how to [build klocust CLI from source][build from source].

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

>>> 1 locust deployments in loadtest namespace. (PREFIX: locust-master-)
+-------+---------------------+-------+------------+-----------+------+
| NAME  | DEPLOYMENT          | READY | UP-TO-DATE | AVAILABLE | AGE  |
+-------+---------------------+-------+------------+-----------+------+
| hello | locust-master-hello | 1/1   | 1          | 1         | 9m5s |
+-------+---------------------+-------+------------+-----------+------+
```

### klocust init
- Create config & locust files before applying. (ex name: hello)
```bash
$ klocust init hello
```

#### Update config & locust files
* Update config & locust files what you need. (ex name: hello)
```bash
$ vi hello-klocust.yaml
$ vi hello-locustfile.py
```

* If you want test locust in your local environments
```bash
$ docker run -p 8089:8089 -v $PWD:/mnt/locust locustio/locust -f /mnt/locust/hello-locustfile.py
or 
$ pip3 install locust
$ locust -f hello-locustfile.py
```

### klocust apply (TODO)
- Create or Update locust cluster with config & locust files. (ex name: hello)
```bash
$ klocust apply hello
or
$ klocust apply -cf hello-klocust.yaml -lf hello-locustfile.yaml
```

### klocust delete (TODO)
- Delete locust cluster (ex name: hello)
```bash
$ klocust delete hello
```



[build from source]: ./docs/build_from_source.md
