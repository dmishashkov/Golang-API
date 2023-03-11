# API проект с CRUD операциями и JWT аутенфикацией

### Вкратце о чём проект  
В этом API пользователи могут регистрироваться и добавлять, удалять, изменять, получать
"статьи", для простоты у статьи есть тело (body) и заголововок (title), аутенфикация осуществляется при помощи
JWT токена, само API написано при помощи популярной библиотеки gin.

### Эндпоинты
Почти все эндпоинты защищены в зависимости от их типа (например, получать статьи может вообще любой человек, даже не зарегестрированный,
но удалять статьи может только аутенфицированный пользователь, причём тот который и является автором)

### Аутенфикация
Пароли шифруются при помощи base64 и дальше хранятся только их хэши и при аутенфикации
они сравниваются, также есть проверка что регистрироваться второй раз нельзя (с одним и тем же логином), 
после аутенфикации пользователю высылается JWT-токен, который он должен передавать на сервер каждый раз при запросе

### База данных

Используется Postgres. Все данные для базы данных и JWT токена хранятся в dotenv файле, который при запуске
все нужные переменные засылает в переменные окружения, откуда они читаются из программы

### Как выглядит ответ от сервера

Ответ от севера приходит в виде JSON, в нём есть три поля: Error, ErrorCode и 
