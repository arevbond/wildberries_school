
## Заместитель (Proxy)

Паттерн Proxy относится к структурным паттернам уровня объекта.

Паттерн Proxy предоставляет объект для контроля доступа к другому объекту.

Другое название паттерна - "Суррогат". В этом понимании, это предмет или продукт, заменяющий собой какой-либо другой предмет или продукт, с которым суррогат имеет лишь некоторые общие свойства, но он не обладает всеми качествами оригинального предмета или продукта.

Паттерна Proxy выдвигается ряд важных требований, а именно то, что оригинальный объект и его суррогат должны взаимодействовать друг с другом, а также должна быть возможность, замещения оригинальным объектом, суррогата в месте его использования, соответственно интерфейсы взаимодействия оригинального объекта и его суррогата должны совпадать.

Вам будет легче понять паттерн, если вы смотрели фильм "Суррогаты".

Требуется для реализации:

1. Интерфейс Subject, являющейся общим интерфейсом для реального объекта и его суррогата;
2. Класс RealSubject, реализующий реальный объект;
3. Класс Proxy, реализующий объект суррогата. Хранит в себе ссылку на реальный объект, что позволяет заместителю обращаться к реальному объект напрямую;

Например, паттерн Proxy можно использовать, если нам нужно управлять ресурсоемкими объектами, но мы не хотим создавать экземпляры таких объектов до момента их реального использования.

Вы можете подумать, что это тоже самое, что и Adapter или Decorator. Но...

Proxy предоставляет своему объекту тот же интерфейс.
Adapter предоставляет другой интерфейс.
Decorator предоставляет расширенный интерфейс.

[!] В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс. Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс. Также в языке Go за место общепринятого наследования используется агрегирование и встраивание.
