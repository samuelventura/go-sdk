# go-sdk

# tools

```bash
cd tools/socks5-netcat
go install
nc -l 31699
socks5-netcat -p 127.0.0.1:64658 -a 127.0.0.1:31699
```
