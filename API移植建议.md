# NGo 项目 API 移植建议

## 项目概述
NGo 是一个将 .NET API 移植到 Go 语言的项目，旨在让开发者能够像使用 .NET 一样编写 Go 代码。目前已经实现了多个核心 API，如 File、Path、BitConverter、DateTime、TimeSpan、Console、String、Math、Convert 等。

## 可移植的 API 建议

### 1. System.IO 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| StreamReader/StreamWriter | 用于读写文本流 | 高 |
| BinaryReader/BinaryWriter | 用于读写二进制数据 | 高 |
| MemoryStream | 内存中的数据流 | 高 |
| FileStream | 文件流操作 | 高 |
| BufferedStream | 带缓冲的流操作 | 中 |
| TextReader/TextWriter | 文本读写抽象基类 | 中 |

### 2. System.Text 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| RegularExpressions (Regex) | 正则表达式操作 | 高 |
| Encoding (扩展现有实现) | 更多编码支持 | 中 |
| StringBuilder (已实现) | 高效字符串构建 | 已完成 |

### 3. System.Collections 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| Queue | 队列数据结构 | 高 |
| Stack | 栈数据结构 | 高 |
| Hashtable | 哈希表实现 | 中 |
| ArrayList | 动态数组 | 中 |

### 4. System.Collections.Generic 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| HashSet<T> | 哈希集合 | 高 |
| SortedList<T> | 排序列表 | 高 |
| SortedDictionary<TKey, TValue> | 排序字典 | 高 |
| LinkedList<T> | 链表实现 | 中 |
| Queue<T> | 泛型队列 | 高 |
| Stack<T> | 泛型栈 | 高 |

### 5. System.Net 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| WebClient | 简单的 HTTP 客户端 | 高 |
| HttpWebRequest/HttpWebResponse | 高级 HTTP 请求/响应 | 高 |
| Uri | URI 处理 | 高 |
| NetworkCredential | 网络凭证 | 中 |
| WebHeaderCollection | HTTP 头部集合 | 中 |

### 6. System.Security.Cryptography 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| MD5 | MD5 哈希算法 | 高 |
| SHA1/SHA256 | SHA 哈希算法 | 高 |
| AES | 对称加密算法 | 高 |
| RSA | 非对称加密算法 | 中 |
| HMAC | 消息认证码 | 中 |

### 7. System.Threading 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| Thread | 线程操作 | 高 |
| Mutex | 互斥锁 | 高 |
| Semaphore | 信号量 | 高 |
| Monitor | 同步监视器 | 中 |
| AutoResetEvent/ManualResetEvent | 事件同步 | 中 |

### 8. System.Timers 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| Timer | 定时器 | 高 |

### 9. System.Random 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| Random | 随机数生成器 | 高 |

### 10. System.Globalization 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| CultureInfo | 文化信息 | 中 |
| DateTimeFormatInfo | 日期时间格式 | 中 |
| NumberFormatInfo | 数字格式 | 中 |

### 11. System.Reflection 命名空间 (部分已实现)

| API | 描述 | 优先级 |
|-----|------|--------|
| Type (扩展现有实现) | 类型信息 | 高 |
| MethodInfo | 方法信息 | 高 |
| PropertyInfo | 属性信息 | 高 |
| FieldInfo | 字段信息 | 高 |
| Assembly | 程序集信息 | 中 |

### 12. System.Diagnostics 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| Process | 进程操作 | 高 |
| ProcessStartInfo | 进程启动信息 | 高 |
| EventLog | 事件日志 | 中 |
| Stopwatch (已实现) | 性能测量 | 已完成 |

### 13. System.Configuration 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| ConfigurationManager | 配置管理 | 高 |
| AppSettings | 应用设置 | 高 |

### 14. System.Xml 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| XmlReader/XmlWriter | XML 读写 | 高 |
| XmlDocument | XML 文档 | 高 |
| XDocument (LINQ to XML) | LINQ 操作 XML | 中 |

### 15. System.Json 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| JsonValue | JSON 值 | 高 |
| JsonObject | JSON 对象 | 高 |
| JsonArray | JSON 数组 | 高 |

### 16. System.Linq 命名空间 (部分已实现)

| API | 描述 | 优先级 |
|-----|------|--------|
| 更多 LINQ 方法 | 扩展现有实现 | 高 |
| IQueryable | 查询接口 | 中 |

### 17. System.Data 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| DataTable | 数据表 | 中 |
| DataSet | 数据集 | 中 |
| SqlConnection/SqlCommand | SQL 数据库操作 | 中 |

### 18. System.Net.Sockets 命名空间

| API | 描述 | 优先级 |
|-----|------|--------|
| Socket | 套接字操作 | 高 |
| TcpClient/TcpListener | TCP 客户端/服务器 | 高 |
| UdpClient | UDP 客户端 | 高 |

### 19. System.Windows.Forms 命名空间 (可选)

| API | 描述 | 优先级 |
|-----|------|--------|
| Form | 窗体 | 低 |
| Button | 按钮 | 低 |
| TextBox | 文本框 | 低 |
| Label | 标签 | 低 |

### 20. System.Drawing 命名空间 (可选)

| API | 描述 | 优先级 |
|-----|------|--------|
| Color | 颜色 | 低 |
| Font | 字体 | 低 |
| Pen | 画笔 | 低 |
| Brush | 画刷 | 低 |

## 实现建议

1. **保持 API 风格一致**：遵循 .NET API 的命名规范和参数顺序，让 .NET 开发者能够无缝使用。

2. **利用 Go 语言特性**：虽然要保持 .NET API 风格，但也要充分利用 Go 语言的特性，如接口、goroutine 等。

3. **注重性能**：确保实现高效，尤其是在处理大量数据时。

4. **完善测试**：为每个 API 编写全面的测试用例，确保功能正确。

5. **文档完善**：提供详细的文档和使用示例，帮助开发者快速上手。

## 优先级排序

### 高优先级 (建议优先实现)
1. System.IO 相关流操作 (StreamReader/StreamWriter, BinaryReader/BinaryWriter, MemoryStream)
2. System.Text.RegularExpressions (Regex)
3. System.Collections.Generic 集合 (HashSet<T>, SortedList<T>, SortedDictionary<TKey, TValue>)
4. System.Net (WebClient, HttpWebRequest/HttpWebResponse, Uri)
5. System.Security.Cryptography (MD5, SHA1/SHA256, AES)
6. System.Threading (Thread, Mutex, Semaphore)
7. System.Timers.Timer
8. System.Random

### 中优先级
1. System.Collections (Queue, Stack, Hashtable)
2. System.Reflection (扩展现有实现)
3. System.Diagnostics (Process, ProcessStartInfo)
4. System.Configuration (ConfigurationManager, AppSettings)
5. System.Xml (XmlReader/XmlWriter, XmlDocument)
6. System.Json (JsonValue, JsonObject, JsonArray)
7. System.Net.Sockets (Socket, TcpClient/TcpListener, UdpClient)

### 低优先级
1. System.Windows.Forms
2. System.Drawing
3. System.Data

## 结论

NGo 项目已经实现了许多核心 .NET API，为 Go 开发者提供了熟悉的编程体验。通过移植更多的 .NET API，可以进一步丰富项目的功能，使其成为一个更完整的 .NET API 兼容层。

建议按照优先级顺序逐步实现这些 API，同时保持代码质量和性能。这样可以为 Go 开发者提供更多熟悉的工具和功能，同时也为 .NET 开发者迁移到 Go 提供了便利。