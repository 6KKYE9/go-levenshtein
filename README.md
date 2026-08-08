# go-levenshtein

零依赖的莱文斯坦（编辑）距离计算工具，支持中文等多字节字符。

## 用法

```bash
go-levenshtein kitten sitting   # 输出 3
```

输出把第一个字符串变成第二个最少需要改的字符数（增 / 删 / 替各算 1）。
库函数 `Distance(a, b string) int` 可直接调用。
