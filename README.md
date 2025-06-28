# simple-registration
A simple command line based people registration service

# Requirements
All requirements were created by claude AI

## Functional Requirements

- Struct "Person" with fields name,birthDate,email,phone,address
- Features: add, search, list, update, delete
- Validations:
    1. valid email 
    2. positive age 
    3. phone number size(brazilizan phone number size used as reference)
    4. Non empty name and address
- Save data in JSON file
- Interactive menu interface

## Non Functional Requirements
- Use slices to store people
- Implement efficcient search (O(n))
- Robust Data Validation
- Modular and reusable code
- error handling when necessary
- suport up to 1000 registers
