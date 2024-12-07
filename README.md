# starter-new

Как использовать:
1. вставить cookie yandex cloud в файл cookie.txt в корне проекта
2. go run main.go
3. ввести имя нового проекта. Например, example
4. если сертификат не будет выпущен в течение 20 минут, в консоли будет сообщение об этом. Сертификат для Gateway нужно будет донастроить руками   

Что получим:
1. сайт example.upikoth.dev
2. репозитории фронт и бэк на github
3. будет настроена вся инфраструктура проекта и ci/cd. Остается только коммитить в репозитории

Что нужно сделать ручками:
1. Настроить oauth (создать новое приложение во всех сервисах, заменить соответсвтующие env переменные)

Oauth:
1. vk - https://id.vk.com/about/business/go/accounts/159787/apps
2. mailru - https://o2.mail.ru/app/
3. yandex - https://oauth.yandex.ru/

После создания нового приложения в oauth сервисах нужно заполнить переменные локально и в ci/cd

- OAUTH_VK_CLIENT_ID=
- OAUTH_VK_CLIENT_SECRET=
- OAUTH_VK_REDIRECT_URL=

- OAUTH_MAIL_CLIENT_ID=
- OAUTH_MAIL_CLIENT_SECRET=
- OAUTH_MAIL_REDIRECT_URL=

- OAUTH_YANDEX_CLIENT_ID=
- OAUTH_YANDEX_CLIENT_SECRET=
- OAUTH_YANDEX_REDIRECT_URL=
