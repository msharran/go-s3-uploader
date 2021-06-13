# Go S3 log uploader

To upload the log files to your s3 bucket, configure your aws cli and add the credentials in `~/.aws/credentials file`

Build binary
```bash
go build
```

run the cli using the following command
```bash 
./s3_uploader \
 -d=/usr/local/var/log/nginx/ \
 -age=5
 ```
