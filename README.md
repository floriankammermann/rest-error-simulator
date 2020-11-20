# rest error simulator

* Simulates errors on REST API (data plane)
* Available Simulations
  * http response code
  * response time
* Rate of errors can be controlled over a REST API (control plane)

## What for a rate do you get

errorratio=1   100% 200
errorratio=2   98% 200
errorratio=3   97% 200
errorratio=4   96% 200
errorratio=5   95% 200
errorratio=6   94% 200
errorratio=7   93% 200
errorratio=8   92% 200
errorratio=9   91% 200
errorratio=10  90% 200
errorratio=11  89% 200
errorratio=12  88% 200
errorratio=13  87% 200
errorratio=14  86% 200
errorratio=15  83% 200
errorratio=16  83% 200
errorratio=17  80% 200
errorratio=18  80% 200
errorratio=25  75% 200
errorratio=33  66% 200
errorratio=50  50% 200
errorratio=100 0%  200

# run server

* set envs (not mandatory) 
  * `export RES_RESPONSECODESUCCESS=200`
  * `export RES_RESPONSECODEFAILURE=500`
  * `export RES_RESPONSECODESUCCESSFAILURERATIO=50`
* `make build-server`
* `bin/rest-error-simulator-server
* get the data `curl localhost:8080/best-tools`
* set the control `curl localhost:8080/control/error?errorratio=50`
* get the metrics `curl localhost:8080/metrics`

# run client

* set envs (not mandatory) 
  * `export RES_REQUESTFREQUENCYINSEC=1`
  * `export RES_ENDPOINT=http://rest-error-simulator.com`
* `make build-client`
* `bin/rest-error-simulator-client
* set the control `curl localhost:8080/control?frequency=3`
