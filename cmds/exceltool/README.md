# config

## 目录

- excel 配置文件
- json 解析出来的json文件
- go golang文件用于读取json配置

## 解析规则:

- 第一行=字段描述
- 第二行=字段名称(驼峰命名)
- 第三行=字段类型 (autocode.go)
    - i = 数字`int32`
    - t = 字典object配置,对应字段名称为sheet名称,对应id数据解析到此字段内

## 结果
```json
{
  "101": {
    "id": 101,
    "name": "101",
    "strAry": [
      "1",
      "2"
    ],
    "num": 0,
    "canSee": 0,
    "moreAry": [
      [
        20000,
        1000
      ]
    ],
    "obj": {
      "201": {
        "productionId": 201,
        "amount": 11
      }
    }
  }
}
```
