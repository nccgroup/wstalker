# WStalker: an easy proxy 

WStalker is a HTTP/HTTPS Proxy with full Request/Response logging. The main purpose of WStalker is to provide an easy way for developers to configure their testing tools to use this proxy and register the valid connections they use to test Web Services. The resulting CSV file can be used in security testing, by importing the requests/responses into your favourite testing tool.

## License

Released as open source by NCC Group Plc - http://www.nccgroup.com/

Developed by Jose Selvi, jose dot selvi at nccgroup dot com

http://www.github.com/nccgroup/wstalker

Released under AGPL see [LICENSE](LICENSE.md) for more information

## Compile or Download WStalker

To use it, you can compile your own binaries from the source code by running `./build.sh`.
You can also download a set of binaries that were already compiled for you [from here](https://github.com/nccgroup/wstalker/releases/).

## Running WStalker

Once you have the binary, it can be executed without any additional parameter. In Windows, it can be executed by double clicking the EXE file.

```
$ ./wstalker
2019/10/18 09:19:30 Creating HTTP Proxy
2019/10/18 09:19:30 Starting in 127.0.0.1:8080
2019/10/18 09:19:31 Saving Request in wstalker.csv
2019/10/18 09:19:31 Stalking Connections...
```

Now, it is necessary to configure the tool that we are using (browser or any other tool that we are using to test your web services) to use http://127.0.0.1:8080 as a proxy. When we start running requests, wstalker will start showing one line per request.

```
2019/10/18 09:20:53 GET - http://ifconfig.co/
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://ifconfig.co/favicon.ico
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
2019/10/18 09:20:54 GET - http://detectportal.firefox.com/success.txt
```

When all the tests have been performed, wstalker can be stopped by pressing CTRL+C or closing the window.

```
2019/10/18 09:21:01 Closing wstalker
```

## Output of WStalker

A file `wstalker.csv` will be generated, containing information about each request: `[REQUEST_IN_BASE64],[RESPONSE_IN_BASE64],METHOD,URL`.
Each new execution of `wstalker` appends the new requests and responses to the existing file.

This file could contain credentials and critical information, so it is recommended to send it encrypted or using a secure mechanisms.

## Do not install wstalker's CA

Note that wstalker is using a CA to generate certificates but, for simplicity purposes, that CA's private key is included in the code, so *PLEASE DO NOT INSTALL IT IN YOUR BROWSER OR OPERATING SYSTEM*, since that would mean that attackers could intercept your connections by generating certificates signed by wstalker's CA.
