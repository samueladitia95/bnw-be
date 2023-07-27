# How to run development

`go run main.go serve --http=localhost:8080`

# How to build

`go build`

# How to run built app

`./bnw-be serve --http=localhost:8080`

# How to connect to remote server lightsail

1. download public ssh key of your account
2. run command `sudo chmod 400 /path/to/private-key.pem`
3. run command `ssh -i /path/to/private-key.pem username@public-ip-address`
