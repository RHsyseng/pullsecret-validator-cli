# pullsecret-validator-cli


## Install it (Windows/macOS/Linux)

- To Install the binary in your system just get the latest release:

[https://github.com/alknopfler/ocp-release/releases/latest](https://github.com/alknopfler/ocp-release/releases/latest)

- Download the file for your OS distribution (Go Binary complied for several OS)
- Unzip/unTar the binary into a folder inside your $PATH


## Use it

To show the instructions or the help message 
```commandline
./ocp-release -h
```

To get the best release candidate tag:
```commandline
./ocp-release -v <version> -c <condition>
```

For example:
```commandline
./ocp-release -v nightly -c assisted-metal
```

You could also match an array of conditions:

```commandline
./ocp-release -v nightly -c assisted-metal aws metal-ipi
```


