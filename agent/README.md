# Topogopher

This project comes with a client and a server with the purpose to measure netword distance between at least two machines.

For bandwith and efficiency sake, hosts only try a fraction of the list in a triangle pattern.

Consider a list of four hosts host1..4. The calls between hosts would take place as follows :

host1 -> host2, host3, host4

host2 -> host3, host4

host3 -> host4.

The client invocation would take the form :

`./main -start=client host1 host2 host3 host4`

On each host, a server must run to respond to sollicitations. This is done by :

`./main -start=server`

A port is optionally settable by :

`./main -start=server -port=xxxx`
Same with the client to address the port:

`./main -start=client -port=xxxx host1 host2 host3 host4`

Note that the port must be the same on each host.

Of course, you can reuse the code in your project without compiling the main executable.