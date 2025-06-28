Please select "Code" view before reading this document

Prerequest
git clone git@github.com:karmegamp/cyderes.git

A. We can execute the application via container by following below steps

1. Build the application container
==================================
$docker image build . -t datatx:latest

2. Start the application container
==================================
$docker container run -p 7777:8888 datatx:latest

3. Start data fetch service for transformation
==============================================
POST   http://0.0.0.0:7777/datatx?cmd=start
Response 
201 Created
Successfully collected data in cloud_store.txt

4. Interrogate container 
========================
# ls
cloud_store.txt  go.mod  main.go
# ls -l
total 44
-rw-r--r-- 1 root root 33018 Jun 26 21:33 cloud_store.txt
-rw-r--r-- 1 root root    25 Jun 14 18:34 go.mod
-rw-r--r-- 1 root root  1889 Jun 26 21:09 main.go



B. Alternatively, 

Executing below command will fetch the data, transform and generate "cloud_store.txt" output file.
$ go run main.go
POST http://0.0.0.0:8888/datatx?cmd=start
Response 
201 Created
Successfully collected data in cloud_store.txt


C. Kubernetes deployment details

Create deployment
$ kubectl create deployment datatx --image=docker.io/karmegamp/datatx:v1

Create service
$ kubectl expose deployment datatx --type=NodePort --port=8888

Create tunnel local port for service testing
$ DATATX_POD=$(kubectl get pods -l app=datatx -o jsonpath='{.items[0].metadata.
name}')
$ kubectl port-forward $DATATX_POD 52951:8888

POST http:///127.0.0.1:52951/datatx?cmd=start
Response 
201 Created
Successfully collected data in cloud_store.txt