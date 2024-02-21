# Ni-nuclei二开

本项目是nuclei的二开项目。[nuclei](https://github.com/projectdiscovery/nuclei)是projectdiscovery项目开发的一款简单的基于YAML的DSL的快速且可定制的漏洞扫描工具。有各种类型的漏洞模板，漏洞规则极其丰富。nuclei很优秀，所以进行二开！最新的二开版本为nuclei 3.1.10。未来的未来重点使用二开nuclei，在重复造轮子的路上渐行渐远。作者：[0e0w](https://github.com/0e0w)

本项目创建于2022年11月22日，最近的更新时间为2024年2月21日。

- [01-项目结构修改](https://github.com/Goqi/Ernuclei#01-%E9%A1%B9%E7%9B%AE%E7%BB%93%E6%9E%84%E4%BF%AE%E6%94%B9)
- [02-项目功能修改](https://github.com/Goqi/Ernuclei#02-%E9%A1%B9%E7%9B%AE%E5%8A%9F%E8%83%BD%E4%BF%AE%E6%94%B9)
- [03-静态特征修改](https://github.com/Goqi/Ernuclei#03-%E9%9D%99%E6%80%81%E7%89%B9%E5%BE%81%E4%BF%AE%E6%94%B9)
- [04-参考项目资源](https://github.com/Goqi/Ernuclei#04-%E5%8F%82%E8%80%83%E9%A1%B9%E7%9B%AE%E8%B5%84%E6%BA%90)

## 01-项目功能修改

- [ ] 将config和nuclei-templates目录自动下载解压到程序本目录pocs文件夹。
  - [config.go](https://github.com/Goqi/Ernuclei/blob/main/pkg/catalog/config/config.go)
  - [template_path.go](https://github.com/Goqi/Ernuclei/blob/main/pkg/utils/template_path.go)
- [ ] 将nuclei-templates打包到程序中。
- [ ] 优化结果保存。
- [ ] 添加自动代理扫描模块。
- [ ] 扫描资产中添加默认http://。
- [ ] 添加xray和goby的pocs扫描。
- [ ] 优化POC：分类与新增。
- [ ] 修改完善pocs模板，从Cell下载。
- [ ] 删掉自认为很少使用的功能。
- [ ] 随机UserAgent。nuclei从2.8.3已经是随机UA。

## 02-静态特征修改

- [ ] 去除日志打印相关内容#Todo
- [ ] 去除nuclei相关的字段内容#Todo

## 03-项目结构修改

本项目基于nuclei 3.1.10。在原程序结构的基础上进行修改。

```

```

## 04-参考项目资源

- [nuclei代码分析报告](https://github.com/Goqi/ErKai/tree/main/0x01/nuclei)
- https://github.com/xm1k3/cent
