# go-ntpd-server-statistics
Obtains the amount of UNIQUE public IPs that access our self-hosted NTPD server during a interval of seconds.


Requires install LibPcap development libs:

```
apt install libpcap-dev
```

Modify config.yml to match your interface and public ip to listen. 

Just run the thing:

```
go run main.go
````

