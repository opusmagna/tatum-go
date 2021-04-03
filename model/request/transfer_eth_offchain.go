package request

type TransferEthOffchain struct {
	Mnemonic                     string                        `validate:"min=1,max=500"`
	Index                        uint32                        `validate:"min=0"`
	PrivateKey                   string                        `validate:"min=66,max=66"`
	Data                         string                        `validate:"max=50000"`
	BaseTransferEthErc20Offchain *BaseTransferEthErc20Offchain `validate:"required"`
}
