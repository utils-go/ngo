# NGo
`ngo` means .net api implement by go

the .net api implement by go. you can write go as like .net

this is based on .net4.7.2 api

api reference:https://learn.microsoft.com/en-us/dotnet/api/?view=netframework-4.7.2&preserve-view=true

source code reference:https://referencesource.microsoft.com/
# Usage
 ```
 go get -u github.com/utils-go/ngo
 ```
file: same as `Sytem.IO.File`
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/io/file"
)

func main() {
	content := "this is ngo example"
	file.WriteAllText("./test.txt", content)
	result, err := file.ReadAllText("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

```
bitconverter: same as `System.BitConverter`,but it is not static methods,so you can change `LittleEndian` or `BigEndian` any where
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/bitconverter"
)

func main() {
	converter := bitconverter.BitConverter{IsLittleEndian: true} //littleEndian at the front
	value := int64(11234567489)
	bytes, err := converter.GetBytesFromInt64E(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes)
	valueNew, err := converter.ToInt64E(bytes, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value == valueNew)
}
```

strings: same as `System.String` with .NET-style string manipulation
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/stringUtils"
)

func main() {
	s := "  Hello, World!  "
	fmt.Println("Original:", s)
	fmt.Println("Trimmed:", stringUtils.Trim(s))
	fmt.Println("Upper:", stringUtils.ToUpper(s))
	fmt.Println("Contains 'World':", stringUtils.Contains(s, "World"))
	formatted := stringUtils.Format("Hello {0}, you are {1}!", "John", 25)
	fmt.Println("Formatted:", formatted)
}
```

math: same as `System.Math`
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/math"
)

func main() {
	fmt.Println("Abs(-5.5):", math.Abs(-5.5))
	fmt.Println("Max(10, 20):", math.Max(10, 20))
	fmt.Println("Pow(2, 3):", math.Pow(2, 3))
	fmt.Println("PI:", math.PI)
}
```

## New: HashSet<T>
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/collections/generic"
)

func main() {
	hs := generic.NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(2) // duplicate, returns false
	fmt.Println("Count:", hs.Count())
	fmt.Println("Contains 2:", hs.Contains(2))
	
	other := generic.NewHashSetFromSlice([]int{2, 3, 4})
	hs.UnionWith(other)
	fmt.Println("Union:", hs.ToSlice())
}
```

## New: Queue<T> / Stack<T>
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/collections/generic"
)

func main() {
	// Queue (FIFO)
	q := generic.NewQueue[string]()
	q.Enqueue("first")
	q.Enqueue("second")
	item, _ := q.Dequeue()
	fmt.Println("Dequeued:", item) // "first"
	
	// Stack (LIFO)
	s := generic.NewStack[int]()
	s.Push(10)
	s.Push(20)
	val, _ := s.Pop()
	fmt.Println("Popped:", val) // 20
}
```

## New: StreamReader / StreamWriter
```go
package main

import (
	"bytes"
	"fmt"
	"github.com/utils-go/ngo/io/stream"
)

func main() {
	// StreamWriter
	var buf bytes.Buffer
	sw := stream.NewStreamWriter(&buf)
	sw.WriteLine("Hello")
	sw.WriteLine("World")
	sw.Flush()
	
	// StreamReader
	sr := stream.NewStreamReader(bytes.NewReader(buf.Bytes()))
	line1, _ := sr.ReadLine()
	line2, _ := sr.ReadLine()
	fmt.Println(line1, line2) // Hello World
}
```

## New: MemoryStream
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/io/stream"
)

func main() {
	ms := stream.NewMemoryStream()
	ms.Write([]byte("data in memory"))
	ms.SetPosition(0)
	buf := make([]byte, 100)
	n, _ := ms.Read(buf)
	fmt.Println(string(buf[:n]))
}
```

## New: BinaryReader / BinaryWriter
```go
package main

import (
	"bytes"
	"fmt"
	"github.com/utils-go/ngo/io/stream"
)

func main() {
	var buf bytes.Buffer
	bw := stream.NewBinaryWriter(&buf)
	bw.WriteInt32(42)
	bw.WriteString("hello")
	bw.WriteBoolean(true)
	
	br := stream.NewBinaryReader(bytes.NewReader(buf.Bytes()))
	v32, _ := br.ReadInt32()
	s, _ := br.ReadString()
	b, _ := br.ReadBoolean()
	fmt.Println(v32, s, b) // 42 hello true
}
```

## New: FileStream
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/io/stream"
)

func main() {
	fs, _ := stream.NewFileStream("./test.dat", stream.FileModeCreate, stream.FileAccessReadWrite)
	defer fs.Close()
	fs.Write([]byte("filestream data"))
	fs.SetPosition(0)
	buf := make([]byte, 1024)
	n, _ := fs.Read(buf)
	fmt.Println(string(buf[:n]))
}
```

## New: Regex
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/text/regex"
)

func main() {
	// Compile and match
	r := regex.MustCompile(`\d+`)
	fmt.Println("IsMatch:", r.IsMatch("abc123def"))
	
	// Groups
	r2 := regex.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	m := r2.Match("user@example.com")
	fmt.Println("Domain:", m.Group(2).Value()) // example
	
	// Replace
	result := regex.ReplaceString("abc123def456", `\d+`, "#")
	fmt.Println(result) // abc#def#
}
```

## New: WebClient / Uri
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/net/uri"
)

func main() {
	wc := uri.NewWebClient()
	wc.SetHeader("User-Agent", "NGo/1.0")
	
	// Download
	content, err := wc.DownloadString("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response:", content[:50], "...")
	
	// Parse URI
	u, _ := uri.ParseUri("https://example.com/path?key=value#section")
	fmt.Println("Host:", u.Host())
	fmt.Println("Path:", u.Path())
	fmt.Println("Query:", u.Query())
}
```

## New: Cryptography
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/security/cryptography"
)

func main() {
	// Hashing
	md5 := cryptography.MD5HashString("hello")
	sha256 := cryptography.SHA256HashString("hello")
	fmt.Println("MD5:", md5)
	fmt.Println("SHA256:", sha256)
	
	// AES encryption
	encrypted, _ := cryptography.AESEncryptString("secret message", "my-32-byte-key-here-1234567890")
	decrypted, _ := cryptography.AESDecryptString(encrypted, "my-32-byte-key-here-1234567890")
	fmt.Println("Decrypted:", decrypted)
	
	// HMAC
	hmac := cryptography.HMACSHA256String("message", "secret-key")
	fmt.Println("HMAC:", hmac)
}
```

## New: Threading
```go
package main

import (
	"fmt"
	"time"
	"github.com/utils-go/ngo/threading"
)

func main() {
	// Mutex
	m := threading.NewMutex()
	m.WaitOne()
	// critical section...
	m.ReleaseMutex()
	
	// Semaphore (max 3 concurrent)
	sem := threading.NewSemaphore(3, 3)
	sem.WaitOne()
	defer sem.Release()
	
	// AutoResetEvent
	evt := threading.NewAutoResetEvent(false)
	go func() {
		time.Sleep(100 * time.Millisecond)
		evt.Set()
	}()
	evt.WaitOne()
	fmt.Println("Event signaled!")
}
```

## New: Timer
```go
package main

import (
	"fmt"
	"time"
	"github.com/utils-go/ngo/timer"
)

func main() {
	t := timer.NewTimer(500 * time.Millisecond)
	t.Elapsed(func() {
		fmt.Println("Tick at", time.Now().Format("15:04:05"))
	})
	t.Start()
	time.Sleep(2 * time.Second)
	t.Stop()
}
```

## New: Random
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/math/random"
)

func main() {
	rng := random.NewRandom()
	fmt.Println("Next:", rng.Next())
	fmt.Println("NextN(10):", rng.NextN(10))
	fmt.Println("NextRange(100, 200):", rng.NextRange(100, 200))
	fmt.Println("NextDouble:", rng.NextDouble())
}
```

# Complete:
- IO.File (Read/Write/Copy/Move/Append/Exists)
- IO.Path
- IO.Directory
- IO.FileInfo
- **IO.Stream** (StreamReader, StreamWriter, BinaryReader, BinaryWriter, MemoryStream, FileStream)
- BitConverter
- DateTime (Enhanced)
- TimeSpan (Enhanced)
- Console
- String (Split, Join, Format, Pad, Trim, etc.)
- Math (Abs, Max, Min, Sin, Cos, Pow, Round, etc.)
- Convert (ToInt32, ToDouble, ToBoolean, ToString, Base64, Hex, etc.)
- Collections.Generic.List<T>
- **Collections.Generic.HashSet<T>**
- **Collections.Generic.SortedList<TKey,TValue>**
- **Collections.Generic.SortedDictionary<TKey,TValue>**
- **Collections.Generic.Queue<T>**
- **Collections.Generic.Stack<T>**
- **Collections.Generic.LinkedList<T>**
- Collections.Generic.Dictionary<K,V>
- Linq (Where, Select, OrderBy, GroupBy, Aggregate, etc.)
- Environment
- Diagnostics.Stopwatch
- Text.Encoding (UTF-8, ASCII, UTF-16, UTF-32)
- **Text.RegularExpressions.Regex**
- Text.StringBuilder
- Reflection
- **Net.WebClient / Net.Uri**
- **Security.Cryptography (MD5, SHA1, SHA256, AES, HMAC)**
- **Threading (Mutex, Semaphore, AutoResetEvent, ManualResetEvent)**
- **Timer (System.Timers.Timer)**
- **Math.Random**
