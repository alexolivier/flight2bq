# Flight2BQ
Stream ADS-B data (dump1080 on a RaspberryPi etc) to Google BigQuery

## Table Schema
```
timestamp	TIMESTAMP	NULLABLE	
hexid	STRING	NULLABLE	
ident	STRING	NULLABLE	
squawk	INTEGER	NULLABLE	
alt	INTEGER	NULLABLE	
speed	INTEGER	NULLABLE	
airGround	STRING	NULLABLE	
lat	FLOAT	NULLABLE	
lon	FLOAT	NULLABLE	
heading	INTEGER	NULLABLE	
```