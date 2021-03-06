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

* **URL Params**

  **Required:**

  `student_id=[string]`

* **Data Params**

  ```
  {
    student_id : "student-id",
    alias_id : "alias-id",
    name : "name"
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

* **URL Params**

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
    const student_id = "student-id"
  
    fetch('http://localhost:8080/members/' + student_id, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
    })
  ```

**入退室状況の保存**
----

* **URL**

  /accesses/:student_id

* **Method:**

  `PUT`

* **URL Params**

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
        method: 'PUT',
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

* **URL Params**

  **Required:**

  `student_id=[string]`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** `{ student_id : "student-id", alias-id : "alias-id", name : "name" }`

* **Error Response:**

    * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const student_id = "student-id"
  
    fetch('http://localhost:8080/accesses/' + student_id, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
    })
  ```

**ログの記録**
----

* **URL**

  /logs

* **Method:**

  `POST`

* **URL Params**
  `None`

* **Data Params**

  {
    student_id : "student-id",
    event : "event",
    date : "20/04/23-12:02:22"
  }

* **Success Response:**

    * **Code:** 200 <br />

* **Error Response:**

    * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    const data = {
      student_id : "student-id",
      event : "event",
      date : "20/04/23-12:02:22"
    }
  
    fetch('http://localhost:8080/logs', {
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

  /logs

* **Method:**

  `GET`

* **URL Params**
  `None`

* **Data Params**
  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** ```[
        {id: "0", student_id: "e1n20021", event: "register", date: "20/04/21-01:14:33"}, 
        {id: "1", student_id: "e1n20021", event: "entry",    date: "20/04/21-01:14:40"}
      ]```

* **Error Response:**

    * undefined [https://github.com/ShebangDog/MemberRegisterApi/issues/2]

* **Sample Call:**

  ```
    fetch('http://localhost:8080/logs', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })
  ```

API Doc Sample:[https://gist.github.com/iros/3426278]
