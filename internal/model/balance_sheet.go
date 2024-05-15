package model

import "value-survey/pkg/idgenerate"

type CurrentAsset struct {
	Cash                 Int   `json:"cash" gorm:"column:cash;bigint;comment:现金"`
	ShortInvestment      Int   `json:"short_investment" gorm:"column:short_investment;bigint;comment:短期投资"`
	ReceivableBill       Int   `json:"receivable_bill" gorm:"column:receivable_bill;bigint;comment:应收票据"`
	ReceivableCash       Int   `json:"receivable_cash" gorm:"column:receivable_cash;bigint;comment:应收账款"`
	ReceivableFinancing  Int   `json:"receivable_financing" gorm:"column:receivable_financing;bigint;comment:应收款融资"`
	PrepayCash           Int   `json:"prepay_cash" gorm:"column:prepay_cash;bigint;comment:预付款项"`
	Stock                Int   `json:"stock" gorm:"column:stock;bigint;comment:存货"`
	ContractAsset        Int   `json:"contract_asset" gorm:"column:contract_asset;bigint;comment:合同资产"`
	OtherReceivableCash  Int   `json:"other_receivable_cash" gorm:"column:other_receivable_cash;bigint;comment:其他应收账款"`
	ShortNonCurrentAsset Int   `json:"short_non_current_asset" gorm:"column:short_non_current_asset;bigint;comment:一年内到期的非流动资产"`
	OtherCurrentAsset    Int   `json:"other_current_asset" gorm:"column:other_current_asset;bigint;comment:其他流动资产"`
	ExtraCurrentAsset    SIMap `json:"extra_current_asset" gorm:"column:extra_current_asset;charset(2048);comment:额外的流动资产"`
}

func (currentAsset *CurrentAsset) GenerateID() {
	if currentAsset.Cash != nil {
		currentAsset.Cash.ID = idgenerate.Generate()
	}
	if currentAsset.ShortInvestment != nil {
		currentAsset.ShortInvestment.ID = idgenerate.Generate()
	}
	if currentAsset.ReceivableBill != nil {
		currentAsset.ReceivableBill.ID = idgenerate.Generate()
	}
	if currentAsset.ReceivableCash != nil {
		currentAsset.ReceivableCash.ID = idgenerate.Generate()
	}
	if currentAsset.ReceivableFinancing != nil {
		currentAsset.ReceivableFinancing.ID = idgenerate.Generate()
	}
	if currentAsset.PrepayCash != nil {
		currentAsset.PrepayCash.ID = idgenerate.Generate()
	}
	if currentAsset.Stock != nil {
		currentAsset.Stock.ID = idgenerate.Generate()
	}
	if currentAsset.ContractAsset != nil {
		currentAsset.ContractAsset.ID = idgenerate.Generate()
	}
	if currentAsset.OtherReceivableCash != nil {
		currentAsset.OtherReceivableCash.ID = idgenerate.Generate()
	}
	if currentAsset.ShortNonCurrentAsset != nil {
		currentAsset.ShortNonCurrentAsset.ID = idgenerate.Generate()
	}
	if currentAsset.OtherCurrentAsset != nil {
		currentAsset.OtherCurrentAsset.ID = idgenerate.Generate()
	}
}

type NonCurrentAsset struct {
	Grant                  Int   `json:"grant" gorm:"column:grant;bigint;comment:发放贷款或垫款"`
	DebtInvestment         Int   `json:"debt_investment" gorm:"column:debt_investment;bigint;comment:债权投资"`
	OtherDebtInvestment    Int   `json:"other_debt_investment" gorm:"column:other_debt_investment;bigint;comment:其他债权投资"`
	LongReceivableCash     Int   `json:"long_receivable_cash" gorm:"column:long_receivable_cash;bigint;comment:长期应收账款"`
	LongStock              Int   `json:"long_stock" gorm:"column:long_stock;bigint;comment:长期股权投资"`
	OtherEquity            Int   `json:"other_equity" gorm:"column:other_equity;bigint;comment:其他权益工具投资"`
	OtherNonCurrentFinance Int   `json:"other_non_current_finance" gorm:"column:other_non_current_finance;bigint;comment:其他非流动金融资产"`
	Estate                 Int   `json:"estate" gorm:"column:estate;bigint;comment:投资性地产"`
	FixedAsset             Int   `json:"fixed_asset" gorm:"column:fixed_asset;bigint;comment:固定资产"`
	ProgressConstruction   Int   `json:"progress_construction" gorm:"column:progress_construction;bigint;comment:在建工程"`
	UseRight               Int   `json:"use_right" gorm:"column:use_right;bigint;comment:使用权资产"`
	IntangibleAsset        Int   `json:"intangible_asset" gorm:"column:intangible_asset;bigint;comment:无形资产"`
	Goodwill               Int   `json:"goodwill" gorm:"column:goodwill;bigint;comment:商誉"`
	LongExpense            Int   `json:"long_expense" gorm:"column:long_expense;bigint;comment:长期待摊费用"`
	DeferredTax            Int   `json:"deferred_tax" gorm:"column:deferred_tax;bigint;comment:延递所得税资产"`
	OtherNonCurrentAsset   Int   `json:"other_non_current_asset" gorm:"column:other_non_current_asset;bigint;comment:其他非流动资产"`
	ExtraNonCurrentAsset   SIMap `json:"extra_non_current_asset" gorm:"column:extra_non_current_asset;charset(2048);comment:额外的非流动资产"`
}

func (n *NonCurrentAsset) GenerateID() {
	if n.Grant != nil {
		n.Grant.ID = idgenerate.Generate()
	}
	if n.DebtInvestment != nil {
		n.DebtInvestment.ID = idgenerate.Generate()
	}
	if n.OtherDebtInvestment != nil {
		n.OtherDebtInvestment.ID = idgenerate.Generate()
	}
	if n.LongReceivableCash != nil {
		n.LongReceivableCash.ID = idgenerate.Generate()
	}
	if n.LongStock != nil {
		n.LongStock.ID = idgenerate.Generate()
	}
	if n.OtherEquity != nil {
		n.OtherEquity.ID = idgenerate.Generate()
	}
	if n.OtherNonCurrentFinance != nil {
		n.OtherNonCurrentFinance.ID = idgenerate.Generate()
	}
	if n.Estate != nil {
		n.Estate.ID = idgenerate.Generate()
	}
	if n.FixedAsset != nil {
		n.FixedAsset.ID = idgenerate.Generate()
	}
	if n.ProgressConstruction != nil {
		n.ProgressConstruction.ID = idgenerate.Generate()
	}
	if n.UseRight != nil {
		n.UseRight.ID = idgenerate.Generate()
	}
	if n.IntangibleAsset != nil {
		n.IntangibleAsset.ID = idgenerate.Generate()
	}
	if n.Goodwill != nil {
		n.Goodwill.ID = idgenerate.Generate()
	}
	if n.LongExpense != nil {
		n.LongExpense.ID = idgenerate.Generate()
	}
	if n.DeferredTax != nil {
		n.DeferredTax.ID = idgenerate.Generate()
	}
	if n.OtherNonCurrentAsset != nil {
		n.OtherNonCurrentAsset.ID = idgenerate.Generate()
	}

}

type CurrentLiability struct {
	ShortLoan                    Int   `json:"short_loan" gorm:"column:short_loan;bigint;comment:短期借款"`
	Deposit                      Int   `json:"deposit" gorm:"column:deposit;bigint;comment:吸收存款及同行存放"`
	BorrowedFund                 Int   `json:"borrowed_fund" gorm:"column:borrowed_fund;bigint;comment:拆入资金"`
	DerivativeFinancialLiability Int   `json:"derivative_financial_liability" gorm:"column:derivative_financial_liability;bigint;comment:衍生金融负债"`
	PayableBill                  Int   `json:"payable_bill" gorm:"column:payable_bill;bigint;comment:应付票据"`
	PayableCash                  Int   `json:"payable_cash" gorm:"column:payable_cash;bigint;comment:应付账款"`
	AdvancesReceived             Int   `json:"advances_received" gorm:"column:advances_received;bigint;comment:预收款项"`
	ContractualLiability         Int   `json:"contractual_liability" gorm:"column:contractual_liability;bigint;comment:合同负债"`
	EmployeeCompensation         Int   `json:"employee_compensation" gorm:"column:employee_compensation;bigint;comment:应付职工薪酬"`
	Taxes                        Int   `json:"taxes" gorm:"column:taxes;bigint;comment:应交税费"`
	NonCurrentLiability          Int   `json:"non_current_liability" gorm:"column:non_current_liability;bigint;comment:一年内到期非流动负债"`
	OtherCurrentLiability        Int   `json:"other_current_liability" gorm:"column:other_current_liability;bigint;comment:其他流动负债"`
	ExtraCurrentLiability        SIMap `json:"extra_current_liability" gorm:"column:extra_current_liability;charset(2048);comment:额外的流动负债"`
}

func (n *CurrentLiability) GenerateID() {
	if n.ShortLoan != nil {
		n.ShortLoan.ID = idgenerate.Generate()
	}
	if n.Deposit != nil {
		n.Deposit.ID = idgenerate.Generate()
	}
	if n.BorrowedFund != nil {
		n.BorrowedFund.ID = idgenerate.Generate()
	}
	if n.DerivativeFinancialLiability != nil {
		n.DerivativeFinancialLiability.ID = idgenerate.Generate()
	}
	if n.PayableBill != nil {
		n.PayableBill.ID = idgenerate.Generate()
	}
	if n.PayableCash != nil {
		n.PayableCash.ID = idgenerate.Generate()
	}
	if n.AdvancesReceived != nil {
		n.AdvancesReceived.ID = idgenerate.Generate()
	}
	if n.ContractualLiability != nil {
		n.ContractualLiability.ID = idgenerate.Generate()
	}
	if n.EmployeeCompensation != nil {
		n.EmployeeCompensation.ID = idgenerate.Generate()
	}
	if n.Taxes != nil {
		n.Taxes.ID = idgenerate.Generate()
	}
	if n.NonCurrentLiability != nil {
		n.NonCurrentLiability.ID = idgenerate.Generate()
	}
	if n.OtherCurrentLiability != nil {
		n.OtherCurrentLiability.ID = idgenerate.Generate()
	}
}

type NonCurrentLiability struct {
	LongLoan                 Int   `json:"long_loan" gorm:"column:long_loan;bigint;comment:长期借款"`
	LeaseLiability           Int   `json:"lease_liability" gorm:"column:lease_liability;bigint;comment:租赁负债"`
	LongPayable              Int   `json:"long_payable" gorm:"column:long_payable;bigint;comment:长期应付款"`
	LongEmployeeCompensation Int   `json:"long_employee_compensation" gorm:"column:long_employee_compensation;bigint;comment:长期应付职工薪酬"`
	DeferredRevenue          Int   `json:"deferred_revenue" gorm:"column:deferred_revenue;bigint;comment:延递收益"`
	DeferredTaxLiabilities   Int   `json:"deferred_tax_liabilities" gorm:"column:deferred_tax_liabilities;bigint;comment:递延所得税负债"`
	ExtraNonCurrentLiability SIMap `json:"extra_non_current_liability" gorm:"column:extra_non_current_liability;charset(2048);comment:额外的非流动负债"`
}

func (n *NonCurrentLiability) GenerateID() {
	if n.LongLoan != nil {
		n.LongLoan.ID = idgenerate.Generate()
	}
	if n.LeaseLiability != nil {
		n.LeaseLiability.ID = idgenerate.Generate()
	}
	if n.LongPayable != nil {
		n.LongPayable.ID = idgenerate.Generate()
	}
	if n.LongEmployeeCompensation != nil {
		n.LongEmployeeCompensation.ID = idgenerate.Generate()
	}
	if n.DeferredRevenue != nil {
		n.DeferredRevenue.ID = idgenerate.Generate()
	}
	if n.DeferredTaxLiabilities != nil {
		n.DeferredTaxLiabilities.ID = idgenerate.Generate()
	}
}
