<template>
  <div class="container">
    <input type="text" placeholder="I need to do.." v-model="desc" v-on:keyup.enter="addTodo" />
    <button id="add" v-on:click="addTodo">
      <svg
        version="1.1"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        x="0px"
        y="0px"
        viewBox="0 0 16 16"
        style="enable-background: new 0 0 16 16;"
        xml:space="preserve"
      >
        <g>
          <path
            class="addFill"
            d="M16,8c0,0.5-0.5,1-1,1H9v6c0,0.5-0.5,1-1,1s-1-0.5-1-1V9H1C0.5,9,0,8.5,0,8s0.5-1,1-1h6V1c0-0.5,0.5-1,1-1
			s1,0.5,1,1v6h6C15.5,7,16,7.5,16,8z"
          />
        </g>
      </svg>
    </button>
  </div>
</template>

<script>
  export default {
    name: "AddTodo",
    data() {
      return {
        desc: ""
      };
    },
    methods: {
      addTodo() {
        if (!this.desc) {
          return;
        }
        this.$emit("add-todo", this.desc);
        this.desc = "";
      }
    }
  };
</script>

<style scoped>
  div {
    width: 100%;
    height: 80px;
    padding: 15px;
    z-index: 3;
    background: rgb(127, 120, 190);
    box-shadow: 0px 2px 4px rgba(74, 89, 146, 0.15);
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;
  }

  div input {
    width: 100%;
    height: 50px;
    float: left;
    color: #fff;
    font-size: 15px;
    font-weight: 400;
    text-indent: 18px;
    padding: 0 60px 0 0;
    background: rgba(255, 255, 255, 0.2);
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
    border-top-right-radius: 25px;
    border-bottom-right-radius: 25px;
    border: 0px;
    box-shadow: none;
    outline: none;

    -webkit-appearance: none;
    -moz-appearance: none;
    -ms-appearance: none;
    -o-appearance: none;
    appearance: none;
  }

  div input::-webkit-input-placeholder {
    color: rgba(255, 255, 255, 0.75);
  }

  div input::-moz-input-placeholder {
    color: rgba(255, 255, 255, 0.75);
  }
  div input:-moz-input-placeholder {
    color: rgba(255, 255, 255, 0.75);
  }
  div input::-ms-input-placeholder {
    color: rgba(255, 255, 255, 0.75);
  }

  div button {
    width: 50px;
    height: 50px;
    position: absolute;
    top: 15px;
    right: 15px;
    z-index: 2;
    border-radius: 25px;
    background: #fff;
    border: 0px;
    cursor: pointer;
    box-shadow: none;
    outline: none;

    -webkit-appearance: none;
    -moz-appearance: none;
    -ms-appearance: none;
    -o-appearance: none;
    appearance: none;
  }

  @media only screen and (min-width: 515px) {
    div button {
      position: relative;
      top: -100%;
      left: 90%;
    }
  }

  div button svg {
    width: 16px;
    height: 16px;
    position: absolute;
    top: 50%;
    left: 50%;
    margin: -8px 0 0 -8px;
  }

  div button svg .addFill {
    fill: rgb(127, 120, 190);
  }
</style>