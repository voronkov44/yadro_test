# **Решение тестового задания YADRO**
Текст задания находится [тут.](https://docs.google.com/document/d/10BTFT11sPh6iAW3Iu6CfXPUsXCwNeLeX/edit?usp=sharing&ouid=109575002357503548475&rtpof=true&sd=true)
## **Установка**
Для скачивания репозитория необходимо перейти в любую удобную директорию и выполнить команду в терминале:

```no-highlight
git clone https://github.com/voronkov44/yadro_test.git
```

Затем необходимо перейти в корневую директорию проекта:

```no-highlight
cd yadro_test
```

## **Использование**

Для сборки и запуска проекта необходимо выполнить следующие шаги:

**1.** Собираем образ Docker следующей командой:

```no-highlight
docker build . -t yadro-test:v1
```

**2.** Просматриваем образ Docker следующей командой:

```no-highlight
docker images
```

Должна появиться такая табличка(IMAGE ID будет различаться)


| REPOSITORY    | TAG        | IMAGE ID           |  CREATED          |  SIZE  | 
| :-----------: |:----------:| :----------------: | :---------------: | :----: |
| yadro-test    | v1         | 80e32ad3s342       | About an hour ago | 9.17MB |


**3.** Запускаем приложение в контейнере Docker следующей командой:

```no-highlight
docker run -it yadro-test
```




