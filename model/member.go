package model

type Member struct {
	MemberNo int64 			`db:"member_no" json:"MemberNo" form:"MemberNo"`
	MemberName string 		`db:"member_name" json:"MemberName" form:"MemberName"`
	MemberId string 		`db:"member_id" json:"MemberId" form:"MemberId"`
	MemberRegDate string	`db:"member_reg_date" json:"MemberRegDate" form:"MemberRegDate"`
	MemberUpdDate string	`db:"member_upd_date" json:"MemberUpdDate" form:"MemberUpdDate"`
}

type InsMember struct {
	MemberNo int64 			`db:"member_no"`
	MemberName string 		`db:"member_id" json:"_MemberId" form:"_MemberId"`
	MemberId string 		`db:"member_name" json:"_MemberName" form:"_MemberName"`
/*	MemberRegDate string	`db:"member_reg_date"`
	MemberUpdDate string	`db:"member_upd_date"`*/
}

/*type MemberInfo struct {
	MemberName string 		`db:"member_name"`
	MemberId string 		`db:"member_id"`
}*/


func NewMember(memNm, memId string) InsMember {
	return InsMember{
		MemberName: memNm,
		MemberId:    memId,
	}
}