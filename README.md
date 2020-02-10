##Terraform to supprt CI using Codebuild and ECR 

Run the follow command, this API open the 8080 Port:

```console
go run *.go
```

for building the docker image, run the follow commands:

```console 
docker build -t bookapi .
```

Test End points

http://localhost:8080/info
http://localhost:8080/books




