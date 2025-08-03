package types

type UCASID struct {
	// 国科大学号 / 职工号
	ID *string `json:"id"`
	// SEP 身份类型，分为本科生 / 硕士研究生 / 博士研究生 / 职工
	Type *string `json:"type"`
}

type User struct {
	Uid        *int64   `json:"uid"`
	Name       *string  `json:"name"`
	Email      *string  `json:"email"`
	UCASIDList []UCASID `json:"ucas_id_list"`
}
