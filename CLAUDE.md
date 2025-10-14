# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Xone Chain is a Cosmos SDK-based blockchain with Ethereum Virtual Machine (EVM) compatibility, built on Evmos v16. It integrates Cosmos SDK modules with Ethereum-compatible smart contract execution, supporting both Cosmos transactions and Ethereum transactions.

**Chain identifier**: `xone_33772211-1`
**Native token**: `uxoc` (micro XOC), with display denomination `xoc` (1 xoc = 10^18 uxoc)
**Binary name**: `xoned`

## Build Commands

```bash
# Build the binary
make build

# Install to $GOPATH/bin
make install

# Initialize chain for development
make init

# Build with Ignite CLI (alternative)
make ignite-build

# Start node (standard)
make start

# Start node with debug mode (includes EVM tracer)
make debug

# Reset: install + init + start
make reset

# Protocol buffer generation
make proto-gen

# Docker build
make docker
```

## Testing

The project uses standard Go testing:
```bash
# Run all tests
go test ./...

# Run tests in a specific package
go test ./x/xone/keeper/...

# Run a single test
go test -run TestSpecificTest ./x/xone/keeper/
```

## Architecture

### Module Structure

**Custom modules:**
- `x/xone`: Main custom module for xone chain functionality
  - Located in `/x/xone/`
  - Standard Cosmos SDK module structure (keeper, types, client/cli)

**Evmos modules (from forked Evmos v16):**
- `x/evm`: Ethereum Virtual Machine support
- `x/feemarket`: EIP-1559 dynamic fee market
- `x/erc20`: Token conversion between Cosmos and ERC-20
- `x/epochs`: Time-based hooks system (added via fork)
- `x/vesting`: Token vesting with EVM support

**Standard Cosmos SDK modules:**
- Bank, Staking, Distribution, Governance, Slashing, etc.
- IBC (Inter-Blockchain Communication) with transfer support
- ICA (Interchain Accounts) controller and host

### Key Architecture Points

1. **Dual Transaction Support**: The chain processes both Cosmos SDK transactions and Ethereum transactions through the same consensus layer.

2. **AnteHandler Chain**: Located in `app/ante/`, this implements custom transaction validation:
   - `ante/cosmos/`: Cosmos transaction handlers (fee validation, authz, vesting)
   - `ante/evm/`: EVM transaction handlers (signature verification, nonce, gas)
   - Configured in `app/ante/handler_options.go`

3. **App Construction**: `app/app.go` contains the main application setup:
   - Keeper initialization order matters (dependencies between keepers)
   - Module manager controls BeginBlock/EndBlock order
   - Store keys are registered in `newStoreKey()`
   - EVM precompiles are registered via `evmKeeper.WithPrecompiles()`

4. **Fork Management**: `app/fork.go` defines chain upgrades at specific block heights:
   - Gas limit increases: ForkGasLimitMainNetHeight, ForkGasLimitTestNetHeight
   - Epochs module additions: ForkAddEpochsMainNetHeight, ForkAddEpochsTestNetHeight
   - Fork logic runs in `BeginBlocker` via `forkBlocker()`
   - Store upgrades handled by `setUpgradeStoreLoader()`

5. **Chain ID Logic**: The chain uses Ethereum-style chain IDs parsed from the Cosmos chain ID string:
   - Mainnet: EvmChainID_Mainnet
   - Testnet: EvmChainID_Testnet
   - Extracted via `evmostypes.ParseChainID()`

### Important Dependencies (Forked)

The project uses custom forks (see `go.mod` replace directives):
- `github.com/hello-xone/xone_cosmos-sdk`: Custom Cosmos SDK fork
- `github.com/hello-xone/xone_go-ethereum`: Custom Ethereum fork
- `github.com/hello-xone/xone_evmos`: Custom Evmos v16 fork

These forks may contain xone-specific modifications not present in upstream.

## Configuration

### Development Configuration

`config.yml` configures the development environment:
- Genesis accounts and validator setup
- API/gRPC endpoints (ports 1317, 9090, 9091)
- EVM denomination set to `uxoc`
- Consensus timeout: 3s
- Fee market parameters (base_fee, min_gas_price)

### Node Configuration

Key parameters for running a node:
- **RPC**: 0.0.0.0:26657
- **JSON-RPC** (Ethereum): 0.0.0.0:8545
- **JSON-RPC WebSocket**: 0.0.0.0:8546
- **API**: 0.0.0.0:1317
- **gRPC**: 0.0.0.0:9090
- **Minimum gas price**: 10000000uxoc

## Protocol Buffers

Proto files in `proto/xone/xone/`:
- `genesis.proto`: Genesis state
- `params.proto`: Module parameters
- `query.proto`: Query service definitions
- `tx.proto`: Transaction message types

Generate code after proto changes:
```bash
make proto-gen
```

## Module Account Permissions

Defined in `maccPerms` in `app/app.go`:
- `xone` module: minter, burner, staking permissions
- `evm` module: minter, burner
- `erc20` module: minter, burner

## Development Notes

- The binary name is `xoned` (not `xone`)
- Command structure: `xoned [command]` (uses Cosmos SDK server commands)
- Version string format: `{Version} {BuildTime} {GitCommit}` (set via ldflags in Makefile)
- Default home directory: `~/.xone`
- Account address prefix: `x` (configured in `app/app.go`)

## Repository Structure

```
app/                  # Application setup and configuration
├── ante/            # Transaction validation (AnteHandlers)
├── app.go           # Main application construction
├── fork.go          # Fork/upgrade logic
├── encoding.go      # Codec setup
cmd/xoned/           # Binary entry point and CLI
x/xone/              # Custom xone module
├── keeper/          # State management
├── types/           # Type definitions
├── client/cli/      # CLI commands
proto/xone/          # Protobuf definitions
docs/                # OpenAPI documentation
testutil/            # Testing utilities
scripts/             # Build and utility scripts
evmos/               # Evmos v16 fork (submodule or vendored)
```