# watcher
Multitool integration to provide system monitoring and alerting, ala EM7 or Nagios

There are a lot of really nice system monitoring tools out there, 
both open source and commercial. There is the ubiquitous nagios, which
works nicely but is a pain to initially configure, is a resource hog, 
and looks pretty unattractive. Then we've got the Python implementation
of the nagios API, Shinken. It's nice that it's in python, but has the
same problems as nagios.

Zabbix and Zenoss are possibilities, but they have never really impressed
me. 

We have newer tools, like sensu (pretty nicely done) and prometheus (VERY
nicely done) to consider.

And finally we have commercial tools like EM7 (from Science Logic) and some
others.

Choosing the right tool for your organization's needs often involves just
playing around the available tools, seeing if they do what you want, then 
calling around to commercial vendors if they don't.

I got tired of all that. We use EM7 where I work, and it's very powerful,
easy to configure, mostly easy to maintain, and mostly reliable. Unfortunately, 
the few times that we have had outages have been at very inopportune times.
This had lead me to want to develop a standby or backup monitoring solution
that we can depend on during those few times when our primary monitor goes 
out. That was what lead to the "Watcher" project.

A good monitoring system needs a good time-series database. Prometheus has
a very nice database, but it isn't really well suited (by their description)
to maintaining archives of detailed point data. Since my need is heavy into
analyzing old point data, something else was needed. That is where InfluxDB
comes in (http://influxdb.com). InfluxDB is open source, written in Go so it
performs pretty well, and is expressly designed to allow access to detailed
point data.

Graphing needs can be handled by a package such as Grafana (http://grafana.org)
or by InfluxDB's tool, Chronograf. Tests using this project will help to determine
which gets used.

This project will be taking place kind of slowly, since developing this is 
not my primary job.

If anyone is interested in participating, send me an email (catdude@gmail.com).

