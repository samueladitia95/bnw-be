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

# How to copy file from local to remote instance

1. run command `scp -i /path/to/private-key.pem -r /path/to/origin-file username@public-ip-address:/path/to/target-file`

# How to copy file from remote instance to local

1. run command `scp -i /path/to/private-key.pem -r username@public-ip-address:/path/to/origin-file  /path/to/target-file`
