package config

const Namespace = "stern-x"

const Default = `
db:  
  driverName: "postgres"
  dataSourceName: "user=postgres password=1234 host=localhost dbname=postgres sslmode=disable" 

server:
  addr: "0.0.0.0:8090"

log:
 level: 
`
