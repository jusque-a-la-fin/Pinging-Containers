## Как запустить:
```bash
git clone git@github.com:jusque-a-la-fin/Pinging-Containers.git && cd Pinging-Containers && sudo docker compose up
```
Подождать, когда контейнер "frontend" отобразится в терминале (Frontend-сервис готов к работе). В браузере перейти по адресу: http://localhost
### Дополнительно:
- Pinger и Backend обмениваются сообщениями через сервис очередей RabbitMQ
- Nginx обслуживает статические файлы Frontend-сервиса и используется как Reverse proxy для доступа к Backend-сервису

