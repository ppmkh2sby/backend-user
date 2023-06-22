# Backend User - RFID
This service is core APIs of user management.

## Main Features
- User Management

## Status
Production

## Running Services
### Backend User API Service
Create backend_user.yaml in _bin
Tamplate config file :
```bash
service_data:
      address: :8001
       
source_data:
 postgresdb_server: localhost
  postgresdb_port: 5432
  postgresdb_name: db
  postgresdb_username: user
  postgresdb_password: pass
  postgresdb_timeout: 5
```


## How To
Build Service
```bash
./_scripts/build.sh
```

Run Service
```bash
./_scripts/run.sh
```

Get Dependencies
```bash
./_scripts/get_dep.sh  
```

Run Unit Testing
```bash
./_scripts/test_unit.sh  
```

Run Unit Testing Race
```bash
./_scripts/test_race.sh
```

Generate Coverage File
```bash
./_scripts/test_coverage_report.sh
```

## License
IT-PPMKH2SBY