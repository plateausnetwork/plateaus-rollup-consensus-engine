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

    # creating config file
    touch "$HOME_DIR/data/data.json"
    cp ./config/config-testing.yml "$HOME_DIR/config/config.yml"
fi

# gen contract implementation in .go
echo "2. generating .go based on contracts abi"
abigen --abi=build/contracts/Lottery.abi --pkg=contract --out=pkg/plateaus/contract/lottery.go --type=lottery
abigen --abi=build/contracts/LotteryValidation.abi --pkg=contracts --out=pkg/network/contracts/lotteryvalidation.go --type LotteryValidation
abigen --abi=build/contracts/PlateausValidation.abi --pkg=contracts --out=pkg/network/polygon/contracts/plateausvalidation.go --type PlateausValidation

echo "'$MAIN_PACKAGE_NAME/config.homeDir=$HOME_DIR/config'"

go build -ldflags="-X '$MAIN_PACKAGE_NAME/config.homeDir=$HOME_DIR/config'" -o $BINARY_NAME ./cmd/consensus/main.go
