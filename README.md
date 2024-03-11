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

## Extra Credit (How i will do it)
To create a local mirror of an html page we need the assets to be stored locally 
and the static html file assets reference needs to be changed to the one in the local
So first of all what I will be doing is: 
```
For every urls given
    get all assets (css, script, images, iframe)
    download all the assets and put in in the separate folder with name format <url>_files
    redirect all the html element attributes to local assets
```
