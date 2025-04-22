## Request UUID Marker
Плагин добавляющий генерацию уникального идентификатора запроса для KrakenD

---

Из коробки KrakenD предоставляет генерацию заголовка с идентификатором запроса путем модификатора, но этого заголовка не будет в ответе, а также нельзя изменить его наименование.

### Установка:
- Склонируйте репозиторий
- Скомпилируйте основной .go файл с **той же версией** go, что и сам KrakenD
- Разместите собранный файл рядом с KrakenD и пропишите плагин в конфиге

**Пример конфига KrakenD:**
```json
"plugin": {
  "pattern": ".so",
  "folder": "/etc/krakend/plugins/" // путь до собранных .so плагинов, которые сканирует и подгружает KrakenD
},
"extra_config": {
  "plugin/http-server": { // тип плагина
    "name": [
      "RequestUUIDMarker" // название плагина
    ]
  }
}
```

В случае использования KrakenD в Docker окружении имеется официальный образ конкретно для сборки плагинов.
<br/>**Пример Dockerfile:**
```Dockerfile
FROM krakend/builder:2.9.4 as plugins

WORKDIR /build

COPY .docker/container/krakend-public/plugins/ .

RUN cd /build/request-uuid-marker/ && go build -buildmode=plugin -o ../output/request-uuid-marker.so ./request-uuid-marker.go

FROM krakend:2.9.4

ENV FC_ENABLE=1

COPY --from=plugins /build/output/ /etc/krakend/plugins/
```

---

**P.S.:** если этот заголовок планируется использовать в браузере, ему необходимо явно разрешить его использовать путем добавления заголовка: `Access-Control-Expose-Headers: X-REQUEST-ID` для Nginx и/или в самом конфиге KrakenD:
```json
"security/cors": {
    ...
    "expose_headers": ["X-Request-Id"]
},
```

**P.S.S:** локально используя Docker и процессоры M* (скорее всего любые ARM) лучше указать конкретную платформу для образов `platform: "linux/x86_64"`, в противном случае придется танцевать с бубном при сборке.
