# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
- Multi-level BOM support
- Generic registry implementation
- Package reorganization

## [0.0.1] - 2025-01-14

### Added
- Component management system
  - Create components with ID, name, description, unit cost, unit of measure, lead time
  - View individual components
  - List all components
  - Delete components with confirmation
- BOM management system
  - Create BOMs with product ID and name
  - Add components with quantities
  - View BOM with line items and total cost
  - Delete BOMs
- Registry pattern for centralized storage
- Input validation (ID format, positive costs, required fields)
- Professional formatted output with box-drawing characters
- CLI menu system with navigation
- Welcome banner with system status

### Technical
- Applied Go concepts from Chapters 1-7
- Proper error handling throughout
- Pointer receivers for state modification
- Constructor functions with validation
- Comma-ok idiom for map lookups
- Constants for validation rules

[unreleased]: https://github.com/7-Dany/production-planner/compare/v0.0.1...HEAD
[0.0.1]: https://github.com/7-Dany/production-planner/releases/tag/v0.0.1
