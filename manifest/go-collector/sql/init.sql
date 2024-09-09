CREATE USER 'speed_report'@'127.0.0.1' IDENTIFIED BY 'IUY*&^*^!2HGJHG886!32';

GRANT ALL PRIVILEGES ON *.* TO 'speed_report'@'127.0.0.1';

FLUSH PRIVILEGES;


CREATE USER 'speed_report'@'localhost' IDENTIFIED BY 'IUY*&^*^!2HGJHG886!32';
GRANT ALL PRIVILEGES ON *.* TO 'speed_report'@'localhost';


CREATE USER 'user_001'@'127.0.0.1' IDENTIFIED BY 'IUY*&^*^!2HGJHG886!32';

GRANT ALL PRIVILEGES ON *.* TO 'user_001'@'127.0.0.1';

FLUSH PRIVILEGES;



CREATE USER 'speed_backup'@'%' IDENTIFIED BY 'bakIUY*&^*^!12H6!326oihjh*(78712YH129-,IUTCJGFZA6761HGqw[ooooPPPP';

GRANT ALL PRIVILEGES ON speed.* TO 'speed_backup'@'%';
GRANT ALL PRIVILEGES ON speed_report.* TO 'speed_backup'@'%';
GRANT ALL PRIVILEGES ON go_fly2.* TO 'speed_backup'@'%';

FLUSH PRIVILEGES;


SELECT
    table_name AS `Table`,
    round(((data_length + index_length) / 1024 / 1024), 2) `Size in MB`
FROM information_schema.TABLES
WHERE table_schema = "speed"
ORDER BY (data_length + index_length) DESC;

SELECT
    table_name AS `Table`,
    round(((data_length + index_length) / 1024 / 1024), 2) `Size in MB`
FROM information_schema.TABLES
WHERE table_schema = "go_fly2"
ORDER BY (data_length + index_length) DESC;