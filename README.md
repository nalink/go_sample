# go_sample
Go Sample 

Simple Go based service

Define GOPATH

Copy all src/sample.com GOPATH/src/

* Local

    > go run server.go
* Docker

    > docker build -t go_sample:latest .
    
    > docker run -p 3000:3000 go_sample:latest

* Kubernetes
  
    > kubectl run gosample --image=go_sample --port=3000

    > kubectl expose deployment gosample --type=LoadBalancer
    

