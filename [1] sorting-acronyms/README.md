# Sorting Acronyms :cyclone:

## Problem :question:
เขียนโปรแกรมเรียงชื่อตาม `ชื่อแบบอักษรย่อ` ของชื่อเต็ม ที่รับ เข้ามาจาก Standard Input โดยมีหลักการย่อคํา ดังนี้ 
1. ชื่อเต็ม จะประกอบด้วยคําภาษาอังกฤษหลายคํา 

2. การสร้างชื่อย่อ จะใช้ตัวอักษรตัวแรกของแต่ละคําในชื่อเต็ม ในกรณีที่ตัวอักษรแรกนั้นๆ เป็นตัวอักษรพิมพ์ใหญ่ 

3. กรณีท ี่ตัวอักษรตัวแรกเป็น ตัวพิมพ์เล็ก จะไม่นำมาประกอบในชื่อย่อ

### Input 
(อ่านจาก Standard Input) บรรทัดแรกของ Input จะเป็น ตัวเลขจํานวนเต็ม N ที่แสดงถึงจํานวนชื่อที่มีใน Input ทั้งหมด และจากนั้น N บรรทัด จะเป็นชื่อแต่ละชื่อที่จะต้องทําการย่อ และนําไปเรียง

### Output
(แสดงผลใน Standard Output) ให้แสดงผลชื่อย่อ ของชื่อเต็มที่รับเข้ามา บรรทัดละ 1 ชื่อ โดยเรียงลําดับจากชื่อที่ยาวที่สุด ไปยังชื่อที่สั้นที่สุด โดยหาก มีความยาวเท่า กัน ให้เรียงตามลําดับอักษร (Alphabetical Order) 


### Test Case
| Input                        | Output |
|------------------------------|--------|
| 5 <br> the United States of America <br> The United States of America <br> Carnegie Mellon University <br> The United States Of America <br> the united states of America                         | TUSOA <br> TUSA <br> CMU <br> USA <br> A <br> <br> |
  
   
     
   
  


## How to Run :confused:
- Clone or Download file from This repo.
- Open **"[1] sorting-acronyms"** folder.
- Run this commad.
    ```
    $ go run main.go
    ```
    or
    ```
    $ ./main (for linux)
    $ main.exe (for windows)
    ```
