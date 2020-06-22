<template>
  <div id="app">
    <AddTodo v-on:add-todo="addTodo" />
  </div>
</template>

<script>
  import AddTodo from "./components/AddTodo";
  import axios from "axios";

  export default {
    name: "TodoListApp",
    components: {
      AddTodo
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
      }
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
