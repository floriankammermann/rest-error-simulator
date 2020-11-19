#!/bin/bash
: ${host:=localhost:8080}
echo "use host: $host"
curl $host/control
echo ""
echo "*********************"
echo "set error ratio to 1"
echo "*********************"
curl $host/control/error?errorratio=1 -X POST
curl $host/control
echo ""
echo "*********************"
echo "get best-tools 3 times" 
echo "*********************"
for i in {1..3}; do
    curl -Is $host/best-tools | head -1 | cut -d " " -f2 && curl -s $host/metrics | egrep "^response.*"
done
echo "*********************"
echo "set error ratio to 100"
echo "*********************"
curl $host/control/error?errorratio=100 -X POST
curl $host/control
echo ""
echo "*********************"
echo "get best-tools 3 times" 
echo "*********************"
for i in {1..3}; do
    curl -Is $host/best-tools | head -1 | cut -d " " -f2 && curl -s $host/metrics | egrep "^response.*"
done
echo "*********************"
echo "set error ratio to 50"
echo "*********************"
curl $host/control/error?errorratio=50 -X POST
curl $host/control
echo ""
echo "*********************"
echo "get best-tools 6 times" 
echo "*********************"
for i in {1..6}; do 
    curl -Is $host/best-tools | head -1 | cut -d " " -f2 && curl -s $host/metrics | egrep "^response.*"
done


