### Golang за 2 дня? 
<em>Мой первый небольшой проект на Go. На изучение и реализацию было всего 2 дня :)</em><br>

### Usage<br><em>
  1. git clone [тут ссылка на репу]
  2. cd terminalApi
  3. docker-compose build<br>
  4. docker-compose up</em><br><br>

<b><em>Не композится?</b></em><br>
ubuntu - попробуйте с "sudo"<br> / запустите docker-machine
windows - измените на volumes local / если не бридж запускайте от админа<br>
macos - все работало прекрасно<br>
<br>

<b>Роуты</b><br>
- Регистрация выдачи терминала в журнале - "/registerTerm", <em>(принимаемые json данные name,locationId,terminalId)</em><br>
- Регистрация возврата терминала в журнале - "/UnregisterTerm", <em>(принимаемые json данные name,locationId,terminalId)</em><br>
- Получить все данные журнала - "/AllTermData", <em>(ничего не принимает, просто выдает данные)</em><br>
- Получить историю об одном терминале - "/InfoAboutTerm", <em>(принимаемые json данные terminalId)</em><br>
- Узнать кто сейчас использует терминал - "/WhoUseTerminal", <em>(принимаемые json данные terminalId)</em><br>
<br>

## Author
* **Ilya Mekke** - [mtytos](https://github.com/mtytos)
