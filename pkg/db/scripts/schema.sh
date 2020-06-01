#!/bin/sh

cd /flyway/sql/ddl

DB_SERVICE=$@
for F in $(ls)
do
    echo "Start process $F"
    /flyway/flyway -X -url=jdbc:mysql://$(DB_SERVICE)/openpitrix -user=root -validateOnMigrate=false -locations=filesystem:/flyway/sql/job migrate
done


