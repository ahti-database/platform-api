# Ahti

Ahti is an open-source database-as-a-service (DBaaS) that you can self-host in your organization or home network. It functions similar to [Turso](https://docs.turso.tech), where you can declaratively provision a database for your application.

It uses [libSQL](https://github.com/tursodatabase/libsql) as the database of choice as it is able to accept database connections via HTTP, thus lends itself well to deployment via ingresses without having to setup open TCP connections in k8s.
