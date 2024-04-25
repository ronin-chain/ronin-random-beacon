# Random Beacon Chain

The `random-beacon-chain` project is a service that generates random numbers and submits them to the RoninChain for rotating validators.

## Background 
------
We run a service that listens for events from the blockchain, generates proofs based on these events, and submits them every 10 minutes daily. The service will retry n times before marking a submission as failed. 

## Flow 
The service uses event trackers and task processing mechanisms to handle the flow of operations:

1. **Chain Tracker**: This component runs at an interval of 3 seconds, tracking the latest processed and finalized blocks on the blockchain.

2. **Random Beacon Tracker**: This component tracks random beacon request events on the blockchain. It operates based on the `tracker_block_size` and `tracker_safe_block_range` settings, and updates the processed block.

3. **Task Creator**: This component creates tasks based on the events tracked by the Random Beacon Tracker.

4. **Task Sender**: This component handles the tasks created by the Task Creator and prepares them for proof submission.

5. **Task Syncer Sent**: This component handles sent tasks and updates their status to either success or failure.

6. **Task Reorg Checker**: This component handles task reorganization based on the finalized block.

7. **Task Cleaner**: This component cleans up tasks that are more than 2 months old.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine.

### Prerequisites

You need to have `make` installed on your machine. You can check if it's installed by running:

```bash
make --version
```

If `make` is not installed, you can install it by following the instructions here.

### Generate code

- Initializes schemas:

```shell
$ go run entgo.io/ent/cmd/ent new --target pkg/schema <your schemas>
```

- Generates assets from schemas:

```shell
$ go run entgo.io/ent/cmd/ent generate ./pkg/schema --target ./pkg/db --feature sql/upsert,sql/execquery
```


- Generates Random Beacon contract interface from ABI: follow [this instruction](https://geth.ethereum.org/docs/dapp/native-bindings)

```shell
$ abigen --abi <ABI JSON file>  --pkg contract --type RandomBeaconCoordinator --out pkg/contract/random_beacon_coordinator.go
```

# Building the Project
To build the project, navigate to the project directory in your terminal and run:

```
make build
```

# Running the Project
To run the project, use the following command:
```
make start ARGS="--config-file=config/"
```

# Generating a Random Number
The `random-beacon-chain` service generates a random number for use in the RoninChain. To generate a random number, use the following command:

```
make random ARGS="--config-file=config/ --period=5 --prev-beacon=10"
```

# Gen VRF Key
The `random-beacon-chain` service also supports generating a key. You can generate a key by running the following command:

```
make generate-key
```
