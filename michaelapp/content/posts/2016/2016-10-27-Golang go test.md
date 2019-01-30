---
title: "Golang go test"
date: 2016-10-27T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---

源码地址：[https://github.com/panyingyun/ducktest](https://github.com/panyingyun/ducktest)

官方参考：[https://blog.golang.org/cover](https://blog.golang.org/cover)


## Goang之go test


go test 包括代码单元测试、性能测试、用例测试、覆盖率、性能分析等5个方面的全面测试框架

usage: go test [build/test flags] [packages] [build/test flags & test binary flags]

### 1. 单元测试

参考 雨痕 《Go 学习笔记 第四版.pdf》

```
参数      		   说明
-c                  仅编译，不执行测试
-v                  显示所有测试函数的执行细节
-run regex          执行指定的测试函数（通过正则表达式匹配）   
-parallel n         并发执行测试函数（默认：GOMAXPROCS）
-timeout  t         单个测试超时时间。（-timeout 2m30s）
```
	
```shell
#go test
PASS
ok      github.com/panyingyun/ducktest  0.434s

#
#go test -v
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
PASS
ok      github.com/panyingyun/ducktest  0.323s

#go test -v -run TestSum
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
ok      github.com/panyingyun/ducktest  0.342s

#go test -v -timeout 3s
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
PASS
ok      github.com/panyingyun/ducktest  0.197s
```

### 2. Benchmark

```shell
参数   
-bench regex   执行指定的性能测试函数（通过正则表达式匹配）
-benchmen      输出内存统计信息
-benchtime -t  设置每个性能测试运行时间   -benchtime 1m20s            
-cpu           设置并发测试。默认 GOMAXPROCS  -cpu 1,2,4,8
```

```shell
#go test -v -bench .
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
BenchmarkSum-8          2000000000               0.28 ns/op
BenchmarkMul-8          2000000000               0.27 ns/op
PASS
ok      github.com/panyingyun/ducktest  1.481s

#go test -v -bench . -benchmem -cpu 1,2,4,8 -benchtime 30s
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
BenchmarkSum            2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkSum-2          2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkSum-4          2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkSum-8          2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkMul            2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkMul-2          2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkMul-4          2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkMul-8          2000000000               0.27 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/panyingyun/ducktest  4.810s

```

### 3. Example

```shell
#go test -v 
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMul
--- PASS: TestMul (0.00s)
=== RUN   ExampleSum
--- PASS: ExampleSum (0.00s)
=== RUN   ExampleMul
--- PASS: ExampleMul (0.00s)
PASS
ok      github.com/panyingyun/ducktest  0.196s
```

Example 代码可输出到文档

### 4.Cover

```shell
参数  
-cover        允许覆盖分析
-covermode    代码分析模式。（set: 是否执行;  count:执行次数; atomic:次数，并发支持)
-coverprofile 输出结果文件。
```

```shell
//输出测试覆盖比例和覆盖测试文件
#go test -cover -coverprofile=cover.out -covermode=count
PASS
coverage: 50.0% of statements
ok      github.com/panyingyun/ducktest  0.132s

//覆盖测试结果
分析
#go tool cover -func=cover.out
github.com\panyingyun\ducktest\duck.go:7:       sum             100.0%
github.com\panyingyun\ducktest\duck.go:11:      mul             100.0%
github.com\panyingyun\ducktest\duck.go:15:      main            0.0%
total:                                          (statements)    50.0%

//浏览器显示
#go tool cover -html=cover.out


```

### PProf
监控程序执行，找出性能瓶颈，以encoding/json包为例子
```shell
//产生文件 cpu.out mem.out  json.test.exe

#go test -v -test.bench . -cpuprofile cpu.out -memprofile mem.out encoding/json
=== RUN   TestMarshal
--- PASS: TestMarshal (0.00s)
=== RUN   TestMarshalBadUTF8
--- PASS: TestMarshalBadUTF8 (0.00s)
=== RUN   TestMarshalNumberZeroVal
--- PASS: TestMarshalNumberZeroVal (0.00s)
=== RUN   TestMarshalEmbeds
--- PASS: TestMarshalEmbeds (0.00s)
=== RUN   TestUnmarshal
--- PASS: TestUnmarshal (0.00s)
=== RUN   TestUnmarshalMarshal
--- PASS: TestUnmarshalMarshal (0.30s)
=== RUN   TestNumberAccessors
--- PASS: TestNumberAccessors (0.00s)
=== RUN   TestLargeByteSlice
--- PASS: TestLargeByteSlice (0.00s)
=== RUN   TestUnmarshalInterface
--- PASS: TestUnmarshalInterface (0.00s)
=== RUN   TestUnmarshalPtrPtr
--- PASS: TestUnmarshalPtrPtr (0.00s)
=== RUN   TestEscape
--- PASS: TestEscape (0.00s)
=== RUN   TestErrorMessageFromMisusedString
--- PASS: TestErrorMessageFromMisusedString (0.00s)
=== RUN   TestRefUnmarshal
--- PASS: TestRefUnmarshal (0.00s)
=== RUN   TestEmptyString
--- PASS: TestEmptyString (0.00s)
=== RUN   TestNullString
--- PASS: TestNullString (0.00s)
=== RUN   TestInterfaceSet
--- PASS: TestInterfaceSet (0.00s)
=== RUN   TestUnmarshalNulls
--- PASS: TestUnmarshalNulls (0.00s)
=== RUN   TestStringKind
--- PASS: TestStringKind (0.00s)
=== RUN   TestByteKind
--- PASS: TestByteKind (0.00s)
=== RUN   TestSliceOfCustomByte
--- PASS: TestSliceOfCustomByte (0.00s)
=== RUN   TestUnmarshalTypeError
--- PASS: TestUnmarshalTypeError (0.00s)
=== RUN   TestUnmarshalSyntax
--- PASS: TestUnmarshalSyntax (0.00s)
=== RUN   TestUnmarshalUnexported
--- PASS: TestUnmarshalUnexported (0.00s)
=== RUN   TestUnmarshalJSONLiteralError
--- PASS: TestUnmarshalJSONLiteralError (0.00s)
=== RUN   TestSkipArrayObjects
--- PASS: TestSkipArrayObjects (0.00s)
=== RUN   TestPrefilled
--- PASS: TestPrefilled (0.00s)
=== RUN   TestInvalidUnmarshal
--- PASS: TestInvalidUnmarshal (0.00s)
=== RUN   TestInvalidUnmarshalText
--- PASS: TestInvalidUnmarshalText (0.00s)
=== RUN   TestInvalidStringOption
--- PASS: TestInvalidStringOption (0.00s)
=== RUN   TestOmitEmpty
--- PASS: TestOmitEmpty (0.00s)
=== RUN   TestStringTag
--- PASS: TestStringTag (0.00s)
=== RUN   TestEncodeRenamedByteSlice
--- PASS: TestEncodeRenamedByteSlice (0.00s)
=== RUN   TestUnsupportedValues
--- PASS: TestUnsupportedValues (0.00s)
=== RUN   TestRefValMarshal
--- PASS: TestRefValMarshal (0.00s)
=== RUN   TestMarshalerEscaping
--- PASS: TestMarshalerEscaping (0.00s)
=== RUN   TestAnonymousNonstruct
--- PASS: TestAnonymousNonstruct (0.00s)
=== RUN   TestEmbeddedBug
--- PASS: TestEmbeddedBug (0.00s)
=== RUN   TestTaggedFieldDominates
--- PASS: TestTaggedFieldDominates (0.00s)
=== RUN   TestDuplicatedFieldDisappears
--- PASS: TestDuplicatedFieldDisappears (0.00s)
=== RUN   TestStringBytes
--- PASS: TestStringBytes (0.10s)
=== RUN   TestIssue6458
--- PASS: TestIssue6458 (0.00s)
=== RUN   TestIssue10281
--- PASS: TestIssue10281 (0.00s)
=== RUN   TestHTMLEscape
--- PASS: TestHTMLEscape (0.00s)
=== RUN   TestEncodePointerString
--- PASS: TestEncodePointerString (0.00s)
=== RUN   TestEncodeString
--- PASS: TestEncodeString (0.00s)
=== RUN   TestEncodeBytekind
--- PASS: TestEncodeBytekind (0.00s)
=== RUN   TestTextMarshalerMapKeysAreSorted
--- PASS: TestTextMarshalerMapKeysAreSorted (0.00s)
=== RUN   TestFold
--- PASS: TestFold (0.00s)
=== RUN   TestFoldAgainstUnicode
--- PASS: TestFoldAgainstUnicode (0.00s)
=== RUN   TestNumberIsValid
--- PASS: TestNumberIsValid (0.00s)
=== RUN   TestCompact
--- PASS: TestCompact (0.00s)
=== RUN   TestCompactSeparators
--- PASS: TestCompactSeparators (0.00s)
=== RUN   TestIndent
--- PASS: TestIndent (0.00s)
=== RUN   TestCompactBig
--- PASS: TestCompactBig (0.20s)
=== RUN   TestIndentBig
--- PASS: TestIndentBig (0.44s)
=== RUN   TestIndentErrors
--- PASS: TestIndentErrors (0.00s)
=== RUN   TestNextValueBig
--- PASS: TestNextValueBig (0.17s)
=== RUN   TestEncoder
--- PASS: TestEncoder (0.00s)
=== RUN   TestEncoderIndent
--- PASS: TestEncoderIndent (0.00s)
=== RUN   TestEncoderSetEscapeHTML
--- PASS: TestEncoderSetEscapeHTML (0.00s)
=== RUN   TestDecoder
--- PASS: TestDecoder (0.00s)
=== RUN   TestDecoderBuffered
--- PASS: TestDecoderBuffered (0.00s)
=== RUN   TestRawMessage
--- PASS: TestRawMessage (0.00s)
=== RUN   TestNullRawMessage
--- PASS: TestNullRawMessage (0.00s)
=== RUN   TestBlocking
--- PASS: TestBlocking (0.00s)
=== RUN   TestDecodeInStream
--- PASS: TestDecodeInStream (0.00s)
=== RUN   TestHTTPDecoding
--- PASS: TestHTTPDecoding (0.00s)
=== RUN   TestStructTagObjectKey
--- PASS: TestStructTagObjectKey (0.00s)
=== RUN   TestTagParsing
--- PASS: TestTagParsing (0.00s)
=== RUN   ExampleMarshal
--- PASS: ExampleMarshal (0.00s)
=== RUN   ExampleUnmarshal
--- PASS: ExampleUnmarshal (0.00s)
=== RUN   ExampleDecoder
--- PASS: ExampleDecoder (0.00s)
=== RUN   ExampleDecoder_Token
--- PASS: ExampleDecoder_Token (0.00s)
=== RUN   ExampleDecoder_Decode_stream
--- PASS: ExampleDecoder_Decode_stream (0.00s)
=== RUN   ExampleRawMessage
--- PASS: ExampleRawMessage (0.00s)
=== RUN   ExampleIndent
--- PASS: ExampleIndent (0.00s)
BenchmarkCodeEncoder-8                       100          14411441 ns/op         134.65 MB/s
BenchmarkCodeMarshal-8                       100          16111611 ns/op         120.44 MB/s
BenchmarkCodeDecoder-8                        30          43971063 ns/op          44.13 MB/s
BenchmarkDecoderStream-8                 5000000               372 ns/op
BenchmarkCodeUnmarshal-8                      30          44804480 ns/op          43.31 MB/s
BenchmarkCodeUnmarshalReuse-8                 30          42670933 ns/op
BenchmarkUnmarshalString-8               3000000               469 ns/op
BenchmarkUnmarshalFloat64-8              3000000               426 ns/op
BenchmarkUnmarshalInt64-8                5000000               383 ns/op
BenchmarkIssue10335-8                    3000000               549 ns/op             272 B/op          4 allocs/op
BenchmarkNumberIsValid-8                100000000               16.6 ns/op
BenchmarkNumberIsValidRegexp-8           2000000               625 ns/op
BenchmarkSkipValue-8                         100          12701270 ns/op         157.50 MB/s
BenchmarkEncoderEncode-8                 3000000               482 ns/op               8 B/op          1 allocs/op
PASS
ok      encoding/json   26.536s

//交互模式分析结果

#go tool pprof json.test.exe mem.out
Entering interactive mode (type "help" for commands)
(pprof) top
6801.15kB of 6801.15kB total (  100%)
Dropped 83 nodes (cum <= 34kB)
Showing top 10 nodes out of 30 (cum >= 2560.66kB)
      flat  flat%   sum%        cum   cum%
 4240.49kB 62.35% 62.35%  4240.49kB 62.35%  bytes.makeSlice
 1536.05kB 22.59% 84.93%  1536.05kB 22.59%  encoding/json.(*decodeState).literalStore
  512.56kB  7.54% 92.47%   512.56kB  7.54%  reflect.unsafe_NewArray
  512.04kB  7.53%   100%   512.04kB  7.53%  reflect.unsafe_New
         0     0%   100%  2086.21kB 30.67%  bytes.(*Buffer).ReadFrom
         0     0%   100%  2154.28kB 31.68%  bytes.(*Buffer).WriteString
         0     0%   100%  2154.28kB 31.68%  bytes.(*Buffer).grow
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*arrayEncoder).(encoding/json.encode)-fm
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*arrayEncoder).encode
         0     0%   100%  2560.66kB 37.65%  encoding/json.(*decodeState).array
(pprof)


//文本输出结果和web输出结果
#go tool pprof -text json.test.exe mem.out
6801.15kB of 6801.15kB total (  100%)
Dropped 83 nodes (cum <= 34kB)
      flat  flat%   sum%        cum   cum%
 4240.49kB 62.35% 62.35%  4240.49kB 62.35%  bytes.makeSlice
 1536.05kB 22.59% 84.93%  1536.05kB 22.59%  encoding/json.(*decodeState).literalStore
  512.56kB  7.54% 92.47%   512.56kB  7.54%  reflect.unsafe_NewArray
  512.04kB  7.53%   100%   512.04kB  7.53%  reflect.unsafe_New
         0     0%   100%  2086.21kB 30.67%  bytes.(*Buffer).ReadFrom
         0     0%   100%  2154.28kB 31.68%  bytes.(*Buffer).WriteString
         0     0%   100%  2154.28kB 31.68%  bytes.(*Buffer).grow
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*arrayEncoder).(encoding/json.encode)-fm
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*arrayEncoder).encode
         0     0%   100%  2560.66kB 37.65%  encoding/json.(*decodeState).array
         0     0%   100%   512.04kB  7.53%  encoding/json.(*decodeState).indirect
         0     0%   100%  1536.05kB 22.59%  encoding/json.(*decodeState).literal
         0     0%   100%  2560.66kB 37.65%  encoding/json.(*decodeState).object
         0     0%   100%  2560.66kB 37.65%  encoding/json.(*decodeState).value
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*encodeState).reflectValue
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*encodeState).string
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*mapEncoder).(encoding/json.encode)-fm
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*mapEncoder).encode
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*sliceEncoder).(encoding/json.encode)-fm
         0     0%   100%  2154.28kB 31.68%  encoding/json.(*sliceEncoder).encode
         0     0%   100%  2086.21kB 30.67%  encoding/json.BenchmarkCodeEncoder
         0     0%   100%  2086.21kB 30.67%  encoding/json.codeInit
         0     0%   100%  2154.28kB 31.68%  encoding/json.interfaceEncoder
         0     0%   100%  2086.21kB 30.67%  io/ioutil.ReadAll
         0     0%   100%  2086.21kB 30.67%  io/ioutil.readAll
         0     0%   100%   512.56kB  7.54%  reflect.MakeSlice
         0     0%   100%   512.04kB  7.53%  reflect.New
         0     0%   100%  2086.21kB 30.67%  runtime.goexit
         0     0%   100%  2086.21kB 30.67%  testing.(*B).run1.func1
         0     0%   100%  2086.21kB 30.67%  testing.(*B).runN

#go tool pprof -web json.test.exe mem.out
Cannot find dot, have you installed Graphviz?

提示需要按装 Graphviz 安装一下即可
```

另外 对于网络服务器还可以使用
net/http/pprof 实时查看 runtime profiling 信息。

```go
package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe("localhost:6060", nil)
	for {
		time.Sleep(time.Second)
	}
}
```
在浏览器中查看 http://localhost:6060/debug/pprof/

