# download-db-apn

Download all available apn sketches (gleisplaene) that are available on [https://trassenfinder.de](https://trassenfinder.de).

## How to install

Download the binary you need from the latest release and execute it. You don't need to pass any cli arguments.

### Cli arguments
```
usage: download-db-apn [-h|--help] [-i|--infrastructure-id <integer>]
                       [-t|--target-directory "<value>"]

                       Download all available apn sketches from
                       trassenfinder.de

Arguments:

  -h  --help               Print help information
  -i  --infrastructure-id  Change the ID of the infrastructure, but I don't
                           think you'll need this. Default: 12
  -t  --target-directory   The directory where the apk sketches will be saved
                           in. Default: target
```