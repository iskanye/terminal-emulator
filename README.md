# Эмулятор терминала на Go
Простой эмулятор терминала Unix с минимальным набором команд. На данный момент команды находят в процессе реализации
## Функционал
* Команды `ls`, `cd`, `exit`
* На этом пока всё...
## To do
* Сделать эмулятор настраиваемым
* Реализовать VFS на основе xml-файла
* Реализовать команды `du`, `tail`, `cat`, `touch`, `rmdir` и `help`
# Сборка и запуск
Для сборки запустить следующие команды
```bash
git clone https://github.com/iskanye/terminal-emulator
cd terminal-emulator
go build
```
После запустить получившийся exe файл
```bash
.\terminal-emulator.exe
```

## Лицензия
Copyright © 2025 [исканье](https://github.com/iskanye)\
Этот проект использует [MIT](LICENSE) лицензию