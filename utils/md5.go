package utils

import (
	"crypto/md5"
	"fmt"
)

const Salt1, Salt2 = "哿", "嬿" // 自定义的盐

func Md5(pwd string) string {
	pwdaddSalt := Salt1 + pwd + Salt2 // 在原密码的基础上进行加盐，在密码中添加一些复杂的自定义字符，也就是盐，使密码很复杂，不常见，这样就很难破解出原密码
	data := []byte(pwdaddSalt)        // md5.Sumd的参数格式是字节切片，所以要先将原数据转换为字节切片

	fmt.Println(md5.Sum(data))                // 加密后返回一个长度16的byte字节切片
	return fmt.Sprintf("%x\n", md5.Sum(data)) // 将字节切片转换为16进制的字符串
	// 使用fmt.Sprintf+ %x 将字节转换为16进制的字符串

	// 加盐: 在密码任意固定位置插入特定的字符串，让散列后的结果和使用原始密码的散列结果不同，这种过程称之为“加盐”
}
