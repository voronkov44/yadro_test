# **Решение тестового задания YADRO**
Текст задания находится [тут.](https://docs.google.com/document/d/10BTFT11sPh6iAW3Iu6CfXPUsXCwNeLeX/edit?usp=sharing&ouid=109575002357503548475&rtpof=true&sd=true)
## **Установка**
Для клонирования репозитория необходимо перейти в любую удобную директорию и выполнить команду в терминале:

```no-highlight
git clone https://github.com/voronkov44/yadro_test.git
```

Затем необходимо перейти в корневую директорию проекта:

```no-highlight
cd yadro_test
```

## **Использование**

Для сборки и запуска проекта необходимо выполнить следующие шаги:

На сегодняшний день (05.06.2024) Docker hub вернулся на территорию Российской Федерации, по этому пункт **0** можно пропустить, этот пункт будет работать до 14.06.2024, https://huecker.io/ прекращает свою работу

Пункт **0** оставлю как исторический факт))

**0.** В связи с последними событиями(уход Docker hub из России (30.05.2024)), необходимо установить нужные нам образы:

```no-highlight
docker pull huecker.io/library/golang:1.22-alpine
```

```no-highlight
docker pull huecker.io/library/alpine:3.17
```

Затем необходимо переименовать наши образы в приличное название xD:

```no-highlight
docker tag huecker.io/library/alpine:3.17 alpine:3.17
```

```no-highlight
docker tag huecker.io/library/golang:1.22-alpine golang:1.22-alpine
```

**источник от куда можно скачать образы:** https://huecker.io/ (будет работать до 14.06.2024)

**1.** Собираем образ Docker следующей командой:

Если не установлен [Docker](https://docs.docker.com/) смотрите [зависимости](https://github.com/voronkov44/yadro_test/blob/main/README.md#%D0%B7%D0%B0%D0%B2%D0%B8%D1%81%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B8)

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


## **Запуск приложения**

**3.** Запускаем приложение в контейнере Docker следующей командой:

Если вы запускаете на **Windows**(powerShell):
```no-highlight
docker run -it -v c:/your_Path/test/input_file.txt:/opt/input_file.txt yadro-test:v1 /opt/input_file.txt
```

**Нужно указать свой путь до файла**


Если вы запускаете на **linux**(Bash):
```no-highlight
docker run -it -v `pwd`/test/input_file.txt:/opt/input_file.txt yadro-test:v1 /opt/input_file.txt
```


## **Входной файл**

Если вы хотите поменять входные данные на свои:

На **Windows**(powerShell):

1)Cоздаете файл в директории /test/

2)Наполняете файл входными данными

3)Запускаете приложение в контейнере Docker, указывая путь к вашему файлу, следующей командой:

```no-highlight
docker run -it -v c:/your_Path/test/your_file.txt:/opt/your_file.txt yadro-test:v1 /opt/your_file.txt
```

На **Linux**(Bash):

1)Cоздаете файл в директории /test/

2)Наполняете файл входными данными

3)Запускаете приложение в контейнере Docker, указывая путь к вашему файлу, следующей командой:

```no-highlight
docker run -it -v `pwd`/your_file.txt:/opt/your_file.txt yadro-test:v1 /opt/your_file.txt

```

## **Зависимости**
Установка пакета [Docker Engine](https://docs.docker.com/engine/install/)










