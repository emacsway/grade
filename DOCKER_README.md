# Развертывание проекта с помощью Docker
1. Склонировав проект, нужно создать файл `.env` в корневой директории проекта, скопировать из `.env.example` параметры и установить новые значения (если требуется)
2. Далее, из этой же корневой директории нужно выполнить ```docker compose build```  
![Alt text](readme_images/docker_readme/docker-compose-build-exmpl.png)  
3. Если все успешно, то будет собран образ:
![Alt text](readme_images/docker_readme/docker-image-builded-exmpl.png)  
4. Теперь из этой же директории выполнить ```docker compose -p 'grade' up -d```  
![Alt text](readme_images/docker_readme/docker-compose-up-d-exmpl.png)  
5. Если все успешно, будут созданы контейнеры:  
![Alt text](readme_images/docker_readme/docker-containers-created-console-exmpl.png)  
![Alt text](readme_images/docker_readme/docker-containers-created-in-docker-desktop-exmpl.png)  
- вывода в консоли может быть больше, если нету уже ранее созданных контейнеров `pgAdmin` и `PostgreSQL`  
6. Далее нужно открыть контейнер с `PostgreSQL` - `db-grade` в консоли (ну или зайти в контейнер через `exec`):
![Alt text](readme_images/docker_readme/db-grade-container-exmpl.png)   
7. В данном контейнере выполнить `bash`  
![Alt text](readme_images/docker_readme/db-grade-bash-cmd-exmpl.png)  
- и перейти в директорию в `app` - `cd app` (по умолчанию откроется сразу данная директория, но лучше убедиться) 
8. После в этом же контейнере и директории `app` выполнить команду ```psql -U devel -p 5432 -d devel_grade < ./grade/internal/infrastructure/sql/init.sql```  
![Alt text](readme_images/docker_readme/db-grade-init-sql-exmpl.png)  
- если все успешно, БД будет "поднята"
9. Теперь осталось настроить `pgAdmin` (если нужно)

# Настройка подключения pgAdmin
1. Открыть в браузера контейнер с `pgAdmin` - `pgadmin-grade`  
![Alt text](readme_images/docker_readme/pgadmin-grade-menu-open-in-browser.png)  
2. Откроется страница с предложением ввести пароль:
![Alt text](readme_images/docker_readme/pgadmin-grade-first-open-in-browser.png)  
- пароль вводим из `.env`, указанный в параметре `DB_PASSWORD`  
![Alt text](readme_images/docker_readme/pgadmin-grade-success-enter-pass-exmpl.png)  
3. Теперь нужно настроить сервер, нажимаем на иконку `Add New Server`  
- вводим имя сервера  
![Alt text](readme_images/docker_readme/pgadmin-grade-name-server-exmpl.png)
- вводим настройки подключения  
![Alt text](readme_images/docker_readme/pgadmin-grade-connection-settings-exmpl.png)  
- если все успешно, сервер будет создан и будет доступна БД и ее таблицы  
![Alt text](readme_images/docker_readme/pgadmin-grade-success-connect-exmpl.png)

# Проверка правильности развертывания
1. Если вышеперечисленные шаги выполнены верно, можно проверить пройдут ли тесты
2. Заходим в контейнер `app-grade` и выполняем команду `make test`:
![Alt text](readme_images/docker_readme/app-grade-success-test-exmpl.png)  
3. И проверим таблицу `member`:
![Alt text](readme_images/docker_readme/app-grade-db-table-member-exmpl.png)
4. Приложение развернуто.