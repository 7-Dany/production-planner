# Production Planning Dashboard

A Bill of Materials (BOM) management system built in Go for tracking components, calculating production costs, and managing product assemblies.

## 📦 Download & Installation

### Option 1: Download Pre-built Binary (Recommended)

**Windows Users (64-bit)** - Most Common ⭐
1. Go to [Releases](https://github.com/yourusername/production-planner/releases/latest)
2. Download: **`production-planner-windows-amd64.exe`**
3. Double-click to run

**Note:** Windows might show a "Windows protected your PC" warning. Click "More info" → "Run anyway". This happens because the app isn't digitally signed yet.

**macOS Users**
- **Intel Mac:** Download `production-planner-darwin-amd64`
- **M1/M2/M3 Mac:** Download `production-planner-darwin-arm64`

```bash
# Make executable
chmod +x production-planner-darwin-amd64
# Run it
./production-planner-darwin-amd64
```

**Linux Users**
- **64-bit:** Download `production-planner-linux-amd64`
- **ARM (Raspberry Pi):** Download `production-planner-linux-arm64`

```bash
chmod +x production-planner-linux-amd64
./production-planner-linux-amd64
```

**Not sure which version?**
- **Windows:** If your PC was made after 2010, use `windows-amd64.exe`
- **macOS:** Run `uname -m` in Terminal
  - `arm64` → Use `darwin-arm64`
  - `x86_64` → Use `darwin-amd64`
- **Linux:** Run `uname -m`
  - `x86_64` → Use `linux-amd64`
  - `aarch64` → Use `linux-arm64`

---

### Option 2: Build from Source

**Prerequisites:** Go 1.21 or higher

```bash
# Clone the repository
git clone https://github.com/yourusername/production-planner.git
cd production-planner

# Build
go build .

# Run
./production-planner  # or production-planner.exe on Windows
```

---

## ✨ Features

### Component Management
- ✅ Create and store component definitions with specifications
- ✅ Track unit costs, lead times, and units of measurement
- ✅ View and list all registered components
- ✅ Delete components with confirmation dialog
- ✅ ID validation (no spaces, duplicates prevented)

### BOM (Bill of Materials) Management
- ✅ Create BOMs for products
- ✅ Add components with quantities
- ✅ Automatic cost calculation (quantity × unit cost)
- ✅ Professional formatted output with detailed breakdowns
- ✅ View total costs and line items
- ✅ Manage multiple BOMs simultaneously

---

## 🚀 Quick Start

### Starting the Application

Run the downloaded binary or:
```bash
go run .
```

You'll see:
```
╔════════════════════════════════════════════╗
║  Production Planning Dashboard v0.0.1      ║
╚════════════════════════════════════════════╝

System Status:
  • Components: 0 registered
  • BOMs: 0 active

1. BOM Menu
2. Components Menu
3. Exit
```

---

## 📖 Usage Guide

### Component Workflow

#### 1. Add a Component
1. Select **Components Menu** (option 2)
2. Select **Add Component** (option 1)
3. Enter details:
   - **ID:** Unique identifier (no spaces, e.g., "frame-01")
   - **Name:** Component name
   - **Description:** Brief description
   - **Unit of Measure:** e.g., "pcs", "kg", "m", "set"
   - **Unit Cost:** Price per unit
   - **Lead Time Days:** Days to procure

#### 2. View Components
- **View one:** Option 2, enter ID
- **List all:** Option 3

#### 3. Delete Component
- Option 4, enter ID
- Confirmation required before deletion

---

### BOM Workflow

#### 1. Create a BOM
1. Select **BOM Menu** (option 1)
2. Select **Create BOM** (option 2)
3. Enter:
   - **Product ID:** Unique identifier
   - **Product Name:** Name of the product

#### 2. Add Components to BOM
1. Select **Select BOM** (option 3)
2. Enter BOM ID
3. Select **Add Item** (option 2)
4. Enter:
   - **Component ID:** Must exist in component registry
   - **Quantity:** Amount needed

#### 3. View BOM Details
- **From BOM menu:** Option 4, enter BOM ID
- **From BOM submenu:** Option 1 (View Details)

**Example Output:**
```
╔═══════════════════════════════════════════════════════╗
║ Product: Mountain Bike Basic                          ║
║ ID: bike-001                                          ║
╠═══════════════════════════════════════════════════════╣
║ Components:                                           ║
║  1. Aluminum Frame       1.00 pcs × $150.00 = $150.00 ║
║  2. 26" Wheel            2.00 pcs × $45.00  = $90.00  ║
║  3. Disc Brake Set       1.00 set × $80.00  = $80.00  ║
╠═══════════════════════════════════════════════════════╣
║                                         Total: $320.00║
╚═══════════════════════════════════════════════════════╝
```

---

## 💡 Example: Creating a Bicycle BOM

### Step 1: Create Components
```
Components Menu → Add Component

Component 1:
  ID: frame-01
  Name: Aluminum Frame
  Description: Lightweight aluminum alloy
  Unit of Measure: pcs
  Unit Cost: 150.00
  Lead Time Days: 14

Component 2:
  ID: wheel-01
  Name: 26" Mountain Wheel
  Description: All-terrain wheel with tire
  Unit of Measure: pcs
  Unit Cost: 45.00
  Lead Time Days: 7

Component 3:
  ID: brake-01
  Name: Hydraulic Disc Brake Set
  Description: Front and rear brake system
  Unit of Measure: set
  Unit Cost: 80.00
  Lead Time Days: 10
```

### Step 2: Create BOM
```
BOM Menu → Create BOM
  Product ID: bike-001
  Product Name: Mountain Bike Basic
```

### Step 3: Add Components to BOM
```
BOM Menu → Select BOM → Enter "bike-001"
  Add Item → Component ID: frame-01, Quantity: 1
  Add Item → Component ID: wheel-01, Quantity: 2
  Add Item → Component ID: brake-01, Quantity: 1
```

### Step 4: View Results
```
View Details → Shows total cost: $320.00
```

---

## 📁 Project Structure

```
production-planner/
├── main.go              # Entry point, main menu, welcome banner
├── component.go         # Component struct, ComponentRegistry, menu
├── bom.go              # BOM struct, BOMRegistry, BOMItem, menu
├── utils.go            # Input utilities (getInput helper)
├── go.mod              # Go module definition
├── README.md           # This file
├── CHANGELOG.md        # Version history
└── .github/
    └── workflows/
        └── release.yml # Automated binary releases
```

---

## 🎓 Learning Journey

This project was built while learning Go from **"Learning Go: An Idiomatic Approach to Real-World Go Programming"** by Jon Bodner.

### Concepts Applied (Chapters 1-7)
- ✅ **Structs and composite types** - Component, BOM, BOMItem structures
- ✅ **Pointers** - Constructor functions return pointers for efficiency
- ✅ **Methods and receivers** - String(), AddItem(), etc.
- ✅ **Error handling** - Validation and graceful error propagation
- ✅ **Maps** - Registry pattern for fast lookups
- ✅ **Slices** - Dynamic collections of components and BOM items
- ✅ **Constructor functions** - NewComponent(), NewBOM() with validation
- ✅ **Comma-ok idiom** - Safe map access patterns

---

## 🗺️ Roadmap

### Current: Phase 1 ✅ (v0.0.1)
- Single-level BOM
- Component management
- CLI interface
- Cross-platform binaries

### Phase 2 (After Chapter 8-10)
- 🔲 Multi-level BOM support (sub-assemblies)
- 🔲 Generic registry refactoring
- 🔲 Custom error types
- 🔲 Package reorganization

### Phase 3 (After Chapter 13)
- 🔲 REST API with net/http
- 🔲 Web UI (htmx + TailwindCSS)
- 🔲 JSON import/export

### Phase 4 (After Chapter 15)
- 🔲 Database persistence (SQLite/PostgreSQL)
- 🔲 Comprehensive test suite
- 🔲 Where-used reports
- 🔲 BOM comparison/versioning

### Phase 5 (Future)
- 🔲 Desktop application (Wails)
- 🔲 Inventory management module
- 🔲 MRP calculations
- 🔲 Supplier management
- 🔲 Advanced reporting

---

## 🏗️ Technical Decisions

### Design Patterns
- **Registry Pattern:** Centralized storage with map-based O(1) lookups
- **Constructor Functions:** Validation at creation time
- **Pointer Receivers:** Methods that modify state use pointer receivers
- **Error Propagation:** Errors returned up the call stack, not hidden

### Why Pointers?
- `*Component` and `*BOM` returned from constructors
- Avoids copying large structs
- Enables shared references in BOM items
- Allows modification through methods

### Data Storage (Current)
- In-memory maps (no persistence yet)
- Fast lookups: O(1) by ID
- Data lost on exit (database coming in Phase 4)

### Code Quality
- Idiomatic Go conventions followed
- Error messages: lowercase, no periods
- Input validation at boundaries
- Confirmation dialogs for destructive operations
- Professional formatted output

---

## 📄 License

MIT License - feel free to use this for learning or as a starting point for your own projects.

---

## 🙏 Acknowledgments

- Built following idiomatic Go practices from **"Learning Go"** by Jon Bodner
- Inspired by real-world manufacturing and operations research problems
- Domain knowledge from mechanical engineering and operations management
- Thanks to the Go community for excellent documentation and tools

---

## 🏷️ Version

**Current Version:** v0.0.1 - Initial Release

See [CHANGELOG.md](CHANGELOG.md) for version history.

---

**⚠️ Note:** This is currently a CLI application. Web and desktop interfaces are planned for future releases. Data is stored in-memory and will be lost when the application closes until database persistence is added in Phase 4.
