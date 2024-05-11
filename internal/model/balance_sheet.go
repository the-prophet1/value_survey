package model

type CurrentAsset struct {
	Cash                 Int // 现金
	ShortInvestment      Int // 短期投资
	ReceivableBill       Int // 应收票据
	ReceivableCash       Int // 应收账款
	ReceivableFinancing  Int // 应收款融资
	PrepayCash           Int // 预付款项
	Stock                Int // 存货
	ContractAsset        Int // 合同资产
	OtherReceivableCash  Int // 其他应收账款
	ShortNonCurrentAsset Int // 一年内到期的非流动资产
	OtherCurrentAsset    Int // 其他流动资产
	Other                map[string]Int
}

type NonCurrentAsset struct {
	Grant                  Int // 发放贷款或垫款
	DebtInvestment         Int // 债权投资
	OtherDebtInvestment    Int // 其他债权投资
	LongReceivableCash     Int // 长期应收账款
	LongStock              Int // 长期股权投资
	OtherEquity            Int // 其他权益工具投资
	OtherNonCurrentFinance Int // 其他非流动金融资产
	Estate                 Int // 投资性地产
	FixedAsset             Int // 固定资产
	ProgressConstruction   Int // 在建工程
	UseRight               Int // 使用权资产
	IntangibleAsset        Int // 无形资产
	Goodwill               Int // 商誉
	LongExpense            Int // 长期待摊费用
	DeferredTax            Int // 延递所得税资产
	OtherNonCurrentAsset   Int // 其他非流动资产
	Other                  map[string]Int
}

type CurrentLiability struct {
	ShortLoan                    Int // 短期借款
	Deposit                      Int // 吸收存款及同行存放
	BorrowedFund                 Int // 拆入资金
	DerivativeFinancialLiability Int // 衍生金融负债
	PayableBill                  Int // 应付票据
	PayableCash                  Int // 应付账款
	AdvancesReceived             Int // 预收款项
	ContractualLiability         Int // 合同负债
	EmployeeCompensation         Int // 应付职工薪酬
	Taxes                        Int // 应交税费
	NonCurrentLiability          Int // 一年内到期非流动负债
	OtherCurrentLiability        Int // 其他流动负债
	Other                        map[string]Int
}

type NonCurrentLiability struct {
	LongLoan                 Int // 长期借款
	LeaseLiability           Int // 租赁负债
	LongPayable              Int // 长期应付款
	LongEmployeeCompensation Int // 长期应付职工薪酬
	DeferredRevenue          Int // 延递收益
	DeferredTaxLiabilities   Int // 递延所得税负债
	Other                    map[string]Int
}
