##Обертка для работы с апи https://www.mango-office.ru/

##### описание api: https://www.mango-office.ru/upload/api/MangoOffice_VPBX_API_v1.9.pdf

работа с библиотекой: <br>
1 вызвать метод инициализации 
	InitMangoCallHandle("Порт внешней системы:", "Уникальный код вашей АТС:", "Ключ для создания подписи:", "Адрес API Виртуальной АТС:")
	
2 Получать события из mango_events.Events

Например:
```
func readerEvents() {
	for {
		call := <-mango_events.Events.Calls
		println(call.CallID)
	}
}
```
