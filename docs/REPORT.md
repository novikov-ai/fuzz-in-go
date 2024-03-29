# Фазз-тестирование

С данным видом тестирования раньше был незнаком. Для более глубокого погружения изучил как оно применяется в Go и даже
[написал](../README.md) введение в фаз-тестирование на этом языке. Очень удобно, что он доступен из коробки с определенной версии. 

Решил попробовать Фазз-тестирование на небольшом проекте, который публикует шахматные задачи. 

[Проект](https://github.com/novikov-ai/chess-daily-puzzle) написан на Go и состоит всего из 400-500 строк кода. 

Несмотря на обилие функций, протестировать я смог только одну.

Функция `GetPictureURL` принимает строку в формате pgn и возвращает URL страницы с шахматной задачей или ошибку, если pgn не удалось обработать.

1. Начал с того, что составил массив из ожидаемых входных данных, который послужил сид-корпусом для будущего тестирования.
2. Описал желаемое поведение функции:
   - url должен быть валидным
   - url должен иметь определенный путь к странице, чтобы можно было ее корректно обработать далее
3. В течение нескольких минут прогонял Фазз-тесты (багов не обнаружил)

При этом при запущенном Фаззинг-тестировании (более 2-х минут) у меня падал тест со следующей ошибкой:
~~~
--- FAIL: FuzzGetPictureURL (135.53s)
    fuzzing process hung or terminated unexpectedly while minimizing: EOF
    Failing input written to testdata/fuzz/FuzzGetPictureURL/d48df77a72f3a23d
    To re-run:
    go test -run=FuzzGetPictureURL/d48df77a72f3a23d
~~~
Также по пути `testdata/fuzz/FuzzGetPictureURL/d48df77a72f3a23d` сохранялся соответствующий файл, но при перезапуске тест успешно проходил.
Как будто это проблема в текущей имплементации Фаззинга в Go, так как многие сталкиваются ровно с тем же (ждем фикса).

Конечная функция для тестирования [здесь](../chess-pgn-generator/pgn_test.go).

Идеологически подход очень понравился, но увидел в нем определенные ограничения при постоянной разработке, а именно:
1. Фаззинг тестирование очень требовательно к железу и времени выполнении (рекомендуют его запускать на долгое время на мощном железе)
2. Конкретно в Go столкнулся с ограничениями на работу с типами данных. К сожалению, нет нативной поддержки структур, а только примитивные типы.
3. Для продуктовой разработки, когда важен минимальный TTM (time to market), такой вид тестирования может сильно замедлять выпуск очередной фичи.

Тем не менее, данный вид тестирования выглядит очень полезным в сложных и ответственных проектах, к которым предъявляются повышенные требования по безопасности и отказоустойчивости.
