---
title: 7-原生sql

date: 2021-12-12	

categories: GORM	

tags: [Go知识点,GORM]
---	

# Sql构建

### 执行原生SQL

```
    db.Exec("DROP TABLE users;")
    db.Exec("UPDATE orders SET shipped_at=? WHERE id IN (?)", time.Now, []int64{11,22,33})

    // Scan
    type Result struct {
        Name string
        Age  int
    }

    var result Result
    db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
```

### sql.Row & sql.Rows

获取查询结果为`*sql.Row`或`*sql.Rows`

```
    row := db.Table("users").Where("name = ?", "jinzhu").Select("name, age").Row() // (*sql.Row)
    row.Scan(&name, &age)

    rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
    defer rows.Close()
    for rows.Next() {
        ...
        rows.Scan(&name, &age, &email)
        ...
    }

    // Raw SQL
    rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows() // (*sql.Rows, error)
    defer rows.Close()
    for rows.Next() {
        ...
        rows.Scan(&name, &age, &email)
        ...
    }
```

### 迭代中使用sql.Rows的Scan

```
    rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
    defer rows.Close()

    for rows.Next() {
      var user User
      db.ScanRows(rows, &user)
      // do something
    }
```