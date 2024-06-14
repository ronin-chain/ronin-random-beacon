# Random Beacon Chain

The `random-beacon-chain` project is a service that generates random numbers and submits them to the RoninChain for rotating validators.

## Background 
------
We run a service that listens for events from the blockchain, generates proofs based on these events, and submits them every 1 minute daily. The service will retry n times before marking a submission as failed. 

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
$ abigen --abi abi/RandomBeacon.abi.json --pkg contract --type RandomBeaconCoordinator --out pkg/contract/random_beacon_coordinator.go
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
make random ARGS="--config-file=config/ --period=5 --prev-beacon=10 --chain-id=2020 --verifying-contract=0xA60c1e07fa030E4B49Eb54950ADb298Ab94dD312"
```

Example Output

```
INFO [06-09|12:27:27.496] Loaded config                            period=5 Prev Beacon=10
INFO [06-09|12:27:27.499] Output                                   res="&{Proof:{Pk:[+93836631653650151269353316240652443878617348030413819343231539155313056694026 +51974711041462004316159568149910806708460680222297123882950673380506412020600] Gamma:[+16161261169515536958452562390271672783929847571978359040698519788687374050289 +30145072877981152360671926035825153273290341499426424463874939044889519383834] C:+38898946712616960857822006366651309325850496135209239141585035907246086220062 S:+109332194886916398362305761008660543127782491667094916998947232884969844049695 Seed:+108590375316084432207008885717060140159794805667776380061032813864368595446299 UWitness:0x5F2830C40430Dc64BB54c2CEB02216c13c52C303 CGammaWitness:[+25672313979954775441181944989962582624878321837874361189037485351937409492632 +19431072094724670564167429239325932030890049088213627524827916329414877729414] SHashWitness:[+36348122650986015621328136828569148696601679562155328689432650485823543701279 +357829478391814333228283967741661477394524155426507391833336708466134744474] ZInv:+60118025322348429685592956744809014495540397930447947433220045315685662522821} RandomRequest:{Period:+5 PrevBeacon:+10 ChainId:+2020 VerifyingContract:0x0000000000000000000000000000000000000000}}"
        1.12 real         0.05 user         0.02 sys
```


# Gen VRF Key
The `random-beacon-chain` service also supports generating a key. You can generate a key by running the following command:

```
make generate-key
```
