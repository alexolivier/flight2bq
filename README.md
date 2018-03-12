# Flight2BQ
Stream ADS-B data (RTLSDR dump1090 on a RaspberryPi etc) to Google BigQuery.

Example live public dataset on London, UK can be found in `alex-olivier.flighttracker_dev.aircraft_stream` in BigQuery.

Please let me know if you would like to add data from your location.

## Example Data
```
| timestamp | hexid | ident | squawk | alt | speed | airGround | lat | lon | heading |
| --------- | ----- | ----- | ------ | --- | ----- | --------- | --- | --- | ------ |
| 2018-03-11 21:33:21.000 UTC |	A023AF | EJM685 | 7660 | 10000 | 305 | A | 51.86767 | -0.02373 | 23	 |
| 2018-03-11 21:33:21.000 UTC | 4CA213 | EIN24K | 2020 | 12825 | 369 | A | 51.29451 | -0.50453 | 306 |
| 2018-03-11 21:33:21.000 UTC | 06A141 | QTR8132 | 7270 | 30550 | 469 | A | 51.27194 | 0.29724 | 98 |
| 2018-03-11 21:33:20.000 UTC | 781103 | CES552 | 2244 | 5125 | 260 | A | 51.53212 | -0.25757 | 67 |
| 2018-03-11 21:33:20.000 UTC | 48AF00 | LOT285 | 2013 | 8000 | 275 | A | 51.64124 | -0.26775 | 269 | 
| 2018-03-11 21:33:20.000 UTC | 40621C | BAW199 | 2217 | 3375 | 250 | A | 51.44371 | -0.33089 | 106 | 
| 2018-03-11 21:33:19.000 UTC | 4CA911 | RYR96JV | 3550 | 35125 | 401 | A | 51.46848 | 1.09154 | 272 |
| 2018-03-11 21:33:19.000 UTC | 4B1691 | SWR24H | 3053 | 10125 | 308 | A | 51.23309 | 0.1207 | 332	 | 
| 2018-03-11 21:33:19.000 UTC | 4CC272 | ICE477 | 4462 | 14775 | 389 | A | 51.49393 | 0.15782 | 332 |
```

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
