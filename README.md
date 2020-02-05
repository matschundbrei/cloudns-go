# ClouDNS-Go

This is an API-Client for the [ClouDNS HTTP API](https://www.cloudns.net/wiki/article/42/) written in [Go](https://golang.org)

## Noob code warning

Currently this software is  written by an [absolute noob in go](https://github.com/matschundbrei). 
Therefore, I am very happy about any pointers toward making the code simpler, faster or more lightweight.

Use at your own risk!

## Usage

There are three structs that you need to know:

 * **Apiaccess**: Holds your authentication parameters (auth-id/sub-auth-id, auth-password)
 * **Zone**: Holds information about a zone
 * **Record**: Holds information about a record
 
These structs have methods, that call the API, most of them return either an Array of the other ones or the updated input struct and an error.
 
 ### Apiaccess Methods
 ### Zone Methods
 ### Record Methods
 


## Limitations

I did not yet touched any of the advanced features, that ClouDNS offers. For details please check the [limitations](limitations.md)
