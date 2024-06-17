# Primary
`SHOW data_directory;` to find $PGDATA
On 11: `/var/lib/postgresql/11/main`
`vi /etc/postgresql/11/main/pg_hba.conf`
add entry
`host    replication     repmgr          172.31.43.17/32           md5`

`create role repmgr with replication login encrypted password 'repmgr123';`

# Stop standy service
`existing process: /usr/lib/postgresql/11/bin/postgres -D /var/lib/postgresql/11/main -c config_file=/etc/postgresql/11/main/postgresql.conf`
to stop it: `/usr/lib/postgresql/11/bin/pg_ctl stop -D /var/lib/postgresql/11/main`
# Backup from Standby
`pg_basebackup -D /var/lib/postgresql/11/backup \
                         -X stream \
                         --write-recovery-conf \
                         --slot=standby1 \
                         --create-slot \
                         --dbname="host=172.31.34.144 dbname=alfreddb user=repmgr"`
`pg_basebackup -Xs -d "hostaddr=172.31.34.144 port=5432 user=repmgr password=repmgr123" -D /var/lib/postgresql/11/main -v -Fp`
`cp /usr/share/postgresql/11/recovery.conf.sample /var/lib/postgresql/11/main/recovery.conf`
```
recovery_target_timeline = 'latest'
standby_mode = on
primary_conninfo = 'host= 172.31.34.144 port=5432 user=repmgr password=repmgr123'
primary_slot_name = 'postgresql_node102'
trigger_file = 'tgfile'
```
select * from pg_create_physical_replication_slot('postgresql_node101');
select * from pg_create_physical_replication_slot('postgresql_node102');



add to pg_hba.conf on standby
```
host all all 172.31.34.144/32 md5
host replication all 172.31.34.144/32 md5
```


# Start standby
`chmod 0700 -R /var/lib/postgresql/11/main/`
`/usr/lib/postgresql/11/bin/pg_ctl start -D /etc/postgresql/11/main`
