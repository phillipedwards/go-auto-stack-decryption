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


With all references of pulumi/pulumi set to 3.