go version go1.17.5 


# Run application


### Compile app 
- instal `golang` 
- run `go run cmd/main.go` 
- `go build cmd/main.go` 

 ### Build Binary 
- `go build cmd/main.go` 

### Run binary build 
- `./main` ※ *no need to install golang* 
- If you got `bash: ./main: permission denied` error 
- Run `chmod +x main` ※ *Permit the app if you got binary build from other PC* 


Run app with logger `go run cmd/main.go --logger` or `./main --logger`

### Permite serial USB in jetson 

1. check group by groups  `sudo adduser $USER dialout`

2. reboot device

3. check group again `sudo groups $USER`

4. then run `chmod +x main` 


### Register auto run use `pm2` package

- install `nodejs`, `npm` and `pm2` 

```
pm2 list 
pm2 start ./main
pm2 stop main
pm2 delete main
pm2 save --force
```