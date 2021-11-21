# Log & Go 

This is common logging utility tht is implemented only just for fun :)

I will be giving the details about how to use this utility to log every event occurred in your applications.

## Steps 

### Step 1

`go get  github.com/yukselcodingwithyou/logandgo@latest`

This command above is necessary for you to use this utility in your application.


### Step 2

`import logger "github.com/yukselcodingwithyou/logandgo"`

Import the downloaded utility to your go files, as visualized above.

### Step 3

Create a new logger with a format and level as you wish, format and level information should be like below.

1. **LogLevels**
    - PANIC          = 0
    - ERROR          = 1
    - WARN           = 2
    - INFO           = 3
    - DEBUG          = 4

2. **Formats**
   
   - JSON
   - TEXT

## Examples

 - **JSON Logger Example**

The example below is going to log in _WARN_ level with _JSON_ format. In case you have configured your log level in application level to be WARN, then our logger utility will not log 
INFO or above levels but just _PANIC, ERROR, WARN_ levels

    package main
    import logger "github.com/yukselcodingwithyou/logandgo"
    func main() {
        jsonLogger := logger.NewLogger(logger.JSON, 2)
        jsonLogger.Debug(logger.LogFields{
            "title": "title",
            "reason": "debug",
            "payload": "payload",
        })
        jsonLogger.Error(logger.LogFields{
            "title": "title",
            "reason": "error",
            "payload": "payload",
        })
        jsonLogger.Info(logger.LogFields{
            "title": "title",
            "reason": "info",
            "payload": "payload",
        })
        jsonLogger.Panic(logger.LogFields{
            "title": "title",
            "reason": "panic",
            "payload": "payload",
        })
        jsonLogger.Warn(logger.LogFields{
            "title": "title",
            "reason": "warn",
            "payload": "payload",
        })
    }

- **JSON Logger Output Example**

    `ERROR: 2021/11/21 16:12:00 main.go:14: {"payload":"payload","reason":"error","title":"title"}`

    `PANIC: 2021/11/21 16:12:00 main.go:26: {"payload":"payload","reason":"panic","title":"title"}`
  
    `WARNING: 2021/11/21 16:12:00 main.go:32: {"payload":"payload","reason":"warn","title":"title"}`



- **TEXT Logger Example**

The example below is going to log in _ALL_ levels defined with _TEXT_ format.

      package main
      
      import logger "github.com/yukselcodingwithyou/logandgo"
    
      func main() {

         textLogger := logger.NewLogger(logger.TEXT, 4)
        
         textLogger.Debug(logger.LogFields{
            "title": "title",
            "reason": "debug",
            "payload": "payload",
         })

         textLogger.Error(logger.LogFields{
            "title": "title",
            "reason": "error",
            "payload": "payload",
        })
         
         textLogger.Info(logger.LogFields{
            "title": "title",
            "reason": "info",
            "payload": "payload",
         })
        
         textLogger.Panic(logger.LogFields{
              "title": "title",
              "reason": "panic",
              "payload": "payload",
         })
        
         textLogger.Warn(logger.LogFields{
              "title": "title",
              "reason": "warn",
              "payload": "payload",
         })
      }

- **TEXT Logger Output Example**

`DEBUG: 2021/11/21 16:39:06 main.go:8: title=title reason=debug payload=payload`

`ERROR: 2021/11/21 16:39:06 main.go:14: reason=error payload=payload title=title`

`INFO: 2021/11/21 16:39:06 main.go:20: reason=info payload=payload title=title`

`PANIC: 2021/11/21 16:39:06 main.go:26: title=title reason=panic payload=payload`

`WARNING: 2021/11/21 16:39:06 main.go:32: title=title reason=warn payload=payload `
