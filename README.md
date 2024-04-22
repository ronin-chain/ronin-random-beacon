# Random Beacon Chain

The `random-beacon-chain` project is a service that generates random numbers and submits them to the RoninChain for rotating validators.

## Background 
------
We run service for listening the events from chain for generating the proof and submit after 10 minutes daily. It will retry n times before marking failed. 

## Flow 
The idea is to use event trackers, task processing for processing via file 

### Chain Tracker
Interval: 3s

Tracking the latest block process.

### Random beacon tracker

Tracking events with setting `tracker_block_size` and `tracker_safe_block_range`

### Task processor 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.

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
