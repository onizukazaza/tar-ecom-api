package exception

import "fmt"

type FailedToCreateAddress struct {
    UserID string
    Reason string
}

func (e *FailedToCreateAddress) Error() string {
    return fmt.Sprintf("failed to create address for user '%s': %s", e.UserID, e.Reason)  
}

// กรณีที่เกิดข้อความ 
// User ไม่มีอยู่ในฐานข้อมูล
// ID ที่ส่งมาไม่ถูกต้อง
// ฐานข้อมูลมีปัญหา
// Query หรือการเชื่อมต่อฐานข้อมูลผิดพลาด