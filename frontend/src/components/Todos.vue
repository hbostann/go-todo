<template>
  <div style="padding: 15px;">
    <ul class="todo incomplete">
      <li v-bind:key="todo.id" v-for="todo in incompleteTodos">
        <TodoItem v-bind:todo="todo" />
      </li>
    </ul>
    <hr />
    <ul class="todo complete">
      <li v-bind:key="todo.id" v-for="todo in completeTodos">
        <TodoItem v-bind:todo="todo" />
      </li>
    </ul>
  </div>
</template>

<script>
  import TodoItem from "./TodoItem";

  export default {
    name: "Todos",
    components: {
      TodoItem
    },
    props: {
      incompleteTodos: Array,
      completeTodos: Array
    }
  };
</script>

<style scoped>
  ul.todo {
    width: 100%;
  }

  ul.todo li {
    width: 100%;
    min-height: 50px;
    font-weight: 500;
    color: #444;
    line-height: 22px;

    background: #fff;
    border-radius: 5px;
    box-shadow: 0px 1px 2px rgba(44, 62, 80, 0.1);
    margin: 0 0 10px 0;
    padding: 2px 2px 2px 20px;
  }

  ul.todo li:last-of-type {
    margin: 0;
  }

  hr {
    background: #d8e5e0;
    margin: 15px;
  }

  ul.todo.incomplete:empty:after {
    content: "You have nothing todo!";
    margin: 15px 0 0 0;
  }

  ul.todo.complete:empty:after {
    content: "Go on, do something!";
  }

  ul.todo.incomplete:after,
  ul.todo.complete:after {
    width: 100%;
    display: block;
    text-align: center;
    font-size: 12px;
    color: #aaa;
  }
</style>