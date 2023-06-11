# Golang DDD Reference Application "Grade"

Система квалификационной классификации участников экспертных сообществ на основе хорошо зарекомендовавшего себя на практике алгоритма спортивной разрядности, где уровень мастерства спортсмена определяется уровнем мастерства других спортсменов.


## Бизнес-требования

Бизнес-требования к системе описаны в разделе "[4.6. Система квалификационной классификации членов Организации](https://github.com/emacsway/charter/blob/main/charter.md#46-%D1%81%D0%B8%D1%81%D1%82%D0%B5%D0%BC%D0%B0-%D0%BA%D0%B2%D0%B0%D0%BB%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D0%BE%D0%B9-%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D0%B8-%D1%87%D0%BB%D0%B5%D0%BD%D0%BE%D0%B2-%D0%BE%D1%80%D0%B3%D0%B0%D0%BD%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D0%B8)" Устава региональной общественной организации "Объединение ИТ-Архитекторов".


## Причины возникновения проекта

Подобные системы существуют во многих общественных объединениях.

Например, в Проектной Ассоциации:

- https://projects.management/infopage.html?Page=vision
- https://projects.management/infopage.html?Page=statuses

В LinkedIn есть система Endorsement.

Большую популярность набирают социальные токены.

В данном проекте предпринята попытка сделать систему максимально объективной, независимой, равноправной, прозрачной, распределенной и простой.
И одновременно с этим - максимально защищенной от фальсификаций, субъективизма и когнитивных искажений.

Ниже перечислены некоторые из причин её появления:

1. Часто приходится слышать о том, что засилье коммерциализированных сертификатов не отражает реальный уровень экспертности.  
Мы считаем, что экспертному сообществу виднее, и решили предоставить именно ему право определять уровень экспертности своих участников.  
Никто не может вмешиваться в этот процесс.  
И руководитель организации, и новичок имеют равные права рекомендовать и быть рекомендованным.

2. Еще Gregor Hohpe подсветил ключевую проблему экспертных сообществ - Эффект Даннинга-Крюгера, по причине которого генерируется большое количество информационных помех в любом экспертном сообществе, что повышает когнитивную нагрузку на участников сообщества и демотивирует грамотных экспертов.\
Система призвана восстановить качество информационного пространства.

3. Другая проблема заключается в том, что тот, кто больше всех занят делом, как правило, наиболее скромен в общении в силу дефицита времени.\
Зачастую это приводит к гегемонии бескомпетентности в информационном пространстве экспертных сообществ - от этого страдает большинство технических пабликов.

4. Закон об общественных объединениях не позволяет принимать решения по самоуправлению дифференцировано, но предусматривает возможность создания добровольных совещательных органов.\
Разные люди обладают разным уровнем экспертности, игнорирование которого не позволило бы максимально полно отразить экспертность в рекомендациях совещательного органа.

5. Большинство систем блокирования спамеров в telegram-сообществах очень примитивны и рискованы.\
Их можно усовершенствовать, если учитывать ценность вклада и уровень экспертности того, кто банит, и того, кого банят.\
Таким образом можно существенно облегчить бан спамеров и защититься от целенаправленных атак против весомых участников telegram-сообщества.

6. К экспертным сообществам нередко обращаются с запросами на консалтинг.\
Если другие компании доверяют сообществу, то сообщество должно стремиться, к тому, чтобы оправдать доверие, и предпринять конкретные шаги к тому, чтобы эти запросы были адресованы в первую очередь к тем, кто обладает наивысшим уровнем экспертности, выраженной конкретным опытом, ценность которого подтверждена другими членами организации.

7. Как reference application, мне этот проект нужен для систематизации своих собственных знаний по организации структуры кода и для обучения программистов.


## Область применения проекта

Проект представляет интерес не только как демонстрационное, но еще и как вполне реально функционирующее приложение, которое, используя принципы диалектики и [теории игр](https://t.me/emacsway_log/1067), позволяет существенно повысить объективность т.н. [карма-движков](https://github.com/pinax/pinax-points), и, как результат, повысить качество информационного пространства экспертных сообществ.

Подробнее о Теории Контрактов можно посмотреть в видео "[Задача о коллективной ответственности](https://youtu.be/pjEhGZpQLn0)" / Алексей Савватеев или в книге "Теория игр: Искусство стратегического мышления в бизнесе и жизни" / Авинаш Диксит и Барри Нейлбафф ("The Art of Strategy: A Game Theorist's Guide to Success in Business and Life" by Avinash K. Dixit, Barry J. Nalebuff).

Проект также призван выступить альтернативой малополезным формально-бюрократическим grade-системам ИТ-компаний на основе матриц компетентностей, которые не способствуют развитию специалистов, а, наоборот, препятствуют, оттягивая ресурсы времени на нерелевантные аспекты в ущерб релевантным.

Самое главное - экспертные сообщества смогут себя квалифицировать и развивать без всяких тестов, эталонов, экзаменаторов и пр. ограничителей развития, установивших монополию на компетентность.

Изначально проект создавался для нужд региональной общественной организации "Объединение ИТ-Архитекторов".
Теперь же спектр его применения просматривается гораздо шире:

- экспертные общественные объединения (голосования совещательных органов с учетом весового коэффициента квалификационного класса участника);
- grade-системы коммерческих ИТ-компаний;
- самоподдержание уровня квалификации исполнителей в ИТ-франчайзинге;
- ранжирование исполнителей на outsourcing marketplaces;
- подключение в качестве более продвинутой версии карма-движка посредством telegram-бота в экспертные telegram-сообщества;
- система бана telegram-сообщества по типу [@banofbot](https://github.com/backmeupplz/banofbot) с учетом коэффициента квалификационного класса голосующего за бан;
- и пр.


## Технические особенности

С технической стороны проект принципиально не использует ORM (чтоб продемонстрировать, как можно без него обходиться);
принципиально соблюдает ключевые принципы OOP, особенно инкапсуляцию, поскольку иначе технически невозможно гарантировать инварианты агрегатов;
использует CQRS/EventSourcing, причем, будет реализовывать Causal Consistency посредством Causal Dependencies.

Задача амбициозная. Взялся за нее потому, что не смог отыскать существующего прецедента.
Реализация сопровождается интенсивной исследовательской работой, и каждая строка кода подкреплена теоретическими изысканиями в десятках архитектурных книг.
Будет документация, ADR, архитектурная документация и трассировка требований - в общем, будет демонстрация всех SDLC-этапов разработки.


## Документация

Сопроводительная методическая информация проекта будет накапливаться [здесь](https://dckms.github.io/system-architecture/emacsway/it/ddd/grade/index.html).


## Contributing

Присоединиться к разработке: https://t.me/emacsway

Единственное требование - наличие теоретической базы DDD или способность прорабатывать теорию на ходу и принимать обоснованные и информированные решения.
