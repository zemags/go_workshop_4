- обработчик(контроллер)
- маршрутизатор\servermux(связь между урлами и обработчиками)
- веб сервер()

#### servermux - маршрутизатор запросов. два типа урлов
- **фиксированные** - урл должен точно совпасть с фиксированным путем (без косой черты в конце урла) - /memo /memo/create
- **многоуровневые** (с косой чертой в конце урал) например / (заканчивается косой чертой, типа корневой домен) или /static/

#### http.DetectContentType()
- определения типа контента, если не сможет то поставит Content-Type: application/octet-stream
- **не может отличить json от текста**, поэтому надо самому поставить в заголовках типа : `w.Header().Set("Content-Type", "application/json")`

#### Curl commands
```bash
curl -i -X GET http://127.0.0.1:8080/memo/create
```

#### project layout
- cmd - разные приложения проекта(каждое приложение отдельная папка), здесь только web
- pkg - вспомогательный код, модели бд, утилиты
- ui - шаблоны, статика, изображения и тд

#### Название шаблонов навазние.роль.tmpl
</br>
 - http.FileServer - Обработчик статических файлов (`fileServer := http.FileServer(http.Dir("./ui/static"))`)
</br>

##### интерфейс http.Handler
- обработчик, контроллер - чтобы объект был обработчиком http-запросов у него должен быть метод `ServeHTTP`
```golang
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
#### пример, метод home удовлетворяющий сигнатуре `ServeHTTP`
```golang
func ServeHTTP(w http.ResponseWriter, r *http.Request) {}
```

- если у функции нет метода ServeHTTP, то она в принципе не обработчик
- чтобы функция стала обработчиком надо использовать адаптер **http.HandlerFunc()**, он добавить функции метод `ServeHTTP`
- это соответствование требованиям интерфейса http.Handler
```golang
mux := http.NewServeMux()
mux.Handle("/", http.HandlerFunc(home))
```

#### логгирование в файл при запуске с терминала
```bash
go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
```