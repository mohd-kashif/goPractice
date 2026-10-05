# Adyen Senior Backend Engineer — System Design
## Payment Service Provider (PSP)

This document contains the complete interview-practice solution for designing a Payment Service Provider for an Adyen-style Senior Backend Engineer system-design interview.

---

## 1. Problem Statement

Design a Payment Service Provider (PSP) that allows merchants to accept payments from their customers.

Initial payment methods:
- Card
- UPI

The system should be extensible to support additional payment methods.

The PSP integrates with multiple external payment processors/banks/payment networks and routes payments to an appropriate provider.

---

## 2. Requirements

### Functional
- Merchants can create/initiate payments.
- Support card and UPI initially.
- Merchants can query payment status.
- Merchants receive asynchronous payment status updates through webhooks.
- Multiple downstream payment processors/providers are supported.
- Payment requests can be routed to an appropriate provider.
- Merchant retries must not create duplicate payments.
- Payments eventually reach a terminal state such as SUCCESS or FAILED.

### Non-functional
- Peak payment creation traffic: ~10,000 requests/sec.
- Payment API acknowledgement p99: <2 seconds.
- Availability target: 99.99%.
- Horizontally scalable.
- Payment correctness is more important than availability during ambiguous payment states.
- No duplicate charges.
- At-least-once internal messaging is acceptable.
- Payment state requires strong correctness.
- Eventual consistency is acceptable for cache/notifications/analytics.

---

## 3. Questions to Clarify

Before designing, ask:
- Which payment methods?
- Direct bank integration or multiple processors?
- Synchronous or asynchronous processing?
- Can merchants retry?
- Refund/capture requirements?
- Webhook requirements?
- Data retention?
- Historical queries?
- Peak TPS?
- Read/write ratio?
- p99 latency?
- Availability?
- Regional requirements?
- RPO/RTO?
- Security/PCI requirements?

---

# 4. Core Invariants

### 1. Merchant retries must not create another payment

The same merchant idempotency key must refer to the same logical payment.

### 2. Provider timeout does not mean failure

An external timeout can mean the provider processed the payment but the response was lost.

Never blindly retry an ambiguous external operation.

### 3. Payment state transitions must be atomic

Two workers must not both acquire the same payment for processing.

### 4. External side effects must be idempotent

Use provider-side idempotency keys wherever supported.

### 5. Webhooks are at-least-once

Exactly-once delivery over an unreliable network cannot be guaranteed. Merchants must deduplicate events.

---

# 5. Payment State Machine

```text
                  +-----------+
                  |  CREATED  |
                  +-----+-----+
                        |
                        v
                  +-----------+
                  | PROCESSING|
                  +---+---+---+
                      |   |
             success  |   | failure
                      |   |
                      v   v
                +-------+ +------+
                |SUCCESS| |FAILED|
                +---+---+ +------+
                    |
                    v
                +---------+
                | CAPTURED|
                +----+----+
                     |
                     v
                +---------+
                | REFUNDED|
                +---------+
```

Important:

`PROCESSING` means the payment has entered the processing workflow, but the final external outcome may still be unknown.

It does NOT mean the external provider definitely processed it.

---

# 6. High-Level Architecture

```text
                         +----------------+
                         |    Merchant    |
                         +-------+--------+
                                 |
                               HTTPS
                                 |
                                 v
                    +--------------------------+
                    | API Gateway / Load Balancer|
                    | Auth + Rate Limiting      |
                    +------------+-------------+
                                 |
                                 v
                    +--------------------------+
                    |    Payment Service       |
                    |       Stateless          |
                    +-------+------------+------+
                            |            |
                       SQL  |            | Outbox Event
                            |            |
                            v            v
                    +------------+   +---------+
                    | Payment DB |   | Outbox  |
                    |   SQL      |   |  Table  |
                    +-----+------+   +----+----+
                                         |
                                         v
                                    +---------+
                                    |  Kafka  |
                                    +----+----+
                                         |
                         +---------------+----------------+
                         |                                |
                         v                                v
                 +---------------+                +---------------+
                 |Payment Workers|                |Webhook Workers|
                 +-------+-------+                +-------+-------+
                         |                                |
               +---------+---------+                      v
               |         |         |                Merchant
               v         v         v                Webhook
          Provider A Provider B Provider C
               |         |         |
               +---------+---------+
                         |
                         v
                 External Networks

                 +----------------------+
                 | Reconciliation       |
                 | Workers              |
                 +----------+-----------+
                            |
                            v
                     Provider APIs

                 +----------------------+
                 | Redis                |
                 | Payment Status Cache |
                 +----------------------+
```

---

# 7. API Design

## Create Payment

```http
POST /v1/payments
Idempotency-Key: abc123
```

Example request:

```json
{
  "merchantId": "merchant_123",
  "amount": 10000,
  "currency": "INR",
  "paymentMethod": {
    "type": "CARD",
    "token": "card_token"
  },
  "returnUrl": "https://merchant.example/callback"
}
```

Response:

```json
{
  "paymentId": "pay_123",
  "status": "PROCESSING"
}
```

## Get Payment

```http
GET /v1/payments/{paymentId}
```

---

# 8. Idempotency

The merchant provides:

```text
Idempotency-Key: abc123
```

Persist it in the database.

Payment table:

```text
payments
------------------------------------------------
payment_id
merchant_id
idempotency_key
amount
currency
payment_method
status
provider
provider_payment_id
created_at
updated_at
version
```

Unique constraint:

```text
UNIQUE(merchant_id, idempotency_key)
```

Flow:

```text
Merchant
   |
   | Idempotency-Key = abc123
   v
Payment Service
   |
   v
INSERT payment
   |
   +---- first request ----> create payment
   |
   +---- duplicate --------> return existing payment
```

### Why not Redis as the source of truth?

Redis can optimize duplicate lookups, but it should not be authoritative.

A short TTL is unsafe because:
- retries can happen after hours;
- applications can crash;
- provider processing can take longer;
- payments can remain unresolved.

The durable DB record is authoritative.

---

# 9. DB + Kafka Consistency: Transactional Outbox

Avoid:

```text
DB write
   |
Kafka publish
```

because DB can succeed while Kafka fails.

Instead:

```text
BEGIN TRANSACTION

INSERT payment
INSERT outbox event

COMMIT
```

Example:

```text
payments
------------------------------------------------
payment_id | status
pay_123    | CREATED
```

```text
outbox
------------------------------------------------
event_id | payment_id | event_type | status
evt_1    | pay_123    | PAYMENT_CREATED | PENDING
```

Then:

```text
Outbox Publisher
       |
       v
     Kafka
```

After successful publication:

```text
outbox.status = PUBLISHED
```

If the publisher crashes, pending events are retried.

This gives at-least-once event publication, so consumers must be idempotent.

---

# 10. Kafka Design

Topic:

```text
payment-events
```

Partition key:

```text
paymentId
```

Why?

Events:

```text
pay_123 CREATED
pay_123 PROCESSING
pay_123 SUCCESS
```

stay ordered within one Kafka partition.

Example:

```text
Partition 0 -> Worker A
Partition 1 -> Worker B
Partition 2 -> Worker C
Partition 3 -> Worker D
```

Consumers belong to the same consumer group.

Scale consumers based on:
- consumer lag
- CPU
- throughput

Important:
- Kubernetes scales consumer pods.
- Kafka does not automatically create more partitions when traffic increases.
- Partition count must be explicitly provisioned/increased.
- Consumer parallelism is bounded by partition count.

---

# 11. Payment Processing

```text
Kafka
  |
  v
Payment Worker
  |
  v
Acquire processing ownership
  |
  v
CREATED -> PROCESSING
  |
  v
Call provider
  |
  +------> SUCCESS
  |
  +------> FAILED
  |
  +------> TIMEOUT / UNKNOWN
```

Use an atomic state transition:

```sql
UPDATE payments
SET status = 'PROCESSING',
    version = version + 1
WHERE payment_id = ?
  AND status = 'CREATED';
```

If one row is updated, the worker owns processing.

If zero rows are updated, another worker already moved the payment forward.

Alternative: `SELECT ... FOR UPDATE` inside a short transaction.

Never hold a DB transaction/row lock while waiting for an external provider.

---

# 12. Provider Idempotency

Suppose:

```text
Worker A
   |
   v
Provider
   |
SUCCESS
   |
Worker A crashes
```

Kafka redelivers.

Worker B must not blindly charge again.

Where supported, send:

```text
Provider Idempotency-Key = paymentId
```

Then duplicate requests with the same key return the same logical result.

---

# 13. Provider Timeout / Ambiguous Result

Critical case:

```text
Worker
  |
  v
Provider
  |
  | customer charged
  |
  X response lost
  |
Worker timeout
```

Do not assume:

```text
timeout = FAILED
```

Keep the payment in `PROCESSING` / `UNKNOWN`.

Then reconciliation checks the provider.

Possible results:

```text
SUCCESS       -> SUCCESS
FAILED        -> FAILED
PENDING       -> retry reconciliation
NOT_FOUND     -> retry only if provider semantics guarantee
                 the original request was never accepted
```

Never blindly retry an ambiguous payment.

---

# 14. Reconciliation

Use distributed workers:

```text
Payment DB
    |
    | PROCESSING / UNKNOWN
    v
Reconciliation Queue
    |
    +---- Worker A
    +---- Worker B
    +---- Worker C
    +---- Worker D
```

Workers claim work atomically.

Example concept:

```sql
UPDATE payments
SET reconciliation_owner = ?,
    reconciliation_until = ?
WHERE payment_id = ?
  AND (
       reconciliation_owner IS NULL
       OR reconciliation_until < NOW()
  );
```

Only the worker that successfully claims the payment processes it.

A delayed queue can also be used for scheduled reconciliation retries.

---

# 15. Reconciliation Retry Policy

Example:

```text
Initial reconciliation
        |
        v
5 minutes
        |
        v
Retry
        |
        v
15 minutes
        |
        v
Retry
        |
        v
1 hour
        |
        v
Retry
        |
        v
Manual intervention / DLQ
```

The exact schedule should depend on provider SLAs and business requirements.

Never immediately mark an ambiguous payment as failed.

---

# 16. Payment Status Read Path

At high read volume:

```text
Merchant
   |
   v
API Gateway
   |
   v
Payment Service
   |
   v
Redis
   |
   +---- HIT ----> return status
   |
   +---- MISS
           |
           v
          DB
           |
           v
        Redis
```

DB remains the source of truth.

When payment changes:

```text
DB -> Kafka -> Cache Updater -> Redis
```

TTL can provide an additional safety mechanism, but correctness should not rely only on TTL.

---

# 17. Database Architecture

Use SQL initially because payment data requires:
- transactions
- unique constraints
- strong consistency
- atomic state transitions
- durable storage
- indexing

Possible architecture:

```text
                Primary DB
                   |
          +--------+--------+
          |                 |
     Read Replica       Backups
```

Useful indexes:

```text
PRIMARY KEY(payment_id)

UNIQUE(merchant_id, idempotency_key)

INDEX(merchant_id, created_at)

INDEX(status, updated_at)
```

For larger scale, partition/shard according to query patterns.

Possible partition keys:
- merchant_id
- hash(payment_id)

Avoid automatically assigning one shard per merchant because large merchants can create hot shards.

---

# 18. Webhook Architecture

```text
Payment Worker
      |
      v
Kafka
      |
      v
Webhook Worker
      |
      v
Merchant Endpoint
```

Webhook:

```json
{
  "eventId": "evt_123",
  "type": "PAYMENT_SUCCEEDED",
  "paymentId": "pay_123",
  "timestamp": 1750000000
}
```

Delivery semantics:

**At-least-once.**

Exactly-once delivery to an external merchant cannot be guaranteed.

Example:

```text
Webhook Worker
      |
      v
Merchant
      |
      v
Merchant processes event
      |
Network timeout
      |
Webhook Worker doesn't receive 200
```

Retry occurs.

Merchant may receive:

```text
evt_123
evt_123
evt_123
```

Merchant deduplicates using `eventId` with a unique constraint.

---

# 19. Webhook Retry

Do not keep a Kafka message unacknowledged for hours.

Use retry scheduling:

```text
Attempt 1
   |
 failure
   v
5 min
   |
Attempt 2
   |
 failure
   v
15 min
   |
Attempt 3
   |
 failure
   v
1 hour
   |
Attempt 4
   |
 ...
   v
DLQ
```

Use:
- exponential backoff
- jitter
- maximum retry duration/count
- DLQ

An unavailable merchant should not block a Kafka partition.

---

# 20. Webhook Security

Sign webhook requests:

```text
signature = HMAC(secret, payload + timestamp)
```

Include:
- eventId
- timestamp
- signature

Merchant verifies:
1. signature
2. timestamp/replay window
3. eventId deduplication

---

# 21. Authentication and Authorization

API Gateway can handle:
- TLS termination
- authentication
- rate limiting
- request validation
- routing

Authorization must also be enforced by the service.

A merchant can access only its own payments.

Use:
- API keys / OAuth as appropriate
- merchant identity propagation
- authorization checks
- encryption in transit
- encryption at rest
- secrets manager

---

# 22. PCI / Sensitive Payment Data

Avoid storing raw card data when possible.

Use tokenization:

```text
Merchant
   |
Card data
   |
Tokenization
   |
Payment Token
   |
PSP
```

Use:
- TLS
- encryption at rest
- secret management
- audit logging
- access controls

---

# 23. Provider Routing

```text
                Payment Router
               /      |                     /       |                Provider A Provider B Provider C
```

Routing can consider:
- payment method
- country
- currency
- provider availability
- provider success rate
- processing cost
- merchant configuration
- regional coverage

Do not blindly fail over after an ambiguous provider timeout.

Reconcile or use provider idempotency first.

---

# 24. Reliability

## Timeouts
Every external call has a bounded timeout.

## Retries
Use:
- exponential backoff
- jitter
- maximum retry count

Only retry operations that are safe to retry.

## Circuit breaker

```text
Provider A
    |
high failure rate
    |
Circuit OPEN
    |
stop sending traffic temporarily
```

## Bulkheads
Prevent one provider or merchant from consuming all resources.

## Backpressure
Kafka buffers work during downstream traffic spikes.

---

# 25. Scaling to 10K TPS

API layer:

```text
                 Load Balancer
                      |
          +-----------+-----------+
          |           |           |
       Payment     Payment     Payment
       Service     Service     Service
```

Scale stateless services horizontally.

Kafka:
- enough partitions for required parallelism
- consumer groups
- monitor consumer lag

Database:
- indexes
- batching where appropriate
- partitioning/sharding
- connection pooling
- read replicas
- careful transaction boundaries

Redis:
- absorb high-volume status reads

---

# 26. Observability

Monitor:

## API
- request rate
- p50/p95/p99 latency
- error rate
- availability

## Kafka
- consumer lag
- throughput
- partition utilization
- DLQ size

## Database
- CPU
- connections
- lock contention
- query latency
- replication lag

## Payments
- SUCCESS rate
- FAILED rate
- PROCESSING duration
- UNKNOWN payments
- reconciliation backlog

## Providers
- latency
- timeout rate
- failure rate
- success rate

## Webhooks
- delivery success rate
- retry rate
- merchant endpoint latency
- DLQ size

---

# 27. Failure Scenarios

### DB fails during payment creation
Do not acknowledge creation unless it is durably persisted.

### Kafka fails
Transactional outbox retains the event.

### Outbox publisher crashes
Pending outbox records are retried.

### Payment worker crashes
Kafka redelivers the event.

### Provider times out
Keep payment ambiguous/processing and reconcile.

### Provider succeeds but worker crashes
Provider idempotency + reconciliation.

### Merchant retries
DB unique constraint on `(merchant_id, idempotency_key)` returns the existing payment.

### Duplicate Kafka event
Consumer checks state/event ID and processes idempotently.

### Webhook timeout
At-least-once retry.

### Redis unavailable
Fall back to DB.

### Provider unavailable
Circuit breaker + routing/failover where safe.

---

# 28. Capacity Discussion

For ~10K payment requests/sec:

API:
- stateless horizontal scaling

Kafka:
- enough partitions to support consumer throughput

DB:
- benchmark actual write/read throughput
- account for payment writes, state transitions and outbox writes
- use indexing and partitioning when needed

Redis:
- handle high-volume status reads

Avoid inventing exact server counts without benchmark data.

---

# 29. Trade-offs

## SQL vs NoSQL

SQL initially because payment state needs:
- transactions
- unique constraints
- strong consistency
- atomic updates
- durable records

## Redis

Read optimization, not authoritative payment state.

## Kafka

Useful for:
- asynchronous processing
- buffering
- decoupling
- replay
- horizontal consumers

## Transactional Outbox

Adds complexity but solves the DB/Kafka dual-write problem.

## Strong vs Eventual Consistency

Strong:
- payment state
- idempotency
- payment ownership/state transitions

Eventual:
- cache
- webhooks
- analytics

---

# 30. 60-Minute Interview Strategy

```text
0–5 min
Requirements + assumptions

5–10 min
Scale + APIs

10–20 min
High-level architecture

20–30 min
Payment flow + DB + idempotency

30–40 min
Kafka + processing + provider integration

40–50 min
Failures + reconciliation + webhooks

50–55 min
Scaling + consistency

55–60 min
Security + trade-offs + final questions
```

Get a simple architecture onto the board quickly. Iterate instead of trying to perfect the first diagram.

---

# 31. Senior-Level Statements Worth Remembering

### Idempotency
"The merchant's idempotency key is persisted with a unique database constraint. Redis can optimize duplicate lookups but isn't the source of truth."

### Transactional Outbox
"I don't want a database commit followed by a potentially failing Kafka publish, so I'll use a transactional outbox."

### External Timeout
"A provider timeout doesn't imply failure because the provider may have processed the payment. I'll reconcile before retrying."

### Concurrency
"I'll use an atomic state transition from CREATED to PROCESSING so only one worker can acquire processing ownership."

### Locks
"I won't hold a database transaction open while calling an external provider."

### Kafka
"I'll partition by paymentId when ordering is required per payment."

### Webhooks
"Webhook delivery will be at-least-once. Exactly-once delivery to an external merchant can't be guaranteed, so we'll include an eventId and require idempotent webhook handling."

### Database
"The database is the source of truth for payment state."

### Redis
"Redis is an optimization layer; if it's unavailable, we can fall back to the database."

---

# 32. Final Architecture

```text
                              MERCHANT
                                  |
                                  v
                         +----------------+
                         | API Gateway/LB |
                         | Auth + Limits  |
                         +-------+--------+
                                 |
                                 v
                         +---------------+
                         | Payment       |
                         | Service       |
                         | Stateless     |
                         +-------+-------+
                                 |
                       +---------+---------+
                       |                   |
                       v                   v
                 +-----------+       +-----------+
                 | Payment DB|       |   Redis   |
                 |   SQL     |       |   Cache   |
                 +-----+-----+       +-----------+
                       |
                  DB Transaction
                       |
                 +-----v------+
                 |   Outbox   |
                 +-----+------+
                       |
                       v
                  +---------+
                  |  Kafka  |
                  +----+----+
                       |
          +------------+-------------+
          |                          |
          v                          v
   +--------------+           +-------------+
   | Payment      |           | Webhook     |
   | Workers      |           | Workers     |
   +------+-------+           +------+------+
          |                          |
          v                          v
   +-------------+             Merchant
   | Payment     |             Webhook
   | Router      |
   +------+------+ 
          |
    +-----+-----+-----+
    |           |     |
    v           v     v
 Provider A Provider B Provider C
    |           |     |
    +-----+-----+-----+
          |
          v
    Payment Networks


              +----------------------+
              | Reconciliation       |
              | Workers              |
              +----------+-----------+
                         |
                         v
                  Provider APIs
```

---

# 33. Five Concepts to Prioritize

If time is limited, focus on:

1. **Merchant idempotency key + DB unique constraint**
2. **Transactional Outbox for DB → Kafka**
3. **Atomic payment state transitions**
4. **Provider idempotency + reconciliation for ambiguous outcomes**
5. **At-least-once webhooks + merchant-side deduplication**

---

# 34. Common Mistakes to Avoid

Avoid saying:

- "Redis will guarantee idempotency."
- "A timeout means payment failed."
- "Kafka guarantees exactly-once end-to-end."
- "We can hold the DB lock while calling the provider."
- "We'll just retry the payment on another provider."
- "One merchant gets one DB shard."
- "Kubernetes automatically increases Kafka partitions."
- "We can guarantee exactly-once webhook delivery."
- "We'll add Kafka/Redis/RabbitMQ" without explaining why.

Always connect a technology choice to a requirement or invariant.

---

# 35. Final Mental Model

Think of the PSP as four major flows:

```text
1. ACCEPT
Merchant
   ↓
API
   ↓
Idempotent DB write


2. PROCESS
DB/Outbox
   ↓
Kafka
   ↓
Worker
   ↓
Provider


3. RECONCILE
Unknown/Processing
   ↓
Reconciliation
   ↓
Provider status
   ↓
Final state


4. NOTIFY
Payment state change
   ↓
Kafka
   ↓
Webhook
   ↓
Merchant
```

The database is the **source of truth**.

Kafka provides **asynchronous decoupling and buffering**.

Redis provides **read optimization**.

The outbox solves **DB/Kafka dual writes**.

Provider idempotency + reconciliation solve **ambiguous external payment outcomes**.

Webhooks use **at-least-once delivery and event-level idempotency**.

This is the core design to be able to explain clearly in an Adyen Senior Backend Engineer system-design interview.
