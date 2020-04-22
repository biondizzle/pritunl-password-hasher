# pritunl-password-hasher

Made to be used in a bash script for automating Pritunl Installs

You can use it like this

```
#!/bin/sh
new_default_password="Kekeroni" 
new_default_password_hash=$(./pritunl-pass-hash "$new_default_password")
mongo --eval 'db.administrators.update({"username":"pritunl"},{$set:{"password":"'$new_default_password_hash'"}});' pritunl
```
