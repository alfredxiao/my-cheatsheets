```
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo -u postgres psql -c "SELECT version();"

vi /etc/postgresql/11/main/postgresql.conf
add/modify listen_addresses = '*'

vi /etc/postgresql/11/main/pg_hba.conf
append two lines at the end:
host    all             all              0.0.0.0/0                       md5
host    all             all              ::/0                            md5

sudo systemctl restart postgresql

sudo su - postgres -c "createuser -P alfred"  # -P means prompt for setting password
sudo su - postgres -c "createdb alfreddb"
sudo -u postgres psql
  GRANT ALL PRIVILEGES ON DATABASE alfreddb TO alfred;
  \c alfreddb
  GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO alfred;

# From another host
psql --host=<DB_HOST_OR_IP> -U alfred --dbname=alfreddb
```
