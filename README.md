# Internship Project

## Description
Your task is to build a basic TODO List API that allows users to perform CRUD operations (Create, Read, Update, Delete) involving TODO items. The API should follow RESTful conventions and should be able to handle JSON data for requests and responses.

## Requirements
1. The API should have endpoints for the following operations:
   - Create a new todo item
   - Retrieve a list of all todo items
   - Retrieve a single todo item by its ID
   - Update a todo item by its ID
   - Delete a todo item by its ID

2. The todo item should have at least the following attributes:
   - ID (unique identifier)
   - Title
   - Description
   - Completion status

3. Use HTTP methods to perform actions:
   - POST for creating
   - GET for reading
   - PUT for updating
   - DELETE for deleting

4. Use a simple, in-memory storage solution (e.g., a slice) for storing todo items. Or you can use a file, to not lose your state between restarts and changes. No need for a database (we will leave this for the future 😉)

5. Try to implement error handling for various scenarios such as invalid requests or non-existent resources.

6. Try to apply your newly aquired knowledge about structures to implement input data validations (e.g., required fields).

7. Try to return appropriate status codes in the API responses
    - 200 OK (ex: when fetching TODO list)
    - 201 Created (ex: when you create a TODO item)
    - 404 Not Found (ex: when you try to edit an item that doesnt exist)
    - 500 Internal Server Error (ex: unexpected error)
    - etc

8. Add tests! Your API code should have unit tests for, at least, the most important parts of your API. The more you can test your code, the better. While making your manual tests help you during development, you want to have written tests so you can avoid regression while working on improvements on a later date (which we will do)

8. Create a documentation for how to use the API, including examples of requests and responses. Someone reading the documenation should be able to follow the exemples and understand how to do their own TODO list.

## Sample Endpoint Structure
- POST /todo (create a new TODO item)
- GET /todo (retrieve all TODO items)
- GET /todo/{id} (retrieve a single TODO item by ID)
- PUT /todo/{id} (update a TODO item by ID)
- DELETE /todo/{id} (delete a TODO item by ID)

## Goals/Milestones

You can organize the API development however you like, make sure to try different things to find what makes you confortable.
With that being said, I have a suggestion on how to break it smaller parts

### **Milestone 1: Project Setup and Planning**
   - Set up your Git (you should fork this project).
   - Set up the project structure.
   - Define the data model for the TODO items (your golang struct).
   - Plan out the API endpoints and overall architecture (exemples of endpoints below)
    - POST /todo (create a new TODO item)
    - GET /todo (retrieve all TODO items)
    - GET /todo/{id} (retrieve a single TODO item by ID)
    - PUT /todo/{id} (update a TODO item by ID)
    - DELETE /todo/{id} (delete a TODO item by ID)

### **Milestone 2: Basic API Implementation**
   - Implement basic HTTP server (already done as an "Hello World" exemple).
   - Create endpoints for retrieving all TODO items and retrieving a single TODO item by its ID.
   - Implement in-memory or in file storage for storing TODO items.
   - Write basic tests for the implemented endpoints.

### **Milestone 3: Create and Update Functionality**
   - Implement endpoints for creating a new TODO item and updating a TODO item by its ID.
   - Add validation for input data (e.g., required fields, data types).
   - Enhance error handling to handle edge cases and invalid requests.
   - Write tests to ensure the create and update functionality works as expected.

### **Milestone 4: Delete Functionality and Error Handling**
   - Implement the endpoint for deleting a todo item by its ID.
   - Improve error handling throughout the API, providing informative error messages that help the user to fix them.
   - Write tests to verify the delete functionality and error handling.

### **Milestone 5: Documentation and API Refinement**
   - Document the API, including endpoint descriptions, request/response formats, and usage examples.
   - Refactor and/or optimize the code as needed. During development you might have gained new insights that can be used now.
   - Review of the API to ensure it meets the project requirements. See if nothing is missing or if you can improve anything or even add something extra.

### **Milestone 6: Testing and Bug Fixing**
   - Perform a full testing of the API, with all endpoints and a full walkthrough of how using this API would look like (create TODOs, list them, edit them, remove them, etc)
   - Identify and fix any bugs or issues discovered during testing.
   - Ensure the API behaves as expected in various scenarios and under different conditions. Try to break it 😉

### **Milestone 7: Final Review**
   - Conduct a final review of the codebase, documentation, and overall project.
   - Address any remaining issues or feedback received during the review process.
   - Follow your own documentation and see if everything is working as expected


## Technologies

You will have to do this project using Golang, as thats the main language the team uses AND thats the one you have been working on.
As for packages you are free to use whatever you think will help you on achieving your goals/milestones
