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

# Analizar los registros con Grafana .

### Configurar el Data Sources .

Inicie sesión en su instancia de Grafana y haga clic en el ícono de ajustes en la barra lateral izquierda. En la ventana emergente resultante, haga clic en Orígenes de datos(Data Sources). En la siguiente ventana, desplácese hacia abajo y seleccione MySQL de la lista. Luego se le presentarán las opciones de configuración necesarias para una conexión de datos MySQL.

```azure
host: ip_address_db:3306 
Database: squidlog
User: squidlog
Password: squidlog
```

### Importar los dashboard.

Para importar un dashboard, haga clic en Importar debajo del icono dashboard en el menú lateral.
Desde aquí puede cargar los archivos JSON que se encuentran en el directorio grafana.
En el paso 2 del proceso de importación, Grafana le permitirá cambiar el nombre del dashboard, elegir qué Data Sources desea que use el dashboard.
