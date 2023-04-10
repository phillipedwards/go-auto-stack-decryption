# Pulumi Configuration Decrypt Repro

To reproduce the issue locally, run after setting all references of pulumi/pulumi to 3.57 or greater:
```
go run main.go
```

To reproduce the issue in a Docker container, run:
```
./docker-repro.sh
```

With all references of pulumi/pulumi set to 3.57 or greater you should expect to see the following output:
```
...
go: downloading gopkg.in/warnings.v0 v0.1.2
go: downloading github.com/cloudflare/circl v1.1.0
go: downloading golang.org/x/text v0.7.0
starting run...
pulumi files removed...
2023/04/10 16:27:01 Pulumi version: 3.57.0
2023/04/10 16:27:05 GetAllConfig: unable to read config: exit status 255
code: 255
stdout: 
stderr: error: could not decrypt configuration value: cipher: message authentication failed
```



With all references of pulumi/pulumi set to 3.57 or less the program works as expected.