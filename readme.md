# ignitetutorialnel

**ignitetutorialnel** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

First, clone this repository:

```
https://github.com/mitchellnel/ignite-tutorial-nel
```

Then, invoke:

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

After the blockchain has started, open a separate terminal window, and invoke:

```
ignite-tutorial-neld q ignitetutorialnel hello
```

This will return:

```
{
  "text": "Hello, Ignite CLI!",
}
```

Alternatively, proceed in a browser to http://localhost:1317 to interact and make queries to the chain via a GUI.
