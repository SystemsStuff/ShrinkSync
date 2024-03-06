# ShrinkSync
A MapReduce implementation

## A 1000000 feet view

- Programming Language - GoLang
- Scope
  - Master Node
    - Scheduling Jobs
    - Maintaining state of worker nodes
    - Work Splitter - Takes in (M,R) and partitions input into job files and assign nodes
    - Job Scheduler - Assigns jobs to map and reduce worker nodes
    - Handle failures - Need not spawn, but just re-assign the work
  - Worker Node(MAP)
    - Message Passing - Communicate with master node
    - Write to disk
  - Worker Node(REDUCE)
    - Message Passing - Communicate with master node
    - Write to Output Location to disk
    - Read data from map worker disk
    - Atomic rename of files[Proposal for now - While committing final output, reduce worker could call endpoint on Datagrid to rename file]
- Infra
  - Setup docker cluster for local development and testing
  - Setup cluster for data grid(to store input data)
  - Cloud deployment maybe later



## What are we building?
- A library to create MapReduce jobs
- A CLI to run the job
  - CLI takes in the MapReduce job files directory
  - This CLI would start the infra with the job files to be executed

## What are we not doing?
- Full feature replication of GFS - Large data should ideally be saved in GFS nodes in chunks. We use a simple FTP server(named `datagrid`) serving chunks of data, simulating GFS to some extent.
- Job orchestration node - Ideally, there should be an orchestration node that receives the job, assigns idle nodes and reports status back to user after execution. But, we skip this step with a CLI that spins up new nodes to run the job.
- gRPC - Some parts of the server-to-server communication should ideally happen using gRPC for improved performance. But, we choose to use HTTP-REST for this setup to reduce complexity.

## Design
### Defining and Extracting a MapReduce job
A map reduce job is defined by extending the `Map` struct and the `Reduce` struct provided in the library. The `map()` and the `reduce()` methods are overridden with user's implementation. There will be an `init()` method for both structs, which will execute some library-level functionalities and then run the `master()`, `map()` or `reduce()` jobs accordingly. The node knows of it's context(whether it is a map, reduce or master node) based on an environment variable. If the environment variable is not set, it indicates that the executable is running on user machine. Handle this with an error log.

### CLI
After defining the job in a directory, a CLI can be used to start the job with files in that directory. Using that CLI, user can set number of map nodes, number of reduce nodes and data location. The job files are copied over to docker containers before starting them. The user code is shared to the MapReduce infra using this approach.

## Usage flow
A user who wants to use ShrinkSync will first download the CLI. Then, they will initialize a folder as a ShrinkSync job(something similar to an `init` command). This initialization will create a skeleton project where users can write their code. There will be a config file that takes in parameters like location of input data, number of map nodes, number of reduce nodes etc. Once they prepared their job with required `map` and `reduce` functions, they use the CLI to start the job(something like a `start` command). Before this, users need to make sure the datagrid from which the job takes input is `UP`. The CLI will then start the infra(master, map and reduce nodes) with the user code already in them.

The CLI utility keeps giving updates on the job through the terminal using the "infra-health" endpoint on master node. Once the job is executed, the infrastructure is brought down and CLI returns the location of the output.