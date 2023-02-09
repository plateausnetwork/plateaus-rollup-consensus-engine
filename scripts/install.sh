#mkdir ./data
#touch ./data/data.json

# gen contract implementation in .go
echo "generating .go based on contracts abi"
abigen --abi=build/contracts/Lottery.abi --pkg=contract --out=pkg/plateaus/contract/lottery.go --type=lottery
abigen --abi=build/contracts/LotteryValidation.abi --pkg=contracts --out=pkg/network/polygon/contracts/lotteryvalidation.go --type LotteryValidation
abigen --abi=build/contracts/PlateausValidation.abi --pkg=contracts --out=pkg/network/polygon/contracts/plateausvalidation.go --type PlateausValidation
