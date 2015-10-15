CREATE TABLE measurement (  -- Description of measurements to be performed
    id INTEGER PRIMAY KEY,
    name V  name VARCHAR(32) NOT NULL,
    class INTEGER NOT NULL,
    db VARCHAR(64) NOT NULL,
    desc VARCHAR(64),
    period INTEGER
);
CREATE TABLE class (        -- Device class (i.e. DNCC, CDS, Linux, CiscoSwitch, etc)
    id INTEGER PRIMARY KEY,
    name VARCHAR(32) NOT NULL
);
CREATE TABLE collector (    -- All hosts that will be collecting data from devices for Watcher
    id INTEGER PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    ip VARCHAR(16) NOT NULL
);
CREATE TABLE site (     -- Locations where devices are clustered (NLV, Detroit, EU, Lux, etc)
    id INTEGER PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    thru INTEGER,       -- If connected through VPN, through which site does VPN run
    collector INTEGER
);
CREATE TABLE oid  (
id INTEGER PRIMARY KEY,
measurement INTEGER,        -- With which measurement is this OID associated?
oid VARCHAR(128),
indexOid INTEGER,           -- If this OID is indexed, what is the numeric OID to the index?
top INTEGER,                -- Is this the top of a subtree? If so, how many octets are index values?
name VARCHAR(128),
class INTEGER,              -- Which device class does this apply to? Default is ALL.
mib VARCHAR(32),
desc TEXT,
type VARCHAR(32),
scanner INTEGER             -- Which data scanner will evaluate the return and stuff into Influxdb?
);
CREATE TABLE hosts (
id INTEGER PRIMARY KEY,
name CHAR(64),
ip CHAR1(16),
class INTEGER,     -- Foreign  key link to "class" table (DNCC, CDS, etc)
site INTEGER,
status INTEGER,      -- 0 = inactive, 1 = active
type INTEGER,       -- 0 = ping only, 1 = SNMP
community CHAR(32) -- SNMP community string, if applicable
);
