# rest error simulator

* Simulates errors on REST API (data plane)
* Available Simulations
  * http response code
  * response time
* Rate of errors can be controlled over a REST API (control plane)

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
  * `export RES_REQUESTFREQUENCYINSEC=200`
  * `export RES_ENDPOINT=500`
* `make build-client`
* `bin/rest-error-simulator-client
* set the control `curl localhost:8080/control?frequency=3`
