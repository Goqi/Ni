# Ernuclei-nuclei二开

本项目是nuclei的二开项目。[nuclei](https://github.com/projectdiscovery/nuclei)是projectdiscovery项目开发的一款简单的基于YAML的DSL的快速且可定制的漏洞扫描器。有各种类型的漏洞模板，漏洞规则极其丰富。但原程序某些功能不太满意，认为存在一些弊端，所以计划进行二开！未来的未来重点使用二开nuclei，尽量避免重复造轮子。作者：[0e0w](https://github.com/0e0w)

本项目创建于2022年11月22日，最近的更新时间为2022年11月26日。

- [01-项目结构修改](https://github.com/Goqi/Ernuclei#01-%E9%A1%B9%E7%9B%AE%E7%BB%93%E6%9E%84%E4%BF%AE%E6%94%B9)
- [02-项目功能修改](https://github.com/Goqi/Ernuclei#02-%E9%A1%B9%E7%9B%AE%E5%8A%9F%E8%83%BD%E4%BF%AE%E6%94%B9)
- [03-静态特征修改](https://github.com/Goqi/Ernuclei#03-%E9%9D%99%E6%80%81%E7%89%B9%E5%BE%81%E4%BF%AE%E6%94%B9)
- [04-参考项目资源](https://github.com/Goqi/Ernuclei#04-%E5%8F%82%E8%80%83%E9%A1%B9%E7%9B%AE%E8%B5%84%E6%BA%90)

## 01-项目功能修改

- [x] 将config和nuclei-templates目录自动下载解压到程序本目录。

  - [config.go](https://github.com/Goqi/Ernuclei/blob/main/pkg/catalog/config/config.go)
  - [template_path.go](https://github.com/Goqi/Ernuclei/blob/main/pkg/utils/template_path.go)

- [ ] 将nuclei-templates打包到程序中。
- [ ] 优化结果保存。
- [ ] 添加自动代理扫描模块。
- [ ] 扫描资产中添加默认http://。
- [ ] 随机UserAgent。
- [ ] 添加xray和goby的pocs扫描。
- [ ] 优化POC：分类与新增。

## 02-静态特征修改

- [ ] 去除日志打印相关内容#Todo
- [ ] 去除nuclei相关的字段内容#Todo

## 03-项目结构修改

本项目基于nuclei-v2.8.3。对项目结构进行了调整，调整后的项目结构如下：

```
│  main.go
│      
├─cmd
│      nuclei.go
│      
└─pkg
    ├─catalog
    │  │  catalogue.go
    │  │  
    │  ├─config
    │  │      config.go
    │  │      config.go.txt
    │  │      
    │  ├─disk
    │  │      catalog.go
    │  │      find.go
    │  │      path.go
    │  │      
    │  └─loader
    │      │  loader.go
    │      │  remote_loader.go
    │      │  
    │      └─filter
    │              path_filter.go
    │              tag_filter.go
    │              
    ├─colorizer
    │      colorizer.go
    │      
    ├─core
    │  │  engine.go
    │  │  execute.go
    │  │  workflow_execute.go
    │  │  workpool.go
    │  │  
    │  └─inputs
    │      │  inputs.go
    │      │  
    │      └─hybrid
    │          │  hmap.go
    │          │  options.go
    │          │  
    │          └─tests
    │                  AS134029.txt
    │                  AS14421.txt
    │                  
    ├─external
    │  └─customtemplates
    │          github.go
    │          s3.go
    │          templates_provider.go
    │          
    ├─input
    │      input.go
    │      
    ├─model
    │  │  model.go
    │  │  worflow_loader.go
    │  │  
    │  └─types
    │      ├─severity
    │      │      severities.go
    │      │      severity.go
    │      │      
    │      ├─stringslice
    │      │      stringslice.go
    │      │      
    │      └─userAgent
    │              user_agent.go
    │              
    ├─operators
    │  │  operators.go
    │  │  
    │  ├─common
    │  │  └─dsl
    │  │          dsl.go
    │  │          
    │  ├─extractors
    │  │      compile.go
    │  │      doc.go
    │  │      extract.go
    │  │      extractors.go
    │  │      extractor_types.go
    │  │      util.go
    │  │      
    │  └─matchers
    │          compile.go
    │          doc.go
    │          match.go
    │          matchers.go
    │          matchers_types.go
    │          validate.go
    │          
    ├─output
    │      doc.go
    │      file_output_writer.go
    │      format_json.go
    │      format_screen.go
    │      output.go
    │      
    ├─parsers
    │      parser.go
    │      workflow_loader.go
    │      
    ├─progress
    │      doc.go
    │      progress.go
    │      
    ├─projectfile
    │      httputil.go
    │      project.go
    │      
    ├─protocols
    │  │  protocols.go
    │  │  
    │  ├─common
    │  │  ├─automaticscan
    │  │  │      automaticscan.go
    │  │  │      doc.go
    │  │  │      
    │  │  ├─compare
    │  │  │      compare.go
    │  │  │      
    │  │  ├─contextargs
    │  │  │      args.go
    │  │  │      contextargs.go
    │  │  │      doc.go
    │  │  │      metainput.go
    │  │  │      variables.go
    │  │  │      
    │  │  ├─executer
    │  │  │      executer.go
    │  │  │      
    │  │  ├─expressions
    │  │  │      expressions.go
    │  │  │      variables.go
    │  │  │      
    │  │  ├─generators
    │  │  │      attack_types.go
    │  │  │      env.go
    │  │  │      generators.go
    │  │  │      load.go
    │  │  │      maps.go
    │  │  │      options.go
    │  │  │      slice.go
    │  │  │      validate.go
    │  │  │      
    │  │  ├─helpers
    │  │  │  ├─deserialization
    │  │  │  │  │  deserialization.go
    │  │  │  │  │  helpers.go
    │  │  │  │  │  java.go
    │  │  │  │  │  
    │  │  │  │  └─testdata
    │  │  │  │          Deserialize.java
    │  │  │  │          README.md
    │  │  │  │          ValueObject.java
    │  │  │  │          
    │  │  │  ├─eventcreator
    │  │  │  │      eventcreator.go
    │  │  │  │      
    │  │  │  ├─responsehighlighter
    │  │  │  │      hexdump.go
    │  │  │  │      response_highlighter.go
    │  │  │  │      
    │  │  │  └─writer
    │  │  │          writer.go
    │  │  │          
    │  │  ├─hosterrorscache
    │  │  │      hosterrorscache.go
    │  │  │      
    │  │  ├─interactsh
    │  │  │      interactsh.go
    │  │  │      
    │  │  ├─marker
    │  │  │      marker.go
    │  │  │      
    │  │  ├─protocolinit
    │  │  │      init.go
    │  │  │      
    │  │  ├─protocolstate
    │  │  │      state.go
    │  │  │      
    │  │  ├─randomip
    │  │  │      randomip.go
    │  │  │      
    │  │  ├─replacer
    │  │  │      replacer.go
    │  │  │      
    │  │  ├─tostring
    │  │  │      tostring.go
    │  │  │      
    │  │  ├─uncover
    │  │  │      uncover.go
    │  │  │      
    │  │  ├─updatecheck
    │  │  │      client.go
    │  │  │      
    │  │  ├─utils
    │  │  │  ├─excludematchers
    │  │  │  │      excludematchers.go
    │  │  │  │      
    │  │  │  └─vardump
    │  │  │          dump.go
    │  │  │          
    │  │  └─variables
    │  │          variables.go
    │  │          
    │  ├─dns
    │  │  │  dns.go
    │  │  │  dns_types.go
    │  │  │  operators.go
    │  │  │  request.go
    │  │  │  
    │  │  └─dnsclientpool
    │  │          clientpool.go
    │  │          
    │  ├─file
    │  │      file.go
    │  │      find.go
    │  │      operators.go
    │  │      request.go
    │  │      
    │  ├─headless
    │  │  │  headless.go
    │  │  │  operators.go
    │  │  │  request.go
    │  │  │  
    │  │  └─engine
    │  │          action.go
    │  │          action_types.go
    │  │          engine.go
    │  │          http_client.go
    │  │          instance.go
    │  │          page.go
    │  │          page_actions.go
    │  │          rules.go
    │  │          util.go
    │  │          
    │  ├─http
    │  │  │  build_request.go
    │  │  │  cluster.go
    │  │  │  http.go
    │  │  │  http_method_types.go
    │  │  │  operators.go
    │  │  │  request.go
    │  │  │  request_annotations.go
    │  │  │  request_condition.go
    │  │  │  request_generator.go
    │  │  │  signature.go
    │  │  │  utils.go
    │  │  │  validate.go
    │  │  │  
    │  │  ├─fuzz
    │  │  │      doc.go
    │  │  │      execute.go
    │  │  │      fuzz.go
    │  │  │      parts.go
    │  │  │      
    │  │  ├─httpclientpool
    │  │  │      clientpool.go
    │  │  │      
    │  │  ├─race
    │  │  │      syncedreadcloser.go
    │  │  │      
    │  │  ├─raw
    │  │  │      doc.go
    │  │  │      raw.go
    │  │  │      
    │  │  ├─signer
    │  │  │      aws-sign.go
    │  │  │      signer.go
    │  │  │      
    │  │  ├─signerpool
    │  │  │      signerpool.go
    │  │  │      
    │  │  └─utils
    │  │          variables.go
    │  │          
    │  ├─network
    │  │  │  network.go
    │  │  │  network_input_types.go
    │  │  │  operators.go
    │  │  │  request.go
    │  │  │  
    │  │  └─networkclientpool
    │  │          clientpool.go
    │  │          
    │  ├─offlinehttp
    │  │      find.go
    │  │      offlinehttp.go
    │  │      operators.go
    │  │      read_response.go
    │  │      request.go
    │  │      
    │  ├─ssl
    │  │      ssl.go
    │  │      
    │  ├─utils
    │  │      utils.go
    │  │      
    │  ├─websocket
    │  │      websocket.go
    │  │      
    │  └─whois
    │      │  whois.go
    │      │  
    │      └─rdapclientpool
    │              clientpool.go
    │              
    ├─reporting
    │  │  reporting.go
    │  │  
    │  ├─dedupe
    │  │      dedupe.go
    │  │      
    │  ├─exporters
    │  │  ├─es
    │  │  │      elasticsearch.go
    │  │  │      
    │  │  ├─markdown
    │  │  │      markdown.go
    │  │  │      
    │  │  └─sarif
    │  │          sarif.go
    │  │          
    │  ├─format
    │  │      format.go
    │  │      
    │  └─trackers
    │      ├─github
    │      │      github.go
    │      │      
    │      ├─gitlab
    │      │      gitlab.go
    │      │      
    │      └─jira
    │              jira.go
    │              
    ├─runner
    │  │  banner.go
    │  │  defaults.go
    │  │  doc.go
    │  │  enumerate.go
    │  │  healthcheck.go
    │  │  inputs.go
    │  │  options.go
    │  │  proxy.go
    │  │  runner.go
    │  │  templates.go
    │  │  update.go
    │  │  
    │  └─nucleicloud
    │          cloud.go
    │          types.go
    │          utils.go
    │          
    ├─templates
    │  │  cluster.go
    │  │  compile.go
    │  │  doc.go
    │  │  log.go
    │  │  preprocessors.go
    │  │  templates.go
    │  │  templates_doc.go
    │  │  templates_doc_examples.go
    │  │  workflows.go
    │  │  
    │  ├─cache
    │  │      cache.go
    │  │      
    │  └─types
    │          types.go
    │          
    ├─testutils
    │  │  integration.go
    │  │  testutils.go
    │  │  
    │  └─testheadless
    │          headless_local.go
    │          headless_runtime.go
    │          
    ├─types
    │      interfaces.go
    │      proxy.go
    │      resume.go
    │      types.go
    │      
    ├─utils
    │  │  insertion_ordered_map.go
    │  │  template_path.go
    │  │  template_path.go.txt
    │  │  utils.go
    │  │  
    │  ├─monitor
    │  │      monitor.go
    │  │      
    │  ├─stats
    │  │      doc.go
    │  │      stats.go
    │  │      
    │  └─yaml
    │          yaml_decode_wrapper.go
    │          
    └─workflows
            doc.go
            workflows.go
```

## 04-参考项目资源

- [nuclei代码分析报告](https://github.com/Goqi/ErKai/tree/main/0x01/nuclei)
- https://github.com/xm1k3/cent
