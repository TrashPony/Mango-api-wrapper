## Обертка для работы с апи https://www.mango-office.ru/

##### описание api: https://www.mango-office.ru/upload/api/MangoOffice_VPBX_API_v1.9.pdf

##### работа с библиотекой: <br>

--- Вызвать метод инициализации и вернется клиент с помощью которого можно отдавать команды, гетать информацию
и получать события

Например:

```
package mango_client

import (
	mango "github.com/TrashPony/Mango-api-wrapper"
)

var client *mango.Client

func init() {
	client = mango.InitMangoCallHandle(
		"Порт внешней системы:",
		"Уникальный код вашей АТС:",
		"Ключ для создания подписи:",
		"Адрес API Виртуальной АТС:")
}

func GetClient() *mango.Client {
	return client
}
```

##### События которые можно получить клиентом "API Realtime"

```
// Уведомление о вызове
func (c *Client) GetCallEvents() 

// Уведомление о нажатиях DTMF клавиш
func (c *Client) GetDTMFEvents() 

// Уведомление о завершении вызова
func (c *Client) GetEndCallsEvents() 

// Уведомление о записи разговора
func (c *Client) GetAddRecordEvents() 

// Уведомление о помещении записи разговора в облачное хранилище
func (c *Client) GetStartRecordEvents() 

// Уведомление о результате отправки SMS
func (c *Client) GetSMSEvents() 
```

##### Команды которые можно получить клиентом "API Команды"
```
// Инициирование вызова от имени сотрудника
func (c *Client) InitCallOfUser(extension, callerNumber, toNumber, lineNumber, sipHeaders string) *mango_objects.Result

// Инициирование вызова от имени группы
func (c *Client) InitCallOfGroup(from, to, lineNumber string) *mango_objects.Result

// Завершение вызова
func (c *Client) EndCall(callID string) *mango_objects.Result

// Отправка SMS
func (c *Client) SendSms(fromExtension, toNumber, smsSender, text string) *mango_objects.Result

// Включение записи разговора
func (c *Client) TurnOnRecordingCall(callID, callPartyNumber string) *mango_objects.Result

// Маршрутизация вызова
func (c *Client) RoutingCall(callID, toNumber, sipHeaders string) *mango_objects.Result

// Перевод вызова
func (c *Client) TransferCall(callID, method, toNumber, initiator string) *mango_objects.Result
```
