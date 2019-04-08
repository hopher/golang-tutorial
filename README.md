# golang-tutorial

Go 语言开发入门教程

## 运行服务
```
hugo server -D --bind=0.0.0.0
```

## 添加文章
```
hugo new posts/201811/my-first-post/index.md
```

## 文件结构
```
|-- deploy.sh       # 发布静态网站
|-- auto-commint.sh # 提交源文件 (排除docs目录)
|-- run.sh          # 本地运行
|-- config
  |-- _default      # 默认配置
  |-- _production   # 生产配置
  |-- staging       # 开发配置
```

## LICENSE
MIT

## 参考资料
- [Github部署说明](https://gohugo.io/hosting-and-deployment/hosting-on-github/)
- [搜索说明](https://gohugo.io/tools/search/)
- [多语言支持说明](https://gohugo.io/content-management/multilingual/)
- [hugo-book 主题](https://github.com/alex-shpak/hugo-book)
- [gitalk 评论插件](https://github.com/gitalk/gitalk)
- [github-markdown-css](https://github.com/sindresorhus/github-markdown-css)