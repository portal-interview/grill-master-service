# Grill Master Service

The GRILL MASTER runs the show! Every order that comes into Flavortown gets processed and routed right here.

This Go service is the heart of our ordering pipeline. When a customer places an order through the Flavor Portal or Triple D Mobile, it lands here first. The Grill Master validates the order, determines which kitchen stations need to handle which items, and routes everything to the Kitchen API for fulfillment.

## Responsibilities

- **Order Validation** — Ensures orders are complete and items are available
- **Station Routing** — Intelligently assigns orders to kitchen stations based on item type and current load
- **Order Lifecycle** — Tracks orders from received through delivered
- **Station Management** — Monitors station capacity and status

## Why Go?

We chose Go for the Grill Master because order processing needs to be FAST. When the Friday night rush hits, we're processing hundreds of orders per minute. Go's concurrency model and low-latency networking make it perfect for this high-throughput service.
