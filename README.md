# Notes-App
Project using Go + Nuxt.js

## Install
Make sure Go is already installed on your PC.

Clone this repository and install all required dependencies.

## Setup
### Environment variables
Create `config.env` file and setup the following variables:
```
db_user = *your_username*
db_pass = *your_password*
db_name = *your_database_name*
db_host = localhost
db_port = *postgres_server_port*
````

## Start
Type the following code in the terminal to run the API server:
```
$ go run server.go
```
Then type the following to run the development server:
```
$ cd client
$ npm run dev

```

## Instruction to Deploy
Notes-App works with postgresql, nuxt, golang.  
Our service checked in deb distros + apache2.
### Git Clone
First, let's clone Notes-App
``` bash
git clone https://github.com/on3dd/Notes-App.git
```
### Database
Let's install the Database
``` bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```
Login to postgresql
``` bash
sudo -u postgres psql
```
Create Database & User  
``` sql
CREATE DATABASE yourdbname;
CREATE USER youruser WITH ENCRYPTED PASSWORD 'yourpass';
GRANT ALL PRIVILEGES ON DATABASE yourdbname TO youruser;
\q
```
Import tables
``` sql
CREATE TABLE "messages" (
	"id" serial NOT NULL,
	"text" TEXT NOT NULL,
	"category_id" int NOT NULL,
	"posted_at" TIMESTAMP NOT NULL,
	"author_id" int NOT NULL,
	CONSTRAINT "messages_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);
CREATE TABLE "categories" (
	"id" int NOT NULL,
	"name" varchar(255) NOT NULL,
	"parent_id" int,
	CONSTRAINT "categories_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);
CREATE TABLE "users" (
	"id" int NOT NULL,
	"name" varchar(255) NOT NULL,
	CONSTRAINT "users_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);
```
Create config.env in Notes-App/
``` bash
cd Notes-App/
echo "db_user = *your_username*" >> config.env
echo "db_pass = *your_password*" >> config.env
echo "db_name = *your_database_name*" >> config.env
echo "db_host = localhost" >> config.env
echo "db_port = *postgres_server_port*" >> config.env
```
### Install Node JS & npm
Please install latest version from PPA
``` bash
curl -sL https://deb.nodesource.com/setup_8.x -o nodesource_setup.sh
sudo bash nodesource_setup.sh
sudo apt install nodejs
rm nodesource_setup.sh
```
Check versions
``` bash
nodejs -v
npm -v
```
App checked in node js v8.16.2 & npm v6.4.1  
  Install nodejs dependencies
``` bash
npm install Notes-App/client
```
### Golang
This will give you latest version of go
``` bash
sudo snap install --classic go
```
You have to move Notes-App/ to $GOROOT
``` bash
cd .. && mv Notes-App/ $GOROOT/src/Notes-App/
```
Check version
``` bash
go version
```
App checked in go v1.13.4
#### Install Goland dependencies
``` bash
cd Notes-App/api
go install
cd .. && go install
```
### Install & Setup Apache2
Install Apache2
``` bash
sudo apt install apache2
```
Touch virtual host
``` bash
sudo vi /etc/apache2/sites-aviable/your_host.conf
```
And describe host
``` apache2
<VirtualHost *:80>
        ServerAdmin ServerAdmin@host.ru
        ServerName your_hostname
        ServerAlias www.your_hostname
        ProxyPreserveHost On
        ProxyPass / http://127.0.0.1:3000/
        ProxyPassReverse / http://127.0.0.1:3000/
</VirtualHost>
```
Lauch your virtual host
``` bash
sudo a2ensite your_host.conf
sudo service apache2 restart
```
### Firewall
Let your Firewall access Apache2
``` bash
sudo ufw allow 'Apache'
```
### Make systemd service
Sure, we can run our App how service. Just create new service
``` bash
sudo vi /etc/systemd/system/Notes-App-server.service
```
And describe service
``` systemd
[Unit]
Description=Notes-App-server
After=network-online.target

[Service]
Restart=on-failure
WorkingDirectory=$GOROOT/src/Notes-App/
ExecStart=/snap/bin/go run server.go
User=USER
[Install]
WantedBy=multi-user.target
```
Where User - who a you. Let repeat this for server.  
Touch Notes-App-client.service  
``` bash
sudo vi /etc/systemd/system/Notes-App-client.service
```
And describe service
``` systemd
[Unit]
Description=Notes-App-client
After=network-online.target

[Service]
Restart=on-failure
WorkingDirectory=$GOROOT/src/Notes-App/client/
ExecStart=/usr/bin/node run dev
User=USER
[Install]
WantedBy=multi-user.target
```
Reload systemd daemon and enable ours services
``` bash
systemctl daemon-reload
systemctl enable Notes-App-client.service
systemctl enale Notes-App-server.service
```
Well, we can see our job in youhostname
