# MemberRegisterApi

## 概要
研究室の入退室を管理するAPIです

## 機能
- ユーザーの追加/削除
- 入退室状況の保存/取得
- ログの記録/取得

**ユーザーの追加**
----

* **URL**

  /members

* **Method:**

  `POST`
  
*  **URL Params**

   **Required:**
 
   `student_id=[string]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const data = {
        student_id : "student-id"
        alias_id : "felica-idm"
        name : "name"
    }
  
    fetch('http://localhost:8080/members', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
  ```
  
  
**ユーザーの削除**
----

* **URL**

  /members/:student_id

* **Method:**

  `DELETE`
  
*  **URL Params**

   **Required:**
 
   `student_id=[string]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const data = {
        student_id : "student-id"
        alias_id : "felica-idm"
        name : "name"
    }
  
    fetch('http://localhost:8080/members/' + data.student_id, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
  ```

 **入退室状況の保存**
----

* **URL**

  /accesses/:student_id

* **Method:**

  `PUT`
  
*  **URL Params**

   **Required:**
 
   `student_id=[string]`

* **Data Params**

  ```
  {
    student_id : "student-id"
    event : "event"
  }
  ```

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const data = {
        student_id : "student-id"
        event : "event"
    }
  
    fetch('http://localhost:8080/accesses/' + data.student_id, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
  ```
  
  
**入退室状況の取得**
----

* **URL**

  /accesses/:student_id

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `student_id=[string]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const data = {
        student_id : "student-id"
        alias_id : "felica-idm"
        name : "name"
    }
  
    fetch('http://localhost:8080', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
  ```
  
  
 **ログの記録**
----

* **URL**

  /members

* **Method:**

  `POST`
  
*  **URL Params**

   **Required:**
 
   `student_id=[string]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const data = {
        student_id : "student-id"
        alias_id : "felica-idm"
        name : "name"
    }
  
    fetch('http://localhost:8080', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
  ```
  
  
  
**ログの取得**
----

* **URL**

  /members

* **Method:**

  `POST`
  
*  **URL Params**

   **Required:**
 
   `student_id=[string]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const data = {
        student_id : "student-id"
        alias_id : "felica-idm"
        name : "name"
    }
  
    fetch('http://localhost:8080', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
  ```
  
 
 API Doc Sample:[https://gist.github.com/iros/3426278]
