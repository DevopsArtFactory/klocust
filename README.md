# klocust
A command-line tool for managing [Locust](https://locust.io/) distributed load testing on [Kubernetes](https://kubernetes.io/)

# Install klocust

# klocust list
```bash
$ klocust list

>>> 1 locust deployments in loadtest namespace. (PREFIX: locust-master-)
+-------+---------------------+-------+------------+-----------+------+
| NAME  | DEPLOYMENT          | READY | UP-TO-DATE | AVAILABLE | AGE  |
+-------+---------------------+-------+------------+-----------+------+
| hello | locust-master-hello | 1/1   | 1          | 1         | 9m5s |
+-------+---------------------+-------+------------+-----------+------+
```

# build klocust binary
```bash
$ make build
```
