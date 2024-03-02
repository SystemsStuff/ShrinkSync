def generate_docker_compose(map_node_count, reduce_node_count):
    compose_file_content = f'''services:'''

    master_definition = f'''
  master:
    build: .
    image: shrink-sync-node
    container_name: master
    environment:
        - MAP_NODE_COUNT={map_node_count}
        - REDUCE_NODE_COUNT={reduce_node_count}
        - NAME=master
'''

    compose_file_content+=master_definition

    for i in range(1, map_node_count + 1):
        service_name = f'map-{i}'
        compose_file_content += f'''\
  {service_name}:
    build: .
    image: shrink-sync-node
    container_name: {service_name}'''
        compose_file_content+=f'''
    environment:
        - MAP_NODE_COUNT={map_node_count}
        - REDUCE_NODE_COUNT={reduce_node_count}
        - NAME={service_name}
'''

    for i in range(1, reduce_node_count + 1):
        service_name = f'reduce-{i}'
        compose_file_content += f'''\
  {service_name}:
    build: .
    image: shrink-sync-node
    container_name: {service_name}'''
        compose_file_content+=f'''
    environment:
        - MAP_NODE_COUNT={map_node_count}
        - REDUCE_NODE_COUNT={reduce_node_count}
        - NAME={service_name}
'''

    network_definition = f'''\
networks:
  default:
    name: shrink-sync-network
    driver: bridge   
    
'''
    compose_file_content += network_definition

    with open('docker-compose.yaml', 'w') as file:
        file.write(compose_file_content)


map_node_count = int(input("Enter the number of map nodes count for the current execution: "))
reduce_node_count = int(input("Enter the number of reduce nodes count for the current execution: "))

generate_docker_compose(map_node_count, reduce_node_count)
print("Completed generation of Docker Compose file")