# LLD Practice Question: Design an In-Memory Notification Service

## Problem Statement

Design an in-memory notification service that can send notifications through different communication channels.

The system should be modular, extensible, and easy to maintain. New notification channels should be added with minimal code changes.

---

## Functional Requirements

The system should support the following operations:

### 1. Send a Notification

Send a notification to a recipient through the specified notification channel.

### 2. Retry a Notification

Retry sending a notification that previously failed.

### 3. Get a Notification

Retrieve a notification by its unique ID.

### 4. Get All Notifications

Retrieve all notifications currently stored in the system.

---

## Notification Model

Each notification should contain the following fields:

| Field | Description |
|-------|-------------|
| ID | Unique identifier |
| Recipient | Recipient of the notification |
| Message | Notification content |
| Channel | Delivery channel |
| Priority | Notification priority |
| Timestamp | Creation timestamp |
| Status | Current notification status |

---

## Supported Channels

Initially, the system should support the following channels:

- Email
- SMS
- Push Notification

Each channel should have its own implementation for sending notifications.

---

## Notification Status

A notification can have one of the following statuses:

- `PENDING`
- `SENT`
- `FAILED`

Rules:

- Every newly created notification starts with the **PENDING** status.
- If sending succeeds, the status becomes **SENT**.
- If sending fails, the status becomes **FAILED**.

---

## Notification Priority

Support the following priorities:

- `LOW`
- `MEDIUM`
- `HIGH`

For this exercise, simply store the priority with the notification. Processing notifications based on priority is **not required**.

---

## Storage

Store all notifications **in memory**.

No database or persistent storage is required.

---

## Constraints

- Everything should be implemented in memory.
- No REST APIs or UI are required.
- No external email/SMS/push providers need to be integrated.
- Sending a notification can be simulated by printing a message to the console.
- Focus on clean object-oriented design rather than production-level networking.

---

## Future Requirements

Your design should make it easy to support additional notification channels in the future, such as:

- WhatsApp
- Slack
- Discord
- Microsoft Teams

Adding a new notification channel should require **minimal or no modification** to the existing business logic.

---

## Expectations

Your solution should demonstrate:

- Clean object-oriented design
- Separation of concerns
- Proper use of interfaces and abstractions
- Extensibility
- Maintainability
- Readable and testable code

You are free to choose the classes, interfaces, and design patterns that best fit the problem.

---

## Bonus (Optional)

If time permits, implement one or more of the following:

- Retry count for failed notifications.
- Search notifications by recipient.
- Filter notifications by status.
- Filter notifications by channel.
- Delete a notification.
- Thread-safe implementation.
- Unit tests.

---

## Estimated Time

**45–60 minutes**