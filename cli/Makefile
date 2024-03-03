setup: infra datagrid

network:
	- docker network create shrink-sync-network -d bridge

infra: network
	- cd infra && python3 DockerComposeGenerator.py
	- cd infra && docker compose up -d --build
	- echo "\nInfra Setup Done"

datagrid: network
	- cd DataGrid && docker build -t datagrid .
	- docker run --name datagrid --network=shrink-sync-network -d -p 21:21 datagrid
	- echo "\nDatagrid Started on port 21"

clean:
	- cd infra && docker compose down --remove-orphans
	- rm -f infra/docker-compose.yaml
	- echo "\nInfra setup cleaned"
	- docker container stop datagrid || true && docker container rm datagrid || true
	- echo "\nDatagrid setup cleaned"

.PHONY: infra datagrid clean