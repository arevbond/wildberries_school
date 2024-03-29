
## Прототип (Prototype)

Паттерн Prototype относится к порождающим паттернам уровня объекта.

Паттерн Prototype позволяет создавать новые объекты, путем копирования (клонирования) созданного ранее объекта-оригинала-продукта (прототипа).

Паттерн описывает процесс создания объектов-клонов на основе имеющегося объекта-прототипа, другими словами, паттерн Prototype описывает способ организации процесса клонирования.

Требуется для реализации:

1. Базовый класс Prototype, объявляющий интерфейс клонирования. Все классы его наследующие должны реализовывать этот механизм клонирования;
2. Класс продукта ConcretePrototypeA, который должен реализовывать этот прототип;
3. Класс продукта ConcretePrototypeB, который должен реализовывать этот прототип.

Обычно операция клонирования происходит через метод clone(), который описан в базовом классе и его должен реализовать каждый продукт.

[!] В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс. Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс. Также в языке Go за место общепринятого наследования используется агрегирование и встраивание.