Ecomerce Build By Golang 
- Concept for me Buildonce deploy multiple
- Clean Artitecture
- SqlX
- Postgresql
- Go Fiber
Go Module 
 - viper
   ใช้ในการโหลด package config 
   ข้อดี ลดความซับซ้อนของตัวโปรเจ็ค
 - validator  
   ใช้ในการvalidate config field ในไม่ใส่มาให้ error ไปเลย
   !ป้องกันช่องโหว sql inject xss!
          

 * ทำไมถึงแยกไฟล์เช่น repository กับ repositoryimpl
  - เพื่อความเป็นระเบียบ
  - เผื่อบางทีเรา implement หลายๆ struct
  - ทำไมถึงมี interface เผื่อเวลากาติดค่อข้อมูลถ้าจะล่วงลูกหรือข้อมูลที่ implement  ไป ต้องผ่านไฟล์ interface เพื่อ inject ++

 *yaml: ................................
  -   allowOrigins "*" ทุก domain สามารถใช้ได้หมด
  -   timeout 30 คือ ถ้ามี client ยิง api มาเกิน 30 วิจะตัด connection อัตโนมัติ

Error :
   


  // connect database !!
  conf := config.ConfigGetting()  ใน main.go รับมาจาก master file config  ฟังชั่น configgetting()