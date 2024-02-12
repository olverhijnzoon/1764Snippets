PRAGMA table_info(ALARMJOURNALS);
INSERT INTO ALARMJOURNALS VALUES(1,'AlarmJournalInternal1764','LOCAL',1,NULL,0);

/* local_journal_dump.sql
PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE alarm_events (id integer ,eventid text,source text,displaypath text,priority integer,eventtype integer,eventflags integer,eventtime numeric,PRIMARY KEY (id));
INSERT INTO alarm_events VALUES(1,'db5cb052-cfe7-4e7c-93aa-7e87a5f55bf2','prov:default:/tag:Snippet1764AlarmTag:/alm:Snippet1764Alarm','',1,1,16,1707733391555);
INSERT INTO alarm_events VALUES(2,'6c4e6364-ba3a-4cc4-a6a7-13651edffafc','prov:default:/tag:Snippet1764AlarmTag:/alm:Snippet1764Alarm','',4,0,0,1707733395229);
INSERT INTO alarm_events VALUES(3,'6c4e6364-ba3a-4cc4-a6a7-13651edffafc','prov:default:/tag:Snippet1764AlarmTag:/alm:Snippet1764Alarm','',4,1,16,1707733422785);
INSERT INTO alarm_events VALUES(4,'80651bb2-a74a-4796-a1c1-5c988de9e253','prov:default:/tag:Snippet1764AlarmTag:/alm:Snippet1764Alarm','',4,0,0,1707733444295);
CREATE TABLE alarm_event_data (id integer,propname text,dtype integer,intvalue integer,floatvalue real,strvalue text);
INSERT INTO alarm_event_data VALUES(1,'eventValue',0,999,NULL,NULL);
INSERT INTO alarm_event_data VALUES(2,'eventValue',0,1764,NULL,NULL);
INSERT INTO alarm_event_data VALUES(3,'eventValue',0,2,NULL,NULL);
INSERT INTO alarm_event_data VALUES(4,'eventValue',0,2000,NULL,NULL);
CREATE INDEX alarm_event_dataidndx ON alarm_event_data(id);
COMMIT;
*/

-- some further notes

-- INSERT INTO TAGCONFIG VALUES('f5d958ef-33c4-4aaa-a20d-7d0044181384',0,NULL,replace('{\n  "valueSource": "memory",\n  "alarms": [\n    {\n      "name": "Alarm1764",\n      "priority": "High"\n    }\n  ],\n  "name": "Tag1764",\n  "value": 1764,\n  "tagType": "AtomicTag"\n}','\n',char(10)),1,'Tag1764');

-- PRAGMA table_info(TAGCONFIG);

/* 
UPDATE TAGCONFIG
SET CFG = json_set(
    CFG,
    '$.alarms[0].priority', 'Critical'
)
WHERE ID = 'f5d958ef-33c4-4aaa-a20d-7d0044181384';
*/