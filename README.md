# copySquidLogToMysql
Copiar registros del squid para una base de datos en mysql

### Otorgar permisos al binario
`> chmod +x ./squidCopyLogMysql`
### Ejecutar el binario 
Para que genere el fichero de configuracion (config)

`> ./squidCopyLogMysql`
```azure
-- System configuration file   --
-- Created by Edgar Javier akosej9208@gmail.com  --
---------------------------------------------------------------------------
#-- Path necessary files
path.AccessLog=./access.log
---------------------------------------------------------------------------
#--Mysql server Ipaddress
mysql.ip=127.0.0.1
#--Mysql username
mysql.user=usuario
#--Mysql Password
mysql.pass=secreto
#--Mysql server Port
mysql.port=3306
#--Mysql database
mysql.db=nameDB
```
### Edite el fichero de configuracion
`> nano ./config`

### Comience a copiar los registros en la base de datos
El binario se queda corriendo percistentemente por lo que se le recomienda que cree un servicio en systemd para su uso definitivo

`nano /etc/systemd/system/squidCopyLogMysql.service`

```azure
[Unit]
Description=squidCopyLogMysql

[Service]
WorkingDirectory=/opt
ExecStart=/opt/squidCopyLogMysql
CPUAccounting=true
#CPUQuota=40%

[Install]
WantedBy=multi-user.target
```

`systemctl enable squidCopyLogMysql.service`

`systemctl daemon-reload`

`systemctl start squidCopyLogMysql.service`

# Listo
Si todo esta bien ya tiene los registro de su proxy en una base de datos en MYSQL
