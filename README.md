##Terraform to supprt CI using Codebuild and ECR 

Run the follow command, this API open the 8080 Port:

```console
go run *.go
```

for building the docker image, run the follow commands:

```console 
docker build -t libraryapi .
```

***Make the Codebuild and ECR:***

add the needed variables to terraform.tfvars and run *terrafom plan* and *terraform apply*


