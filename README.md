# 一个简易的问答社区

属于速通了，功能没做很完善，基本的api做了

# 亮点

- 使用go-zero做的微服务框架
- 外接了腾讯云的sms
- 发布问题或者回答的时候使用nsq做了异步处理
- 鉴权加了个中间件用了jwt双token

- 用redis实现点赞
- validator校验数据
- 密码的md5加密和加盐处理

先就做个简易版的，以后有机会优化

接口文档；https://console-docs.apipost.cn/preview/d83654d608c970e2/309712078c877426