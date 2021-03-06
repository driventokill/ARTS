# arts-1

## Algorithm

[0001. two sum](./0001-two-sum)

## Tech post review

[Why we rewrote Cloudflow's CLI from Go to Scala](https://www.lightbend.com/blog/writing-kubectl-plugins-with-scala-or-java-with-fabric8-kubernetes-client-on-graalvm)

### 背景

决定重写的原因：

1. 原有的Go程序已经背上了沉重的技术债
2. Go（中的库）对于HOCON（Human-Optimized Config Object Notation）的支持不好，也不是很好修复

### 技术选型考量

| 要求 | 技术选择 |
|----|-----|
| 团队熟悉的编程语言 | Scala/Java |
| Native 性能 | GraalVM AOT |
| 行业标准库 | Fabric8 Kubernetes Client |
| 坚实的生态（技术栈） | HOCON/Scopt/Airframe logging/Pureconfig |

### 挑战

最大的考量感觉就是要使用 JVM 生态（上面的一大把成熟好用的功能库），要去找到满足 Native 性能的方法，主要的助力就是 GraalVM 的 AOT 编译器，面临的挑战：

1. GraalVM 编译的配置维护的成本和风险都很高，特别是使用重度依赖反射的库（比如 Jackson）的时候。还好我们可以完全控制我们程序的边界，只要在一个真实的集群上面训练好配置就可以了。然后通过修改 `sbt generateGraalVMConfig` 的运行，实际上运行另一个程序的Main，在记录 `Assisted configuration` 时尽可能的覆盖大部分可能的代码路径和不变式
2. 另一个挑战是 GraalVM 超长的编译时间，现在 CLI 需要差不多5分钟来编译 native 包。日常开发中直接使用 `sbt run` 去校验改动的内容，不需要编译。
3. 最后一个问题是 GraalVM 不支持交叉编译，我们通过CI为不同的 runner 在目标架构环境下本地编译。

### 成果

1. 让我们回到了舒爽的 JVM 生态圈
2. 拥有了一个内部日志系统可以按需注入，平时静默状态用户无感只的拥有全部的程序运行性能。需要的时候只要命令结尾加一个 `-v trace` 就可以了。这让debug用户问题的时候依然保留着之前Go版本的工效（似乎这是之前Go版本的一个很大优势）
3. Jackson 提供了开箱即用的 `case class` 到 JSON 和 YAML 的序列化能力
4. CLI 可以直接在 ScalaTest 中复用，我们当前已经有了全 Scala 的集成测试，不再需要单独跑一个进程，并且出错的情况下可以直接访问到有用的错误信息
5. 对于整体代码结构很自豪，强壮的关注点分离使得所有东西都容易理解并方便测试

### 总结

1. 从工效、反馈和校验的用户体验
2. 更容易的加入了更多的特性和功能
3. 代码库现在已经更熟悉并更容易理解

## Tip from work

