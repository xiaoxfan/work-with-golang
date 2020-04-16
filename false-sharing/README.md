### 在看sync.Pool源码时发现poolLocal结构体内部有一个pad变量

```go
type poolLocal struct {
	poolLocalInternal

	// Prevents false sharing on widespread platforms with
	// 128 mod (cache line size) = 0 .
	pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte
}
```

pad是为了防止在多核CPU下的遇到缓存的false sharing问题。

- 缓存 false sharing：一个 CPU 核更新变量会强制其他 CPU 核更新缓存。而我们都知道从缓存中读取 CPU 的变量比从内存中读取变量要快得多。因此，虽然该变量一直存在于多核中，但这会显著影响性能。
解决该问题的常用方法是缓存填充：在变量之间填充一些无意义的变量。使一个变量单独占用 CPU 核的缓存行，因此当其他核更新时，其他变量不会使该核从内存中重新加载变量。

具体参考: 
-  [原文](https://medium.com/@genchilu/whats-false-sharing-and-how-to-solve-it-using-golang-as-example-ef978a305e10) 
-  [译文](https://juejin.im/post/5d0519e05188257a78764d5d?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com#comment)
         