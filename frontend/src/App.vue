<template>
  <div id="app">
    <AddTodo v-on:add-todo="addTodo" />
    <Todos
      v-bind:incomplete-todos="incompleteTodos"
      v-bind:complete-todos="completeTodos"
      v-on:del-todo="deleteTodo"
      v-on:done-todo="toggleTodo"
    />
  </div>
</template>

<script>
  import AddTodo from "./components/AddTodo";
  import Todos from "./components/Todos";
  import axios from "axios";

  export default {
    name: "TodoListApp",
    components: {
      AddTodo,
      Todos
    },
    data() {
      return {
        completeTodos: [],
        incompleteTodos: []
      };
    },
    methods: {
      addTodo(description) {
        axios
          .post("http://localhost:8000/addTodo", null, {
            params: { description }
          })
          .then(
            resp => (this.incompleteTodos = [resp.data, ...this.incompleteTodos])
          )
          .catch(err => console.log(err));
      },
      deleteTodo(id) {
        axios
          .delete("http://localhost:8000/todo/" + id)
          .then(res => {
            if (res.data.deleted) {
              this.completeTodos = this.completeTodos.filter(
                todo => todo.id !== id
              );
              this.incompleteTodos = this.incompleteTodos.filter(
                todo => todo.id !== id
              );
            }
          })
          .catch(err => console.log(err));
      },
      toggleTodo(id) {
        let remove_list = this.incompleteTodos;
        let add_list = this.completeTodos;
        let todo_item = remove_list.find(el => el.id == id);
        if (todo_item === undefined) {
          [remove_list, add_list] = [add_list, remove_list];
          todo_item = remove_list.find(el => el.id == id);
          if (todo_item === undefined) {
            console.log("Can't find any item with id: " + id);
            return;
          }
        }
        axios
          .post("http://localhost:8000/todo/" + todo_item.id, null, {
            params: { done: !todo_item.done }
          })
          .then(res => {
            if (res.data.updated) {
              todo_item.done = !todo_item.done;
              remove_list.splice(remove_list.indexOf(todo_item), 1);
              add_list.unshift(todo_item);
            }
          })
          .catch(err => console.log(err));
      }
    },
    created() {
      axios
        .get("http://localhost:8000/todo")
        .then(res => {
          this.completeTodos = res.data
            .filter(todo => todo.done)
            .sort()
            .reverse();
          this.incompleteTodos = res.data
            .filter(todo => !todo.done)
            .sort()
            .reverse();
        })
        .catch(err => console.log(err));
    }
  };
</script>

<style>
  @import "./assets/reset.css";
  @import url("https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&display=swap");

  body,
  input,
  button {
    font-family: "Roboto", sans-serif;
  }

  body {
    background: #edf0f1;
  }

  .noFill {
    fill: none;
  }

  @media only screen and (min-width: 515px) {
    html {
      background: #e8e6f3;
    }
    body {
      max-width: 515px;
      margin: 0 auto;
      padding: 0 15px;
      border-style: none solid;
      border-width: 2px;
      border-color: #dedbef;
    }
  }
</style>
