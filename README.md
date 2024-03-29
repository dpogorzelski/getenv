## What's this?
`getenv` will fetch Google Cloud (Compute Engine) metadata attributes and print them as K=V pairs ready to be `exported` as environment variables.

Example:
```bash
getenv
MY_VARIABLE=SOMEVALUE
```

#### How to use

Export all metadata attributes as ENV variables for current shell session and store them in /etc/environment for future sessions:
```bash
export $(getenv)
getenv 1>> /etc/environment
```
or
```bash
getenv 1>> /etc/environment
source /etc/environment
```
