package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	buffer2 "target-management/internal/buffer"
	"target-management/internal/buffer/model"
	"target-management/internal/client/fact_api"
	"target-management/internal/config"
)

var b buffer2.Buffer
var l *log.Logger

func main() {
	httpConfig, err := config.NewHTTPConfig()
	if err != nil {
		log.Fatal("HTTP server config error")
	}
	http.HandleFunc("/load-facts", loadFacts) // HTTP-сервер можно вынести в отдельный пакет internal/server

	l = log.New(os.Stdout, "", 0)

	ctx := context.Background()

	// Все параметры нужно брать из переменных среды
	b = buffer2.NewBuffer(1000)

	// Нужно везде прокинуть контекст и выходить из горутины для gracefull stopping service
	fmt.Println("Start save facts process")
	go saveFacts(ctx)

	// Здесь тоже нужно слушать сигнал от ОС остановки сервиса и корректно останавливать сервис
	fmt.Println("Start HTTP server...")
	err = http.ListenAndServe(httpConfig.Address(), nil)
	if err != nil {
		log.Fatal("HTTP server start error")
	}
}

func loadFacts(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error body close")
		}
	}(r.Body)

	// Добавить валидация для всех принимаемых значений
	periodStart, _ := time.Parse(time.DateOnly, r.FormValue("period_start")) // Обработка ошибки
	periodEnd, _ := time.Parse(time.DateOnly, r.FormValue("period_end"))     // Обработка ошибки
	factTime, _ := time.Parse(time.DateOnly, r.FormValue("fact_time"))       // Обработка ошибки

	b.Push(&model.Fact{
		PeriodStart:         periodStart,
		PeriodEnd:           periodEnd,
		PeriodKey:           r.FormValue("period_key"),
		IndicatorToMoID:     toIntParamRequest(r.FormValue("indicator_to_mo_id")),
		IndicatorToMoFactID: toIntParamRequest(r.FormValue("indicator_to_mo_fact_id")),
		Value:               toIntParamRequest(r.FormValue("value")),
		FactTime:            factTime,
		IsPlan:              toIntParamRequest(r.FormValue("is_plan")),
		AuthUserID:          toIntParamRequest(r.FormValue("auth_user_id")),
		Comment:             r.FormValue("comment"),
	})

	Log("Save fact to buffer")

	// В зависимости от ошибок отдавать подходящий код ответа
	w.WriteHeader(http.StatusOK)
}

func Log(msg string) {
	l.SetPrefix(time.Now().Format(time.RFC3339Nano) + " ")
	l.Println(msg)
}

func toIntParamRequest(v string) int {
	param, _ := strconv.Atoi(v) // Обработка ошибки
	return param
}

func saveFacts(ctx context.Context) {
	// Конфиг клиента API
	factAPIConfig, err := config.NewFactAPI()
	if err != nil {
		log.Fatal("Fact API config error")
	}

	// Клиент API для сохранения фактов
	factAPIClient := fact_api.NewClient(factAPIConfig.GetAPIKey())

	// Таймер, чтобы по времени проверять буффер, есть ли там что-то для сохранения в API
	ticker := time.NewTicker(5 * time.Second) // Будем проверять раз в пять секунд
	defer ticker.Stop()

	// Можно использовать канал для сигнализования что в буфере появились новые факты и их нужно сохранить по API

	for {
		select {
		case <-ticker.C:
			for !b.IsEmpty() { // Пока буфер фактов не пустой
				// Берем из очереди факт
				fact, err := b.Peek()
				if err != nil {
					// TODO: Handle error
				}

				// Отправляем факт на сохранение
				err = factAPIClient.SaveFact(context.Background(), fact)
				if err != nil {
					fmt.Println(err.Error())
					// Тут можно поретраить запрос до тех пор, пока не сохранится факт по ручке API
				} else {
					Log("Send fact to API")
					// Если сохранили без ошибок в API, извлекаем факт из очереди
					b.Pop()
				}

				// Какой-то таймаут, чтобы не грузить ручку API
				// Можно реализовать разными способами
				time.Sleep(1 * time.Second) // Const. Retry. Exponential backoff.
			}
		case <-ctx.Done(): // Нужно для завершения горутины, чтобы не была утечка горутин и памяти
			fmt.Println("Stop save facts process")
			return
		}
	}
}
