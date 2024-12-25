├── models/           # Core data models (new directory)
│   ├── response.go   # API response structures
│   └── shipment.go   # Shipment data structures
pkg/llm/
├── engine/
│   ├── engine.go      # Core tinygrad integration interface
│   ├── processor.go   # Text processing and tokenization
│   ├── inference.go   # Model inference implementation
│   └── config.go      # Engine configuration
│
├── models/
│   ├── model.go       # Model interface definitions
│   ├── weights.go     # Weight management
│   ├── embedding.go   # Embedding layer implementations
│   └── attention.go   # Attention mechanism implementations
│
└── training/
    ├── trainer.go     # Training orchestration
    ├── dataset.go     # Training data management
    ├── optimizer.go   # Optimization algorithms
    └── metrics.go     # Training metrics and evaluation
    pkg/usps/
├── client/
│   ├── client.go      # USPS API client interface
│   ├── auth.go        # Authentication handling
│   ├── request.go     # Request building
│   └── response.go    # Response parsing
│
├── rates/
│   ├── calculator.go  # Rate calculation interface
│   ├── domestic.go    # Domestic shipping rates
│   ├── international.go # International rates
│   └── zones.go       # Shipping zone calculations
│
└── tracking/
    ├── tracker.go     # Tracking interface
    ├── status.go      # Shipment status management
    ├── events.go      # Tracking event handling
    └── notifications.go # Status notification system
    pkg/core/
├── config/
│   ├── config.go      # Configuration management
│   ├── environment.go # Environment variable handling
│   ├── validation.go  # Config validation
│   └── defaults.go    # Default configurations
│
├── logger/
│   ├── logger.go      # Logging interface
│   ├── formatter.go   # Log formatting
│   ├── levels.go      # Log levels
│   └── output.go      # Log output management
│
└── errors/
    ├── errors.go      # Custom error definitions
    ├── handlers.go    # Error handling
    ├── recovery.go    # Panic recovery
    └── codes.go       # Error codes and messages
pkg/analytics/
├── tracking/
│   ├── tracker.go     # Usage tracking interface
│   ├── events.go      # Event definition and handling
│   ├── collector.go   # Data collection
│   └── storage.go     # Analytics storage
│
└── reporting/
    ├── reporter.go    # Reporting interface
    ├── aggregator.go  # Data aggregation
    ├── formatter.go   # Report formatting
    └── exporter.go    # Report export functionality