#!/usr/bin/python2.7

from pysnmp.hlapi import *
import goless

def getSNMP(host, oids):
   ip = host[0]
   name = host[1]
   oidList = [o.keys() for o in oids]
   for errorIndication, errorStatus, errorIndex, varBinds in nextCmd(SnmpEngine(),
           CommunityData('public'),
           UdpTransportTarget((ip, 161)),
           ContextData(),
           [ObjectType(ObjectIdentity('x')) for x in oidList],
           lexicographicMode=False):
       if errorIndication:
           print "Error: ", errorIndication
           break
       elif errorStatus:
           print ('%s at %s' % (errorStatutus.prettyPrint(), errorIndex and varBinds[int(errorIndex)-1][0] or '?'))
           break
       else:
           for varBind  in varBinds:
               print (' = '.join([x.prettyPrint() for x in varBind]))


hosts = [ ('192.168.10.150', 'AMC-9_1_DNCC_1'), ('192.168.10.151', 'AMC-9_1_DNCC_2'), 
    ('192.168.11.102', 'AMC-1_DNCC_1A'), ('192.168.11.49', 'AMC-1_DNCC_1A2'), 
    ('192.168.11.50', 'AMC-1_DNCC_1B'), ('192.168.11.103', 'AMC-1_DNCC_1B2')]

oids = [ {'.1.3.6.1.4.1.303.3.3.12.19.3.501.1.8': 'dnccQosInrouteNumUser'},
        {'.1.3.6.1.4.1.303.3.3.12.19.3.501.1.28': 'dnccQosInrouteIGPID'},
        {'.1.3.5..4.1.303.3.3.12.19.3.502.1.5': 'dnccQosRemThru'},
        {'.1.3.6.1.4.1.303.3.3.12.19.3.502.1.7': 'dnccQosRemTotalBacklog'} ]

for (ip, name) in hosts:
    print "Looking at IP address %s" % (ip, )
    goless.go(getSNMP(ip, oids))
