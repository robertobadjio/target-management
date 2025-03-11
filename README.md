### Запуск сервиса
```go run cmd/buffer/buffer.go ```

### Загрузить факты
```install-deps-linux```

```make load-testing```

```
Start save facts process
Start HTTP server...
2025-03-11T19:36:25.860383324+03:00 Save fact to buffer
2025-03-11T19:36:25.860822299+03:00 Save fact to buffer
2025-03-11T19:36:25.8609327+03:00 Save fact to buffer
2025-03-11T19:36:25.86058666+03:00 Save fact to buffer
2025-03-11T19:36:25.861179712+03:00 Save fact to buffer
2025-03-11T19:36:25.861189649+03:00 Save fact to buffer
2025-03-11T19:36:25.861345753+03:00 Save fact to buffer
2025-03-11T19:36:25.861456441+03:00 Save fact to buffer
2025-03-11T19:36:25.861509755+03:00 Save fact to buffer
2025-03-11T19:36:25.861584821+03:00 Save fact to buffer
2025-03-11T19:36:29.666504388+03:00 Send fact to API
2025-03-11T19:36:30.816971448+03:00 Send fact to API
2025-03-11T19:36:31.943592814+03:00 Send fact to API
2025-03-11T19:36:33.21320479+03:00 Send fact to API
2025-03-11T19:36:34.622898418+03:00 Send fact to API
2025-03-11T19:36:35.739502337+03:00 Send fact to API
2025-03-11T19:36:36.978788216+03:00 Send fact to API
2025-03-11T19:36:38.092533121+03:00 Send fact to API
2025-03-11T19:36:39.231546014+03:00 Send fact to API
2025-03-11T19:36:40.391259733+03:00 Send fact to API
```

```
./bin/bombardier \
        -c 10 \
        -m POST http://localhost:8081/load-facts \
        -b "period_start=2024-12-01&period_end=2024-12-31&period_key=month&indicator_to_mo_id=227373&indicator_to_mo_fact_id=0&value=3&fact_time=2024-12-31&is_plan=0&auth_user_id=40&comment=buffer Fatkhullin" \
        -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" \
        -n 10 \
        -l
Bombarding http://localhost:8081/load-facts with 10 request(s) using 10 connection(s)
 10 / 10 [================================================================================================================================================================================================================================================================================] 100.00% 49/s 0s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec      2105.51       0.00    2105.51
  Latency        3.27ms   537.43us     4.46ms
  Latency Distribution
     50%     3.03ms
     75%     3.43ms
     90%     3.84ms
     95%     4.46ms
     99%     4.46ms
  HTTP codes:
    1xx - 0, 2xx - 10, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     0.86MB/s

```