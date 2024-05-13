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
docker build .
```

**2.** Просматриваем образ Docker следующей командой:

```no-highlight
docker images
```

**3.** Копируем IMAGE ID


| REPOSITORY    | TAG        | IMAGE ID           |  CREATED          |  SIZE  | 
| :-----------: |:----------:| :----------------: | :---------------: | :----: |
| none          | none       | 80e32ad3(копируем) | About an hour ago | 9.17MB |


**4.** Запускаем приложение в контейнере Docker следующей командой, указывая IMAGE ID после флага -it, который ранее копировали:

```no-highlight
docker run -it (указываем ID IMAGE который копировали)
```




