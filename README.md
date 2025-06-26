Please select "Code" view before reading this document

Executing below command will fetch the data, tranform and generate "cloud_store.txt" output file.
$ go run main.go
POST http://0.0.0.0:8888/datatx?cmd=start
Response 
201 Created
Successfully collected data in cloud_store.txt


Alternatively, we can execute the application via container by following below steps

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

4. Introgate container 
======================
# ls
cloud_store.txt  go.mod  main.go
# ls -l
total 44
-rw-r--r-- 1 root root 33018 Jun 26 21:33 cloud_store.txt
-rw-r--r-- 1 root root    25 Jun 14 18:34 go.mod
-rw-r--r-- 1 root root  1889 Jun 26 21:09 main.go

