function showEditForm(todo) {
  document.getElementById("todo-id").value = todo.id; // Set the todo ID
  document.getElementById("todo-title").value = todo.title; // Set the title
  document.getElementById("todo-due-date").value = todo.dueDate; // Set the due date
  document.getElementById("form-title").textContent = "Edit Todo"; // Change form title

  // Set the form action to PUT for editing
  document
    .getElementById("todo-form")
    .setAttribute("hx-put", `/todos/${todo.id}`);

  // Adjust the submit button text or attributes if necessary
  document.getElementById("submit-button").innerText = "Update Todo";
}

function resetForm() {
  document.getElementById("todo-id").value = ""; // Clear the todo ID
  document.getElementById("todo-title").value = ""; // Clear the title
  document.getElementById("todo-due-date").value = ""; // Clear the due date
  document.getElementById("form-title").textContent = "Add Todo"; // Reset form title

  // Reset the form action to POST for adding a new todo
  document.getElementById("todo-form").setAttribute("hx-post", "/todos");

  // Set the submit button back to Add
  document.getElementById("submit-button").innerText = "Add Todo";
}
