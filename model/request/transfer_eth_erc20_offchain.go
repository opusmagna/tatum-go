package request

type TransferEthErc20Offchain struct {
	BaseTransferEthErc20Offchain *BaseTransferEthErc20Offchain `validate:"required"`
	Mnemonic                     string                        `validate:"min=1,max=500"`
	Index                        uint32                        `validate:"min=0"`
	PrivateKey                   string                        `validate:"min=66,max=66"`
}
