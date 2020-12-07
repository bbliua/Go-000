##### 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

##### 思考
如果不Wrap这个error，可能的一种做法是: 直接返回 err = nil 以及 nil 结构体。
上层为了知道是否有相关记录的行，需要判断 result == nil 且 err == nil，这可能是一种做法，但感觉繁琐，不够直接。
经分析，倾向于 Wrap 这个 error，但如果直接 Wrap 这个 error，由于dao的多数据源service层又不好判定“sql.ErrNoRows”？ 
... 我们可以抽象定义统一的错误码，Wrap这个错误码，抛出去。
