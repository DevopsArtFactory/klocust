# Build klocust CLI from Source

1. Verify that you have Go 1.13+ installed.
    ```bash
    $ go version
    ```
    If `go` is not installed, follow instructions on [the Go website](https://golang.org/doc/install).

2. Clone this repository
    ```bash
    $ git clone https://github.com/DevopsArtFactory/klocust/klocust.git 
    $ cd klocust
    ```

3. Build the project
    ```bash
    $ make
    ```

4. Move the resulting `out/klocust` executable to somewhere in your PATH
    ```bash
    $ sudo mv ./out/klocust /usr/local/bin/
    ```

5. Run `klocust version` to check if it worked. 
