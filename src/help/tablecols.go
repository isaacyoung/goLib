package main

import (
	"fmt"
	"strings"
)

func main() {
	cols := `bill_no varchar(20) null comment '票据编号',
	school_code varchar(50) null comment '学校编码',
	school_name varchar(100) null comment '学校名称',
	order_id bigint null comment '关联订单id',
	payment_unit varchar(100) null comment '交款单位/个人',
	id_card varchar(30) null comment '身份证号',
	user_class varchar(30) null comment '班级',
	receiving_unit varchar(100) null comment '收款单位',
	receiving_person varchar(50) null comment '收款人',
	payment_amount decimal(10,2) null comment '收款金额',
	payment_amount_zn varchar(50) null comment '收款金额中文',
	remark varchar(300) null comment '备注',
	delete_flag smallint(1) null comment '删除标识 0正常 1删除',
	billing_date datetime null comment '开票日期',
	create_time datetime null comment '创建日期',`

	arrStr := strings.Split(cols, "\n")
	for _, col := range arrStr {
		rows := strings.Split(col, " ")
		colName := strings.Trim(rows[0], "\t")
		subName := strings.Split(colName, "_")
		var javaName = ""
		if len(subName) == 1 {
			javaName = colName
		} else {
			for i, sub := range subName {
				if i == 0 {
					javaName += sub
				} else {
					javaName += strings.ToUpper(sub[0:1]) + sub[1:]
				}
			}
		}

		colComment := rows[4]
		fmt.Printf("%-20s %-20s %-20s\n", colName, javaName, colComment)
	}

}
