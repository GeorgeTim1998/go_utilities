# Стратегия
**Стратегия** — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

# Применимость
- Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.

Стратегия позволяет варьировать поведение объекта во время выполнения программы, подставляя в него различные объекты-поведения (например, отличающиеся балансом скорости и потребления ресурсов).

- Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.

Стратегия позволяет вынести отличающееся поведение в отдельную иерархию классов, а затем свести первоначальные классы к одному, сделав поведение этого класса настраиваемым.

- Когда вы не хотите обнажать детали реализации алгоритмов для других классов.

Стратегия позволяет изолировать код, данные и зависимости алгоритмов от других объектов, скрыв эти детали внутри классов-стратегий.

- Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма.

Стратегия помещает каждую лапу такого оператора в отдельный класс-стратегию. Затем контекст получает определённый объект-стратегию от клиента и делегирует ему работу. Если вдруг понадобится сменить алгоритм, в контекст можно подать другую стратегию.

# Плюсы
- Горячая замена алгоритмов на лету.
- Изолирует код и данные алгоритмов от остальных классов.
- Уход от наследования к делегированию.
- Реализует принцип открытости/закрытости.

# Минусы
- Усложняет программу за счёт дополнительных классов.
- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

# Реальный пример
Навигационное приложение для путешественников должно показывать удобную карту, позволяя легко ориентироваться в незнакомом городе. Одной из ключевых функций является поиск и прокладывание маршрутов. Пользователь указывает начальную точку и пункт назначения, а навигатор прокладывает оптимальный путь.

В навигаторе применяются разные алгоритмы: маршруты только для автомобилей, пешие маршруты, маршруты для общественного транспорта. А также маршруты по велодорожкам и туристические маршруты.

Паттерн Стратегия позволяет легко добавлять новые типы маршрутов, не изменяя существующий код. Каждый алгоритм прокладывания маршрута реализуется в отдельном классе (стратегии), и навигатор может переключаться между ними.

Класс навигатора становится контекстом, который использует различные стратегии маршрутов. Вместо выполнения алгоритмов внутри себя, навигатор делегирует эту работу стратегиям, обеспечивая простоту и модульность кода.

Все стратегии маршрутов имеют общий интерфейс. Это позволяет навигатору работать с любыми стратегиями, не зависимо от их реализации, и легко менять алгоритмы в зависимости от условий.