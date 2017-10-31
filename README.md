# gosparrow

A minimal toy go api server.

## Quick Start

    mkdir gosparrow_dir && export GOPATH=`pwd`
    mkdir src && cd $_
    git clone https://github.com/lylex/gosparrow.git
    cd gosparrow
    glide install
    cd cmd/gosparrow
    go build

## Architecture & Design

Are you kidding? Does this poor toy project need a architecture?

## Contributing

Learning togerher is greate thing. But I do not think you are intrested in such a humble project. If you wish, change anything you want.

## FAQ

**Q: Why need this project?**

It does nothing but a toy.


**Q: How to generate the proto file**

```sh
cd pkg/gosparrow
protoc --go_out=plugins=grpc:. *.proto
```
**Q: How to update the libraries and add them the glide.yml**

```sh
glide get github.com/Masterminds/cookoo
```
