# vandame

## Config
``yaml
machine:
  name: "caxias"
  networkInterfaces:
  - name: "enps3"
    ip: "192.168.0.2"
  - name: "enp4s"
    ip: "10.0.2.3"

cluster:
  name: "rio"
  state: "new"
  machines:
  - name: "macae"
    clusterIp: "192.168.0.3"
  - name: "campos"
    clusterIp: "192.168.0.4"
  - name: "araruama"
    clusterIp: "192.168.0.5"
``

Rodar o comando:

``bash
./config-generator.sh config.yaml
```
