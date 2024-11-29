Requirements:

1.Application should handle atleast 10k requests /sec
2.One Get Restpoint /api/verve/accept where id and endpoint is given as a parameter with it  
3.ensoure deduplication also works when put behind an load balancer
4.send the count of unique received ids to a distributed streaming service of your choice.

Implementations:
1. a go application can easily handle 10k requests/sec and when behind a lb we can run at any scale if the backends are idempotent
2. to ensure deduplication also works when put behind an load balancer redis is used , we can use any database but as redis is an inmemory database and offers milli second latency so that we can easily handle the requests
3. kafka is used to recieve the unique ids for event processing 