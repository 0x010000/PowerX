package heal

import (
	"PowerX/pkg/idx"
	"fmt"
)

//
// Oid
//  @Description: 档案ID
//  @return string
//
func Oid() string {
	return fmt.Sprintf(`O%s`, idx.SnowFlake.Sid())
}

//
// Cid
//  @Description: 方案ID
//  @return string
//
func Cid() string {
	return fmt.Sprintf(`C%s`, idx.SnowFlake.Sid())
}

//
// Pid
//  @Description: 评估ID
//  @return string
//
func Pid() string {
	return fmt.Sprintf(`P%s`, idx.SnowFlake.Sid())
}
