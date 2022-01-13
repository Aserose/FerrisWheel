# FerrisWheel (ENG)

**«FerrisWheel»** is a telegram bot. The main function of the bot is collection of images indexed by publication date and geo-data (address name or coordinates) from social networks - if it's publicly available, of course. 
The address can be specified in any common language of the world. And besides that, there is an option to search by time: either until a certain date, or some interval.
The search is currently only includes the vk social network.

The bot also has a blacklist function - if the user does not want to see images of a certain author, then he can add him to the blacklist and they will not be displayed anymore.

Although the application is functional, in its current state it still requires some major architectural fixes. So instead of fixing it, I decided to create a new and more user-friendly application with similar features.

____
### Examples

*will be soon*
____
### Development Information

The connection to the Telegram API is made using Long polling method.

Repository: BoltDB

# FerrisWheel (RUS)

**«FerrisWheel»** — это чатбот для мессенджера Telegram. Основным назначением является поиск изображений по их дате публикации (среди опций: ДО определенной даты, С определенной даты или же ОТ И ДО) и гео-данным (координаты или адрес) в социальных сетях (пока только ВК). Выдача результата поиска осуществляется следующим образом: пользователь получает два сообщения, где первое — это альбом из 10 найденных изображений, а второе — структурированная информация о публикации(дата, место размещения) и цифры, которые нужно отправить боту для добавления учетной записи (с которого было размещено изображение) в "черный список". Затем пользователь может задать новые параметры поиска или продолжить просмотр изображений по уже заданным, нажимая кнопку меню "next". И как было сказано, чатбот имеет ещё одну функцию — "черный список" — это перечень учетных записей, чьи изображения не будут включены в выдачу результата поиска. Пользователь может добавлять учетные записи в "черный список", удалять их оттуда и просматривать, что он в себя включает. 

Планируется сделать аналогичное приложение, которое будет, во-первых, дополнено новыми функциями, во-вторых, сделано более удобным для пользования и, соотвественно, в-третьих, созданным таким образом, чтобы его легко было поддерживать в техническом плане. 
___
### Примеры

*Описание составляется*

___
### Конструкция приложения

* Соединение с API Телеграма сделано по методу Long polling.

* Репозиторий: BoltDB
