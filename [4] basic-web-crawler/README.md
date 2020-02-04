# Floating Prime :rocket:

## Description
ทีม Marketing โครงการ Internship ต้องการทํา API สําหรับ Logo บริษัทที่เข้ามาร่วมในโครงการ แต่เนื่องจากไม่
อยากรบกวนทีม Developer ของโครงการ (ที่แต่ละคนเป็นอาสาสมัครจากบริษัทที่เข้าร่วมโครงการ ที่งานล้นมืออยู่
แล้ว) จึงคิดวิธีการง่ายๆ ที่จะ “ดึงเฉพาะ Logo URL มาจากหน้าเว็บของโครงการ The Internship” และขอให้น้องๆ
ช่วยเขียนโปรแกรมนี้

โจทย์ข้อนี้มี 2 ส่วน โดยที่น้องๆ สามารถทําเฉพาะส่วนแรก (4.1) ก็ได้แต่ถ้าอยากได้คะแนนเพิ่ม ก็ทําส่วนที่สอง (4.2)
ด้วยจะเยี่ยมมากเลย ..... แต่ส่วนที่สองจะทําได้ต้องทําส่วนแรกได้ก่อนนะ :-)

## Problem [4.1] :question:
1. เขียนโปรแกรมเพื่อดึงข้อมูล URL ของ Logo ทุกบริษัท ที่อยู่ในส่วน “ใครมาบ้าง / PARTICIPATING STARTUPS” ในเว็บไซต์ The Internship (https://theinternship.io/) โดยข้อมูลต้นทางจะเป็น HTML และมีข้อมูลที่เป็น URL ของ Logo แต่ละบริษัทอยู่

2. ให้ Output ข้อมูลโดยเรียงลําดับตามความยาวของ “คําอธิบายบริษัท” ซึ่งอยู่ใต้รูป โดยเรียงจากน้อยไปมาก

3. ถ้ามีการเปลี่ยนแปลงจํานวนของบริษัทบนหน้าเว็บ (มีบริษัทเข้ามาใหม่ หรือ ออกจากโครงการ) เมื่อรันโปรแกรมใหม่ผลลัพธ์จะต้องเปลี่ยนไปตามข้อมูลบนหน้าเว็บจริง

### Example Output
company/wisible_logo.png <br>
company/codeapp_logo.png <br>
company/horganice_logo.png <br>
...


## Problem [4.2] (optional) :question:
ให้น้องๆ เอาข้อมูลที่ได้มาจากข้อ 4.1 มาเปิดเป็น JSON API สําหรับเข้าถึงข้อมูล โดยมี Spec ดังนี้
- มี Route คือ /companies และรองรับเฉพาะ GET Request
- เมื่อเรียก API แล้วจะได้ JSON ที่มีโครงสร้างดังตัวอย่าง

```
{ 
    "companies" : [
        { "logo": "https://theinternship.io/company/wisible_logo.png"},
        { "logo": "https://theinternship.io/company/codeapp_logo.png"}
    ]
}
```

โดยที่น้องๆ สามารถใช้ภาษาโปรแกรมไหนก็ได้, Framework ไหนก็ได้ที่ถนัด หรืออยากใช้ :-)

## How to Run :confused:
- Clone or Download file from This repo.
- Open **"[4] basic-web-crawler"** folder.
- Run this commad.
    ```
    $ npm install
    ```
    for 4.1
    ```
    $ npm run start1 
    ```
    for 4.2 (it run on port 3000)
    ```
    $ npm run start2
    ```
