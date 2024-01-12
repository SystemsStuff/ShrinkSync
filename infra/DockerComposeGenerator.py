def generate_docker_compose(map_node_count, reduce_node_count):
    port_number = 8080
    compose_file_content = f'''services:'''

    master_definition = f'''
  master:
    build: .
    image: shrink-sync-node
    container_name: master
    ports:
        - "{port_number}:8080"
    volumes:
        - /var/run/docker.sock:/var/run/docker.sock    
'''

    compose_file_content+=master_definition
    port_number+=1

    for i in range(1, map_node_count + 1):
        service_name = f'map-{i}'
        compose_file_content += f'''\
  {service_name}:
    build: .
    image: shrink-sync-node
    container_name: {service_name}
    ports:
        - "{port_number}:8080"
    volumes:
        - /var/run/docker.sock:/var/run/docker.sock
'''
        port_number+=1

    for i in range(1, reduce_node_count + 1):
        service_name = f'reduce-{i}'
        compose_file_content += f'''\
  {service_name}:
    build: .
    image: shrink-sync-node
    container_name: {service_name}
    ports:
        - "{port_number}:8080"
    volumes:
        - /var/run/docker.sock:/var/run/docker.sock
'''
        port_number+=1

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