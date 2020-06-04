<template>
  <div id="app">
    <Editor/>
  </div>
</template>

<script>
import Editor from "./components/Editor.vue";
import "./assets/css/main.css";
import './assets/css/easymde.min.css'

export default {
  name: "app",
  components: {
    Editor
  }
};
</script>
