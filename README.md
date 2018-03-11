# Flight2BQ
Stream ADS-B data (RTLSDR dump1080 on a RaspberryPi etc) to Google BigQuery.

Example live public dataset on London, UK can be found in `alex-olivier.flighttracker_dev.aircraft_stream` in BigQuery.

Please let me know if you would like to add data from your location.

## Todo
- [ ] Get airframe information from https://opensky-network.org/
- [ ] Schedules dedupe query

## Table Schema
```
timestamp	TIMESTAMP	NULLABLE	
hexid STRING	NULLABLE	
ident	STRING	NULLABLE	
squawk	INTEGER	NULLABLE	
alt	INTEGER	NULLABLE	
speed	INTEGER	NULLABLE	
airGround	STRING	NULLABLE	
lat	FLOAT	NULLABLE	
lon	FLOAT	NULLABLE	
heading	INTEGER	NULLABLE	
```
