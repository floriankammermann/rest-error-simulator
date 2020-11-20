# rest error simulator

* Simulates errors on REST API (data plane)
* Available Simulations
  * http response code
  * response time
* Rate of errors can be controlled over a REST API (control plane)

## What for a rate do you get

| errorratio | percentage | responsecode |
|------------|------------|--------------|
| 1          | 100        | 200          |
| 2          | 98         | 200          |
| 3          | 97         | 200          |
| 4          | 96         | 200          |
| 5          | 95         | 200          |
| 6          | 94         | 200          |
| 7          | 93         | 200          |
| 8          | 92         | 200          |
| 9          | 91         | 200          |
| 10         | 90         | 200          |
| 11         | 89         | 200          |
| 12         | 88         | 200          |
| 13         | 87         | 200          |
| 14         | 86         | 200          |
| 15         | 83         | 200          |
| 16         | 83         | 200          |
| 17         | 80         | 200          |
| 18         | 80         | 200          |
| 25         | 75         | 200          |
| 33         | 66         | 200          |
| 50         | 50         | 200          |
| 100        | 0          | 200          |

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
