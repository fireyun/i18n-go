# 国际化

## 国际化使用方式

1. 在源码中有调用golang.org/x/text/message包中的`func (p *Printer) Sprintf/Fprintf/Printf`三个方法中的任意一个
2. 首次部署项目或没有gotext命令行工具时执行`make init`，该操作会安装gotext命令行工具<br>
   （该工具能自动从源码中提取要翻译的message以及合并message，并生成对应的目录文件）
3. 执行`make i18n`，从源码中提取要翻译的message以及合并message并更新到pkg/i18n目录下的对应文件中
4. 将pkg/i18n/translations/locales/zh/messages.gotext.json文件中未翻译的内容进行补充，补充完后再次执行`make i18n`，<br>
   这次不会有翻译缺失提示，所有message均有了对应翻译，out.gotext.json和catalog.go文件也都被更新，重新编译部署项目后生效
5. 在请求中指定语言从而获取对应语言的message，如果是http请求，可考虑从header或cookie中获取

## 使用方式详解

### 1.在源码中有调用golang.org/x/text/message包中的`func (p *Printer) Sprintf/Fprintf/Printf`三个方法中的任意一个

- 调用示例

```go
fmt.Println(l.Translate("Welcome!"))
fmt.Println(l.Translate("Who are you? How old are you?"))
fmt.Println(l.Translate("My name is %s, I'm %d years old.", name, age))
```

### 2. 首次部署项目或没有gotext命令行工具时执行`make init`，该操作会安装gotext命令行工具<br>

### （该工具能自动从源码中提取要翻译的message以及合并message，并生成对应的目录文件）

具体操作可见Makefile文件：

```makefile
.PHONY: init
init:
	@echo Download gotext
	go install golang.org/x/text/cmd/gotext@v0.14.0
```

### 3. 执行`make i18n`，从源码中提取要翻译的message以及合并message并更新到pkg/i18n目录下的对应文件中

主要会执行如下操作：

- 从源码中自动提取到要翻译的message到pkg/i18n/translations/locales/{language}/out.gotext.json文件，<br>
  源码中任何调用了golang.org/x/text/message包中的`func (p *Printer) Sprintf/Fprintf/Printf`三个方法中任意一个的message均会被提取
- 合并pkg/i18n/translations/locales/{language}下的out.gotext.json和messages.gotext.json文件，并更新到out.gotext.json<br>
- 自动更新所有语言的翻译内容映射关系到pkg/i18n/translations/locales/catalog.go文件
- 复制pkg/i18n/translations/locales/{language}下的out.gotext.json到messages.gotext.json文件，便于下次修改补充翻译内容
- **注意：如果没有messages.gotext.json文件，out.gotext.json将只是提取后的内容，之前补充的翻译后内容会丢失，<br>
  因为messages.gotext.json文件的存在，重复执行`make i18n`，已翻译过的内容会依然存在**

具体操作可见Makefile文件：

```makefile
.PHONY: i18n
i18n:
	@go generate ./pkg/i18n/translations/translations.go
	@cp ./pkg/i18n/translations/locales/zh/out.gotext.json ./pkg/i18n/translations/locales/zh/messages.gotext.json
```

关于gotext命令行工具的使用见pkg/i18n/translations/translations.go文件中的注释说明：

```go
/*
* 通过gotext命令行工具，能够自动从源码中提取要翻译的message以及合并message，并生成对应的目录文件
* -srclang 指定在源码代中使用的语言，这里是英语en, 语言名称需符合BCP 47规范https://en.wikipedia.org/wiki/IETF_language_tag
* update 该子命令用来从源码中提取要翻译的message以及进行合并操作，并生成对应的目录文件
* -out 指定要生成的message catalog文件，主要是存放不同语言翻译前后的映射关系
* -lang 指定要翻译的目标语言，这里是英语en和中文zh，多个语言之间以逗号分隔
* 最后的参数表示要提取翻译message的包路径，多个包路径以空格分隔（这里使用需要翻译的模块入口所在package，会自动加载这些包的依赖包）
* gotext命令行工具更多使用，见gotext help输出
* gotext源码使用example：https://cs.opensource.google/go/x/text/+/refs/tags/v0.14.0:cmd/gotext/examples/
* 更多实现细节见gotext源码：https://cs.opensource.google/go/x/text/+/refs/tags/v0.14.0:cmd/gotext/main.go
 */
//go:generate gotext -srclang=en update -out=catalog.go -lang=en,zh github.com/fireyun/i18n-go/cmd
```

执行操作后的示例, 如果存在没有翻译的message，输出会提示哪些message缺少对应的翻译：

```shell
$ make i18n
zh: Missing entry for "Welcome!".
zh: Missing entry for "Who are you? How old are you?".
zh: Missing entry for "My name is {Name}, I'm {Age} years old.".
```

查看./pkg/i18n/translations/locales/zh/messages.gotext.json文件，能看到translation字段为空，等待去填充翻译：

```json
{
  "language": "zh",
  "messages": [
    {
      "id": "Welcome!",
      "message": "Welcome!",
      "translation": ""
    },
    {
      "id": "Who are you? How old are you?",
      "message": "Who are you? How old are you?",
      "translation": ""
    },
    {
      "id": "My name is {Name}, I'm {Age} years old.",
      "message": "My name is {Name}, I'm {Age} years old.",
      "translation": "",
      "placeholders": [
        {
          "id": "Name",
          "string": "%[1]s",
          "type": "string",
          "underlyingType": "string",
          "argNum": 1,
          "expr": "name"
        },
        {
          "id": "Age",
          "string": "%[2]d",
          "type": "int",
          "underlyingType": "int",
          "argNum": 2,
          "expr": "age"
        }
      ]
    }
  ]
}
```

### 4. 将pkg/i18n/translations/locales/zh/messages.gotext.json文件中未翻译的内容进行补充，补充完后再次执行`make i18n`，<br>

### 这次不会有翻译缺失提示，所有message均有了对应翻译，out.gotext.json和catalog.go文件也都被更新，重新编译部署项目后生效

**
说明：因为源码语言用的en，翻译后的message和源码相同，会自动更新生成out.gotext.json和messages.gotext.json文件，<br>
所以pkg/i18n/translations/locales/en目录不需要做任何调整**

为translation字段增加翻译内容后的示例：

```json
{
  "language": "zh",
  "messages": [
    {
      "id": "Welcome!",
      "message": "Welcome!",
      "translation": "欢迎！"
    },
    {
      "id": "Who are you? How old are you?",
      "message": "Who are you? How old are you?",
      "translation": "你是谁？你多大了？"
    },
    {
      "id": "My name is {Name}, I'm {Age} years old.",
      "message": "My name is {Name}, I'm {Age} years old.",
      "translation": "我的名字叫{Name}，我{Age}岁了。",
      "placeholders": [
        {
          "id": "Name",
          "string": "%[1]s",
          "type": "string",
          "underlyingType": "string",
          "argNum": 1,
          "expr": "name"
        },
        {
          "id": "Age",
          "string": "%[2]d",
          "type": "int",
          "underlyingType": "int",
          "argNum": 2,
          "expr": "age"
        }
      ]
    }
  ]
}
```

再次执行`make i18n` ，这时所有message均有了对应翻译，out.gotext.json和catalog.go文件也都被更新 <br>
重新编译部署项目后生效

### 5. 在请求中指定语言从而获取对应语言的message，如果是http请求，可考虑从header或cookie中获取
示例见[main方法](./cmd/main.go)
