
Современные процессоры (CPU) имеют несколько ядер, и некоторые поддерживают гиперпоточность, поэтому они могут обрабатывать несколько инструкций одновременно. Чтобы полностью утилизировать современные CPU, нужно использовать конкурентное программирование.

*Конкурентные вычисления* — это форма вычислений, когда несколько инструкций выполняются, пересекаясь, в течение одного временного периода.

Например есть 2 инструкции, которые нужно выполнить: *А* и *B*. При выполнении инструкций конкурентно ядро процессора вычисляет инструкцию *A*, при ожидающей *B*. Когда первая инструкция *A* заканчивает активное вычисление, она ставится на паузу, и ядро переключается на вычисление инструкции *B*.

Рассмотрим простой пример конкурентности в реальной жизни: HTTP запрос к стороннему сервису по сети:
- вычисление *А* делает HTTP запрос
- пока *А* ждет ответа, начинает выполняться вычисление *B*
- когда ответ пришел, *А* возвращается в работу

Если в процессоре присутствует более одного ядра, обработка происходит *параллельно*. Возвращаясь к примеру выше, инструкции *A* и *B* будут выполняться независимо в один момент времени.

Конкурентные вычисления могут происходить в программе, компьютере или сети. В данном курсе рассматривается только уровень программ.

Как это все относится к Go? Разберемся в следующем уроке.
