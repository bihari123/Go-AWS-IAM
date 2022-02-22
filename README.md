# Go-AWS-IAM
Multi-purpose tool built for Go  AWS SDK 

# Setting the environment variable
```
$ export AWS_ACCESS_KEY_ID=<<Your-api-access-key>>
$ export AWS_SECRET_ACCESS_KEY=<<your-secret-access-key>>
```

# Running the operations
```
$ go run aws_IAM.go createUser <username>
$ go run aws_IAM.go listAccessKey <username>
$ go run aws_IAM.go listUsers
$ go run aws_IAM.go updateUserName <currentName> <newName>
$ go run aws_IAM.go deleteAccessKeyID <userName> <accessKeyId>

```
