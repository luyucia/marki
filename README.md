# marki

markdown内容服务器

- 一个命令即可把你的文档目录变成在线访问的服务
- 自动遍历目录，构建菜单
- 文件修改，自动更新，无需重启服务

## 如何编译
```
Linux:
执行build.sh

windows:
执行build.cmd
```

## 如何使用
```shell script
marki ./
第一个参数是文档的根目录
目录里需要有.md后缀的文件

详细:
marki -path ./ -host 0.0.0.0 -port 8088

```


## 如何开发前端
```shell script
进到web目录
执行
npm run serve

打包生产文件
npm run build


vue ui
```

## 如何打包静态资源到二进制中
先执行 go generate
assert_generate.go是生成脚本


## todo
1. 美化前端展示，完善全部样式
1. 全文检索
1. 线上编辑
1. 思维导图生成
1. 图表生成
1. 二进制生成
1. 升级检测
1. 快速使用发布三个平台: go平台 yum brew
