# Production Planning Dashboard

A Bill of Materials (BOM) management system built in Go for tracking components, calculating production costs, and managing product assemblies.

## Features

### Component Management
- Create and store component definitions with specifications
- Track unit costs, lead times, and units of measurement
- View and list all registered components
- Delete components from the registry

### BOM (Bill of Materials) Management
- Build multi-level BOMs for products
- Define component quantities for each product
- Automatic cost calculation based on component prices
- View detailed BOM breakdowns with total costs
- Manage multiple BOMs simultaneously

## Installation

### Prerequisites
- Go 1.21 or higher

### Setup
```bash
# Clone the repository
git clone https://github.com/yourusername/production-planner.git
cd production-planner

# Initialize Go module (if not already done)
go mod init github.com/yourusername/production-planner

# Run the application
go run .
```

## Usage

### Starting the Application
```bash
go run .
```

You'll be presented with the main menu:
```
1. BOM Menu
2. Components Menu
3. Exit
```

### Component Workflow

1. **Add a Component**
   - Navigate to Components Menu → Add Component
   - Enter component details:
     - ID (unique identifier)
     - Name
     - Description
     - Unit of Measure (e.g., "pcs", "kg", "m")
     - Unit Cost (in your currency)
     - Lead Time Days

2. **View/List Components**
   - View a specific component by ID
   - List all registered components

### BOM Workflow

1. **Create a BOM**
   - Navigate to BOM Menu → Create BOM
   - Enter Product ID and Product Name

2. **Add Items to BOM**
   - Select BOM → Add Item
   - Enter Component ID and Quantity
   - System automatically calculates costs

3. **View BOM Details**
   - See all components, quantities, and total cost
   - Useful for cost estimation and planning

## Example Usage

### Creating a Bicycle BOM

```
1. Create Components:
   - ID: "frame-01", Name: "Aluminum Frame", Cost: 150.00, Unit: "pcs"
   - ID: "wheel-01", Name: "26\" Wheel", Cost: 45.00, Unit: "pcs"
   - ID: "brake-01", Name: "Disc Brake Set", Cost: 80.00, Unit: "set"

2. Create BOM:
   - Product ID: "bike-001"
   - Product Name: "Mountain Bike Basic"

3. Add Items:
   - frame-01: quantity 1
   - wheel-01: quantity 2
   - brake-01: quantity 1

4. View BOM:
   Product: Mountain Bike Basic (ID: bike-001)
   - Aluminum Frame x1: $150.00
   - 26" Wheel x2: $90.00
   - Disc Brake Set x1: $80.00
   Total Cost: $320.00
```

## Project Structure

```
production-planner/
├── main.go           # Main entry point and menu system
├── component.go      # Component struct and registry
├── bom.go           # BOM struct and management
├── utils.go         # Input utilities
└── README.md        # This file
```

## Learning Journey

This project was built while learning Go from **"Learning Go: An Idiomatic Approach to Real-World Go Programming"** by Jon Bodner.

### Concepts Applied (Chapters 1-7)
- ✅ Structs and composite types
- ✅ Pointers and value vs. reference semantics
- ✅ Methods and receivers
- ✅ Error handling patterns
- ✅ Maps for registry pattern
- ✅ Slices for collections
- ✅ Constructor functions
- ✅ Comma-ok idiom

### Roadmap
- **Phase 2** (After Chapter 13): Add HTTP server and web UI
- **Phase 3** (After Chapter 15): Add comprehensive test suite
- **Phase 4**: Package as desktop application using Wails
- **Future**: Multi-level BOMs (nested assemblies), cost analysis reports, export to CSV

## Technical Decisions

### Why Pointers?
- `*Component` and `*BOM` are returned from constructors to allow modification
- Avoids copying large structs
- Enables shared references in BOM items

### Registry Pattern
- Centralized storage for components and BOMs
- Fast lookups using map with string keys
- Prevents duplicate IDs

### Error Handling
- Validation at creation time (constructors)
- Graceful error propagation
- User-friendly error messages

## Contributing

This is a learning project, but suggestions and improvements are welcome!

## License

MIT License - feel free to use this for learning or as a starting point for your own projects.

## Acknowledgments

- Built following idiomatic Go practices from "Learning Go" by Jon Bodner
- Inspired by real-world manufacturing and operations research problems
- Thanks to the Go community for excellent documentation and tools

---

**Note:** This is a CLI application. Web and desktop interfaces are planned for future releases.
