GreenEye energy monitor web app

This is a first stab at trying something of a new web app for power monitoring

-------
Setup the database
-------

cockroach start
cockroach user set energymon
cockroach sql -e 'CREATE DATABASE powerread'
cockroach sql -e 'GRANT ALL ON DATABASE powerread TO energymon'


