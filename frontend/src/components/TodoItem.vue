<template>
  <div class="container">
    <div class="description">{{ todo.desc }}</div>
    <div class="buttons">
      <button class="remove" v-on:click="$emit('del-todo', todo.id)">
        <svg
          enable-background="new 0 0 22 22"
          version="1.1"
          viewBox="0 0 22 22"
          xml:space="preserve"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            class="removeFill"
            d="m16.1 3.6h-1.9v-0.3c0-1.3-1-2.3-2.3-2.3h-1.7c-1.3 0-2.4 1-2.4 2.3v0.2h-1.9c-1.3 0-2.3 1-2.3 2.3v1.3c0 0.5 0.4 0.9 0.9 1v10.5c0 1.3 1 2.3 2.3 2.3h8.5c1.3 0 2.3-1 2.3-2.3v-10.4c0.5-0.1 0.9-0.5 0.9-1v-1.3c-0.1-1.3-1.1-2.3-2.4-2.3zm-7-0.3c0-0.6 0.5-1.1 1.1-1.1h1.7c0.6 0 1.1 0.5 1.1 1.1v0.2h-3.9v-0.2zm7.2 15.4c0 0.6-0.5 1.1-1.1 1.1h-8.5c-0.6 0-1.1-0.5-1.1-1.1v-10.5h10.6l0.1 10.5zm0.9-11.7h-12.4v-1.1c0-0.6 0.5-1.1 1.1-1.1h10.2c0.6 0 1.1 0.5 1.1 1.1v1.1z"
          />
          <path
            class="removeFill"
            d="m11 18c-0.4 0-0.6-0.3-0.6-0.6v-6.8c0-0.4 0.3-0.6 0.6-0.6s0.6 0.3 0.6 0.6v6.8c0 0.3-0.2 0.6-0.6 0.6z"
          />
          <path
            class="removeFill"
            d="M8,18c-0.4,0-0.6-0.3-0.6-0.6v-6.8C7.4,10.2,7.7,10,8,10c0.4,0,0.6,0.3,0.6,0.6v6.8C8.7,17.7,8.4,18,8,18z"
          />
          <path
            class="removeFill"
            d="m14 18c-0.4 0-0.6-0.3-0.6-0.6v-6.8c0-0.4 0.3-0.6 0.6-0.6 0.4 0 0.6 0.3 0.6 0.6v6.8c0 0.3-0.3 0.6-0.6 0.6z"
          />
        </svg>
      </button>
      <button
        class="done"
        v-on:click="$emit('done-todo',  todo.id)"
        v-bind:class="{'completed': todo.done}"
      >
        <svg
          version="1.1"
          xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink"
          x="0px"
          y="0px"
          viewBox="0 0 22 22"
          style="enable-background: new 0 0 22 22;"
          xml:space="preserve"
        >
          <rect y="0" class="noFill" width="22" height="22" />
          <g>
            <path
              class="doneFill"
              d="M9.7,14.4L9.7,14.4c-0.2,0-0.4-0.1-0.5-0.2l-2.7-2.7c-0.3-0.3-0.3-0.8,0-1.1s0.8-0.3,1.1,0l2.1,2.1l4.8-4.8 c0.3-0.3,0.8-0.3,1.1,0s0.3,0.8,0,1.1l-5.3,5.3C10.1,14.3,9.9,14.4,9.7,14.4z"
            />
          </g>
        </svg>
      </button>
    </div>
  </div>
</template>

<script>
  export default {
    name: "TodoItem",
    props: ["todo"]
  };
</script>

<style scoped>
  .container {
    display: flex;
    align-items: center;
    flex-wrap: nowrap;
  }
  .description {
    flex-grow: 1;
  }
  .buttons {
    width: 100px;
    height: 50px;
    flex-wrap: nowrap;
    display: flex;
  }

  .buttons button {
    width: 50px;
    height: 50px;
    background: none;
    position: relative;
    border: 0;
    box-shadow: none;
    outline: none;
    cursor: pointer;

    -webkit-appearance: none;
    -moz-appearance: none;
    -ms-appearance: none;
    -o-appearance: none;
    appearance: none;
  }

  .buttons button:last-of-type:before {
    content: "";
    width: 1px;
    height: 30px;
    background: #edf0f1;
  }

  .buttons button svg {
    width: 22px;
    height: 22px;
    position: absolute;
    top: 50%;
    left: 50%;
    margin: -11px 0 0 -11px;
  }

  /* Done Button Style */
  .buttons button.done svg {
    border-radius: 11px;
    border: 1.2px solid rgb(191, 187, 222);
    transition: background 0.2s ease-in-out;
  }
  .buttons button.done svg .doneFill {
    fill: rgb(127, 120, 190);
  }
  .buttons button.done:hover svg {
    background: rgba(127, 120, 190, 0.5);
  }
  .buttons button.done:hover svg .doneFill {
    fill: #fff;
  }
  .buttons button.done.completed svg {
    background: rgb(127, 120, 190);
    border: 0px;
  }
  .buttons button.done.completed svg .doneFill {
    fill: #fff;
  }

  /* Remove Button Style */
  .buttons button.remove svg .removeFill {
    fill: #c0cecb;
  }
  .buttons button.remove:hover svg .removeFill {
    fill: #e85656;
  }
</style>
