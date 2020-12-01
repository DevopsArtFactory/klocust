# Contributing Guide
If you want to contribute to `klocust`, then please follow this guide.

## Process
1. Set up local development environment.
2. Change something...
3. Run test.
4. Create PR.

### 1. Set up local development environment
- You have to install...
    - golang ( >= 1.14 )
- You have to have and kubernetes cluster with kubernetes version 1.16 or higher

```bash
$ cd $GOPATH/src
$ mkdir -p github.com/DevopsArtFactory
$ cd github.com/DevopsArtFactory
$ git clone https://github.com/DevopsArtFactory/klocust.git
$ cd klocust
```

### 2. Change something
- Change codes
- If you create new function, then please **make unit test**.
- Please run `make fmt` in order to do formatting

### 3. Run test
- `make linters`: This will check the rules for clean code.
- `make test`: Run unit test

### 4. Create PR
- Thank you so much for your Pull Request!!
