# go-levenshtein

一堆文本要处理，开编辑器点来点去太慢，命令行一把梭最省事。

## 用法

```bash
go-levenshtein kitten sitting   # 输出 3
```

输出把第一个字符串变成第二个最少需要改的字符数（增 / 删 / 替各算 1）。
库函数 `Distance(a, b string) int` 可直接调用。
