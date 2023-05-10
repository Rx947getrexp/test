package config

type Gas struct {
	CreateTransaction    int `mapstructure:"create-transaction" json:"createTransaction"`
	TriggerSmartContract int `mapstructure:"trigger-smart-contract" json:"triggerSmartContract"`
	BindChain            int `mapstructure:"bind-chain" json:"bindChain"`
	ImgTransfer          int `mapstructure:"img-transfer" json:"imgTransfer"`
	ImgChainTransfer     int `mapstructure:"img-chain-transfer" json:"imgChainTransfer"`
	BindDirectAddress    int `mapstructure:"bind-direct-address" json:"bindDirectAddress"`
	CreateContract       int `mapstructure:"create-contract" json:"createContract"`
	TakeMiningFin        int `mapstructure:"take-mining-fin" json:"takeMiningFin"`
	CreateFlowPool       int `mapstructure:"create-flow-pool" json:"createFlowPool"`
	TakeFlowPool         int `mapstructure:"take-flow-pool" json:"takeFlowPool"`
	SendSwap             int `mapstructure:"send-swap" json:"sendSwap"`
	TakeFlowPoolProfit   int `mapstructure:"take-flow-pool-profit" json:"takeFlowPoolProfit"`
	AddPosition          int `mapstructure:"add-position" json:"addPosition"`
	Authorize            int `mapstructure:"authorize" json:"authorize"`
	AddFlow              int `mapstructure:"add-flow" json:"addFlow"`
	CrossChainHashFlash  int `mapstructure:"cross-chain-hash-flash" json:"crossChainHashFlash"`
}
