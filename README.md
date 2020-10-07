# Network and Subnet Calculator
## How to run
* 1. `docker build -t net-calc .`
* 2. `docker run -it --rm --name net-app net-calc`

## How to use
* Put Network address as `155.100.128.0`
* Put Mask address as `255.255.148.0`
* Put count host as `255`
* Then view the answer

### In the future
* Solution of the problem of the second subnet
* Dockeriza
* Terminal with Color
* More flexible


* cuando tenemos la misma red repetida, es la ultima direccion de subred mas 2 (porque uno es el broadcast y 2 es la primera direccion de red)