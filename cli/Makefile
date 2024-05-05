setup: infra datagrid

network:
	- docker network create shrink-sync-network -d bridge

infra: network datagrid
	- python3 DockerComposeGenerator.py
	- docker compose up -d --build
	- echo "\nInfra Setup Done"

datagrid: network
	- cd ../DataGrid && docker build -t datagrid .
	- docker run --name datagrid --network=shrink-sync-network -d -p 21:21 datagrid
	- echo "\nDatagrid Started on port 21"

clean:
	- docker compose down --remove-orphans
	- docker network rm shrink-sync-network
	- rm -f docker-compose.yaml
	- echo "\nInfra setup cleaned"
	- docker container stop datagrid || true && docker container rm datagrid || true
	- echo "\nDatagrid setup cleaned"

.PHONY: infra datagrid network clean