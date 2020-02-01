# Limitations

The whole idea behind this module is to enable easy creation of zones and records through the api.
The ClouDNS API however covers a lot more features than that. If you are looking for something that
can fetch stats, create very specific zones and records using the advanced features of ClouDNS, like
HTTP/email forwarding and Geo-Loadbalancing, I have to tell you that this is currently not covered by this Module.

Further limitations are:

- Updating zones is currently not possible
- Accounts with more than 100 zones will only list up to 100 with the Listzones() function
- ...
