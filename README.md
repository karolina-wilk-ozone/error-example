```
➜  example-error go run .                                                               
2026/09/16 13:59:16 listening on :7777
{"time":"2026-09-16T13:59:20.164766+02:00","level":"ERROR","msg":"REQUEST_ERROR","method":"POST","uri":"/tasks","status":400,"latency":95667,"host":"localhost:7777","bytes_in":"12","bytes_out":37,"user_agent":"curl/8.7.1","remote_ip":"::1","request_id":"jhQKxMdaVtsYYlgHIwMBKeGSzWKPenQK","error":"name of task is required"}
{"time":"2026-09-16T14:00:03.723214+02:00","level":"ERROR","msg":"REQUEST_ERROR","method":"POST","uri":"/tasks","status":400,"latency":53542,"host":"localhost:7777","bytes_in":"8","bytes_out":37,"user_agent":"curl/8.7.1","remote_ip":"::1","request_id":"LECPbVEOYqXvSsPARVZTqrGungIiaOcT","error":"name of task is required"}
{"time":"2026-09-16T14:00:22.84134+02:00","level":"INFO","msg":"REQUEST","method":"GET","uri":"/tasks","status":200,"latency":604625,"host":"localhost:7777","bytes_in":"","bytes_out":116,"user_agent":"curl/8.7.1","remote_ip":"::1","request_id":"FkKJvgoLCdlSiIyMgGqgRscVoaJGhgct"}
2026/09/16 14:00:48 request_id= path=/tasks: code=400, message=Bad Request, err=invalid character 'o' in literal null (expecting 'u')
{"time":"2026-09-16T14:00:48.093569+02:00","level":"ERROR","msg":"REQUEST_ERROR","method":"POST","uri":"/tasks","status":400,"latency":103166,"host":"localhost:7777","bytes_in":"15","bytes_out":33,"user_agent":"curl/8.7.1","remote_ip":"::1","request_id":"NnXGVxskSYFgUczuqVQdPccBVTURCdCG","error":"invalid request body"}
2026/09/16 14:00:48 request_id= path=/tasks: code=400, message=Bad Request, err=invalid character 'o' in literal null (expecting 'u')
```