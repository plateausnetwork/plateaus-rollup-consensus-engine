# creating data files
NAME=".plateaus-consensus-rollup"
BINARY_NAME="plateaus-consensus"
HOME_DIR="$HOME/$NAME"
MAIN_PACKAGE_NAME="github.com/rhizomplatform/plateaus-rollup-consensus-engine"

if [ -d "$HOME_DIR" ]; then
    echo "$HOME_DIR already exists"
    echo "installation was aborted"
    exit
else
    mkdir $HOME_DIR
    mkdir "$HOME_DIR/data"
    mkdir "$HOME_DIR/config"
    touch "$HOME_DIR/data/data.json"
    echo "{}" > "$HOME_DIR/data/data.json"

    # creating config file
    touch "$HOME_DIR/data/data.json"
    cp ./config/config-testing.yml "$HOME_DIR/config/config.yml"
fi

go build -ldflags="-X '$MAIN_PACKAGE_NAME/config.homeDir=$HOME_DIR'" -o $BINARY_NAME ./cmd/consensus/main.go

mv $BINARY_NAME "$HOME/go/bin"
