# Web Fetcher
CLI Application that can be used to fetch web into the current directory. 

## How To Use (Go Compiler)
- Install golang
- Run the code below
```
go mod download
go build -o fetch main.go
./fetch <url1> <url2>
./fetch --metadata <url1>
```

## How To Use (Docker)
- Install docker
- Run the code below
```
docker build -t fetcher .
docker run --rm -it -v $(pwd):/docs -e DOCKER_MOUNTED_VOL='docs' fetcher <url1> <url2>
docker run --rm -it -v $(pwd):/docs -e DOCKER_MOUNTED_VOL='docs' fetcher --metadata <url1>   
```


