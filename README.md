# fsexplorer

Very quick-and-dirty sketch of filesystem explorer with web UI.

## Build instructions

```bash
# You need to have Node / YARN and Go install
git clone https://github.com/lnsp/fsexplorer.git
cd fsexplorer/client && yarn install && yarn run generate
cd .. && make
```

You should now have a `fsexplorer` binary in the top-level directory.

## Usage

```
fsexplorer [-static=false] [-address=localhost:9876] [-basedir=/]
```

To run `fsexplorer` for your current working directory, execute

```bash
fsexplorer -static -address=localhost:9876 -basedir=$(pwd)
```

## Navigation

It is highly recommended to use the keyboard to navigate the UI.

key | command
-|-
h | go back
j | move cursor down
k | move cursor up
l | enter directory / view file
esc | close file viewer
d | download file (when viewing file)
```
