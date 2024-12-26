[]		smail/
[x]		├── cmd/
[x]		│   └── server/
[x]		│       ├── main.go        # Application entry point
[x]		│       ├── routes.go      # Route configurations
[x]		│       └── handlers.go    # HTTP request handlers
[]		│
[x]		├── pkg/                   # Go packages
[x]		│   ├── models/           # Shared data structures
[]		│   │   ├── response.go   # API response structures
[]		│   │   └── shipment.go   # Shipment data models
[]		│   │
[x]		│   ├── mail/            # Email processing
[]		│   │   ├── parser/
[]		│   │   │   ├── parser.go      # Main parsing interface
[]		│   │   │   ├── mime.go        # MIME handling
[]		│   │   │   ├── attachment.go  # Attachment processing
[]		│   │   │   └── types.go       # Type definitions
[]		│   │   ├── validator/
[]		│   │   │   ├── validator.go   # Main validation
[]		│   │   │   ├── syntax.go      # Syntax rules
[]		│   │   │   ├── domain.go      # Domain validation
[]		│   │   │   └── rules.go       # Custom rules
[]		│   │   └── sender/
[]		│   │       ├── sender.go      # Sending interface
[]		│   │       ├── smtp.go        # SMTP implementation
[]		│   │       ├── queue.go       # Queue management
[]		│   │       └── retry.go       # Retry logic
[]		│   │
[x]		│   ├─pkg/llm/           # LLM package in Go
[x]			├── engine/        # Core LLM engine components
[]			│   ├── engine.go      # Main engine interface
[]			│   ├── processor.go   # Text processing
[]			│   ├── inference.go   # Model inference
[]			│   └── config.go      # Engine configuration
[]			│
[x]			├── models/        # Model definitions
[]			│   ├── model.go       # Model interface
[]			│   ├── weights.go     # Weight management
[]			│   ├── embedding.go   # Embedding implementations
[]			│   └── attention.go   # Attention mechanisms
[]			│
[x]			├── training/      # Training orchestration
[]			│   ├── trainer.go     # Training interface
[]			│   ├── dataset.go     # Training data
[]			│   ├── optimizer.go   # Optimization
[]			│   └── metrics.go     # Training metrics
[]			│
[x]			├── bridge/        # Bridge to Python implementation
[]			│   ├── bridge.go      # Bridge interface
[]			│   └── client.go      # HTTP client
[]			│
[x]			└── types/         # Shared type definitions
[]		    	└── types.go       # Common types
[]		│   │
[x]		│   ├── usps/            # USPS integration
[x]		│   │   ├── client/
[]		│   │   │   ├── client.go    # USPS API client
[]		│   │   │   ├── auth.go      # Authentication
[]		│   │   │   ├── request.go   # Request building
[]		│   │   │   └── response.go  # Response parsing
[x]		│   │   ├── rates/
[]		│   │   │   ├── calculator.go     # Rate calculation
[]		│   │   │   ├── domestic.go       # Domestic rates
[]		│   │   │   ├── international.go  # International
[]		│   │   │   └── zones.go          # Zone logic
[x]		│   │   └── tracking/
[]		│   │       ├── tracker.go        # Tracking interface
[]		│   │       ├── status.go         # Status management
[]		│   │       ├── events.go         # Event handling
[]		│   │       └── notifications.go  # Notifications
[]		│   │
[x]		│   ├── core/            # Core functionality
[]		│   │   ├── config/
[]		│   │   │   ├── config.go      # Configuration
[]		│   │   │   ├── environment.go # Env variables
[]		│   │   │   ├── validation.go  # Config validation
[]		│   │   │   └── defaults.go    # Defaults
[x]		│   │   ├── logger/
[]		│   │   │   ├── logger.go      # Logging interface
[]		│   │   │   ├── formatter.go   # Log formatting
[]		│   │   │   ├── levels.go      # Log levels
[]		│   │   │   └── output.go      # Output handling
[x]		│   │   └── errors/
[]		│   │       ├── errors.go      # Error definitions
[]		│   │       ├── handlers.go    # Error handling
[]		│   │       ├── recovery.go    # Panic recovery
[]		│   │       └── codes.go       # Error codes
[]		│   │
[x]		│   └── analytics/       # Analytics system
[x]		│       ├── tracking/
[]		│       │   ├── tracker.go     # Usage tracking
[]		│       │   ├── events.go      # Event handling
[]		│       │   ├── collector.go   # Data collection
[]		│       │   └── storage.go     # Data storage
[x]		│       └── reporting/
[]		│           ├── reporter.go    # Reporting interface
[]		│           ├── aggregator.go  # Data aggregation
[]		│           ├── formatter.go   # Report formatting
[]		│           └── exporter.go    # Export logic
[]		│
[x]		├── python/              # Python/tinygrad implementation
[x]		│   └── llm/
[x]		│       ├── engine/
[]		│       │   ├── __init__.py
[]		│       │   ├── model.py      # tinygrad model
[]		│       │   └── trainer.py    # Training logic
[x]		│       ├── serving/
[]		│       │   ├── __init__.py
[]		│       │   └── server.py     # FastAPI server
[]		│       └── requirements.txt
[]		│
[x]		├── web/                # Frontend assets
[x]		│   ├── static/
[]		│   │   ├── css/
[]		│   │   ├── js/
[]		│   │   └── images/
[x]		│   └── templates/
[]		│       └── layout/
[]		│
[x]		├── deployments/        # Deployment configurations
[]		│   ├── docker/
[]		│   └── kubernetes/
[]		│
[]		├── docs/              # Documentation
[]		├── tests/             # Integration tests
[]		├── .env.example       # Environment template
[]		├── docker-compose.yml # Docker composition
[]		├── Dockerfile         # Main dockerfile
[]		└── README.md          # Project documentation
