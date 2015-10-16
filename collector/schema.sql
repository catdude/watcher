CREATE TABLE measurement (      -- Description of measurements to be performed
    measID INTEGER PRIMAY KEY,
    name VARCHAR(32) NOT NULL,  -- Name of this measurement
    class INTEGER NOT NULL,     -- Foreign key link to class.id
    db VARCHAR(64) NOT NULL,    -- Name of Influxdb database 
    desc VARCHAR(64),           -- Text description of the measurements
    period INTEGER              -- How often should test tun (minutes)
);

CREATE TABLE class (            -- Device class (i.e. DNCC, CDS, Linux, CiscoSwitch, etc)
    classID INTEGER PRIMARY KEY,
    name VARCHAR(32) NOT NULL
);

CREATE TABLE collector (        -- All hosts that will be collecting data from devices for Watcher
    collID INTEGER PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    ip VARCHAR(16) NOT NULL
);

CREATE TABLE site (             -- Locations where devices are clustered (NLV, Detroit, EU, Lux, etc)
    siteID INTEGER PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    thru INTEGER,               -- If connected through VPN, through which site does VPN run
    collector INTEGER           -- Foreign key link to collector.id
);

CREATE TABLE host (
    hostID INTEGER PRIMARY KEY,
    name CHAR(64),
    ip CHAR1(16),
    class INTEGER,              -- Foreign  key link to class.classID (DNCC, CDS, etc)
    site INTEGER,               -- Foreign key link to site.id
    status INTEGER,             -- 0 = inactive, 1 = active
    type INTEGER,               -- 0 = ping only, 1 = SNMP
    community CHAR(32)          -- SNMP community string, if applicable
);

CREATE TABLE test_master (      --  Single-line OIDs and the top (table) entry for table-based OIDs
    masterID INTEGER PRIMARY KEY,
    oid CHAR(128),              -- Numeric OID for single-line and top-of-table (entry) OIDs
    name CHAR(128),             -- Text name (optional)
    mib CHAR(64),               -- Name of MIB defining OID (optional)
    class INTEGER,              -- Might be deleted later - class of item being read
    measurement INTEGER         -- Foreign key to measurement.measID
);

CREATE TABLE test_item (        -- single-line OIDs and line-item OIDs from above table-based OIDs
    itemID INTEGER PRIMARY KEY,
    oid CHAR(128),
    measurement INTEGER,        -- Foreign key link to measurement.measID
    series CHAR(64)             -- Name of InfluxDB series to store values
);

CREATE TABLE test_parse (       -- Indicates how many levels deep indexes go for table-based OIDS
    parseID INTEGER PRIMARY KEY,
    oid INTEGER,                -- Foreign Key link to test_item.id
    levels INTEGER,             -- Number of levels indexes are nested
    measurement INTEGER         -- Foreign key link to measurement.measID
);

