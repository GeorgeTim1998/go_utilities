# Состояние
**Состояние** — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния. Извне создаётся впечатление, что изменился класс объекта.

# Применимость
- Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.

Паттерн предлагает выделить в собственные классы все поля и методы, связанные с определёнными состояниями. Первоначальный объект будет постоянно ссылаться на один из объектов-состояний, делегируя ему часть своей работы. Для изменения состояния в контекст достаточно будет подставить другой объект-состояние.

- Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.

Паттерн предлагает переместить каждую ветку такого условного оператора в собственный класс. Тут же можно поселить и все поля, связанные с данным состоянием.

- Когда вы сознательно используете табличную машину состояний, построенную на условных операторах, но вынуждены мириться с дублированием кода для похожих состояний и переходов.

Паттерн Состояние позволяет реализовать иерархическую машину состояний, базирующуюся на наследовании. Вы можете отнаследовать похожие состояния от одного родительского класса и вынести туда весь дублирующий код.

# Плюсы
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.
- Упрощает код контекста.

# Минусы
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

# Реальный пример
Реальный пример использования паттерна "Состояние" можно найти в системах управления принтерами.

Принтер может находиться в различных состояниях, таких как "готов к печати", "печатает", "нуждается в бумаге", "нуждается в обслуживании", "ошибка". Каждое из этих состояний требует различных действий в ответ на одни и те же события.

Без паттерна "Состояние" каждое действие системы управления принтером (например, попытка печати документа) должно было бы включать большое количество условных операторов (if-else или switch) для определения текущего состояния принтера и соответствующего действия. Это бы усложнило и запутало код.

Паттерн "Состояние" позволяет инкапсулировать поведение, связанное с конкретным состоянием, в отдельные классы. Это упрощает добавление новых состояний и изменение существующих, поскольку изменения затрагивают только соответствующие классы состояний, а не весь код системы.

В случае с принтером класс "Принтер" (контекст) не должен знать детали каждого состояния. Он просто делегирует действия текущему объекту состояния. Это делает класс "Принтер" более простым и понятным.