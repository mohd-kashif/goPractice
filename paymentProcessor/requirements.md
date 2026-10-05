# LLD Practice Question: Design an In-Memory Payment Processing System

## Problem Statement

Design an in-memory payment processing system that supports multiple payment methods.

The system should be modular, extensible, and easy to maintain. New payment methods should be added with minimal changes to the existing code.

---

# Functional Requirements

## 1. Register a Payment Method

The system should support registering different payment methods.

Initially support:

- Credit Card
- Debit Card
- UPI
- Net Banking

---

## 2. Process a Payment

Process a payment using the specified payment method.

The system should validate the request and invoke the appropriate payment processor.

A successful payment should be recorded.

A failed payment should also be recorded.

---

## 3. Refund a Payment

Refund a previously successful payment.

A payment cannot be refunded more than once.

---

## 4. Get Payment Details

Retrieve the details of a payment using its payment ID.

---

## 5. List All Payments

Return all payments processed by the system.

---

## 6. Get Payment Status

Retrieve the current status of a payment.

Possible statuses include:

- INITIATED
- SUCCESS
- FAILED
- REFUNDED

---

# Constraints

- Everything should be implemented in memory.
- No external payment gateways.
- No database.
- No REST APIs.
- Payment processing can be simulated by printing logs or returning success/failure.
- Focus on clean object-oriented design.

---

# Future Requirements

Your design should make it easy to support additional payment methods without modifying existing business logic.

Examples:

- Wallet
- PayPal
- Apple Pay
- Google Pay
- Cryptocurrency

---

# Expectations

Your solution should demonstrate:

- Clean object-oriented design
- Separation of concerns
- Proper use of interfaces
- Extensibility
- Maintainability
- Readable and testable code

Choose the classes, interfaces, and design patterns you believe are appropriate.

---

# Bonus (Optional)

If time permits, implement one or more of the following:

- Payment history for a user
- Search payments by status
- Search payments by payment method
- Retry failed payments
- Thread-safe repositories
- Unit tests

---

# Estimated Time

45–60 minutes