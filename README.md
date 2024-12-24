<a id="readme-top"></a>


<br />
<h3 align="center">gowebcalc</h3>

  <p align="center">
    Простой калькулятор и API к нему, имплементированные на go. 
  </p>
</div>


<details>
  <summary>Контент</summary>
  <ol>
    <li>
      <a href="#about-the-project">О проекте</a>
      <ul>
        <li><a href="#built-with">Стек</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Запуск проекта</a>
      <ul>
        <li><a href="#prerequisites">Предварительные требования</a></li>
        <li><a href="#installation">Установка и запуск</a></li>
      </ul>
    </li>
    <li><a href="#usage">Использование</a></li>
  </ol>
</details>



## О проекте

<p align="right">(<a href="https://t.me/Basquade">Связаться со мной</a>)</p>


Решение gowebcalc позволяет использовать надежный функционал вычисления арифметических выражений. Проект включает API с HTTP-эндпоинтом /api/v1/calculate, который позволяет отправлять запросы на вычисление выражений.
Gowebcalc работает на стеке net-http, zap для логгирования и viper для оперирования переменными окружения.
Калькулятор следующие арифметические операции:
* Сложение (+)
* Вычитание (-)
* Умножение (*)
* Деление (/)
* Круглые скобки для указания порядка

При возникновении багов или прочих проблем, пожалуйста, свяжитесь со мной:
Тык ->[![Telegram][tg]][tg-url]<- Тык


### Стек

* [![Go][Go]][Go-url]
* [net-http](https://pkg.go.dev/net/http)
* [uber zap](https://github.com/uber-go/zap)
* [viper](https://github.com/spf13/viper)

<p align="right">(<a href="#readme-top">вернуться назад</a>)</p>


## Запуск проекта 

Данный проект можно запустить локально следуя простым шагам:

### Предварительные требования


* Go https://go.dev/doc/install


### Установка и запуск

1. Установите go на ваш компьютер, он доступен почти для всех ОС
2. Склонируйте репозиторий
   ```sh
   git clone https://github.com/AtariOverlord09/gowebcalc.git
   ```
3. Загрузка необходимых пакетов(zap, viper)
   ```sh
   go mod tidy
   ```
4. Создайте и заполните .env файл
   ```.env
   HOST=127.0.0.1
   PORT=8080
   ```
5. Компиляция или запуск
   ```sh
   # go build  -trimpath -ldflags="-s -w" -o gowebcalc.exe cmd/main.go

   # Запуск:
      go run cmd/main.go   
   ```
6. Запуск проекта
   ```sh
   ./gowebcalc
   ```

<p align="right">(<a href="#readme-top">вернуться назад</a>)</p>



<!-- USAGE EXAMPLES -->
## Использование

POST /api/v1/calculate/
Request:
```json
{
    "expression": "2+2"
}
```
Response:
```json
{
    "result": "4.000000"
}
```

Проект может эффективно использоваться в вычислениях арифметических выражение, так как использует обратную польскую нотацию для расчетов. Так же есть api, повзоляющая сделать gowebcalc отличным микросервисом.


<p align="right">(<a href="#readme-top">вернуться назад</a>)</p>



[net-http-url]: https://pkg.go.dev/net/http
[Go-url]: https://go.dev/
[Go]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge
[tg-url]: https://t.me/Basquade 
[tg]: https://img.shields.io/badge/Telegram-2CA5E0?style=flat-square&logo=telegram&logoColor=white
