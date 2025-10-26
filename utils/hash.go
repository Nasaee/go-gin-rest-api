package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	/*
		[]byte(password) คือ การแปลง string ให้เป็นชุด byte เพื่อให้ฟังก์ชันพวกเข้ารหัส (เช่น bcrypt) ทำงานได้ เพราะมันไม่ได้ทำงานกับ “ตัวอักษร” โดยตรง แต่กับ “ข้อมูลดิบ” ของ string ในรูปของ byte
		ตัวอย่างให้เห็นชัด:
		password := "abc"
		fmt.Println([]byte(password)) -> [97 98 99]
		เพราะ: 'a' → 97, 'b' → 98, 'c' → 99
	*/
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
