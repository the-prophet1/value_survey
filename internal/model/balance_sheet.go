package model

type CurrentAsset struct {
	Cash                 int // 现金
	ShortInvestment      int // 短期投资
	ReceivableBill       int // 应收票据
	ReceivableCash       int // 应收账款
	ReceivableFinancing  int // 应收款融资
	PrepayCash           int // 预付款项
	Stock                int // 存货
	ContractAsset        int // 合同资产
	OtherReceivableCash  int // 其他应收账款
	ShortNonCurrentAsset int // 一年内到期的非流动资产
	OtherCurrentAsset    int // 其他流动资产
	Other                map[string]int
}

type NonCurrentAsset struct {
	Grant                  int // 发放贷款或垫款
	DebtInvestment         int // 债权投资
	OtherDebtInvestment    int // 其他债权投资
	LongReceivableCash     int // 长期应收账款
	LongStock              int // 长期股权投资
	OtherEquity            int // 其他权益工具投资
	OtherNonCurrentFinance int // 其他非流动金融资产
	Estate                 int // 投资性地产
	FixedAsset             int // 固定资产
	ProgressConstruction   int // 在建工程
	UseRight               int // 使用权资产
	IntangibleAsset        int // 无形资产
	Goodwill               int // 商誉
	LongExpense            int // 长期待摊费用
	DeferredTax            int // 延递所得税资产
	OtherNonCurrentAsset   int // 其他非流动资产
	Other                  map[string]int
}

type CurrentLiability struct {
	ShortLoan                    int // 短期借款
	Deposit                      int // 吸收存款及同行存放
	BorrowedFund                 int // 拆入资金
	DerivativeFinancialLiability int // 衍生金融负债
	PayableBill                  int // 应付票据
	PayableCash                  int // 应付账款
	AdvancesReceived             int // 预收款项
	ContractualLiability         int // 合同负债
	EmployeeCompensation         int // 应付职工薪酬
	Taxes                        int // 应交税费
	NonCurrentLiability          int // 一年内到期非流动负债
	OtherCurrentLiability        int // 其他流动负债
	Other                        map[string]int
}

type NonCurrentLiability struct {
	LongLoan                 int // 长期借款
	LeaseLiability           int // 租赁负债
	LongPayable              int // 长期应付款
	LongEmployeeCompensation int // 长期应付职工薪酬
	DeferredRevenue          int // 延递收益
	DeferredTaxLiabilities   int // 递延所得税负债
	Other                    map[string]int
}
