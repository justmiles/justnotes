<template>
  <div class="container">
    <button @click="openFile" type="button">Open</button>
    <textarea id="mde"></textarea>

  </div>
</template>

<script>

import EasyMDE from 'easymde';
// import * as Wails from '@wailsapp/runtime';

export default {
  data() {
    return {
      message: " "
    };
  },
  methods: {
    openFile: function() {
      var self = this;
      window.backend.Backend.OpenFile().then(result => {
        self.message = result;
      });
    }
  },
  mounted() {
    window.backend.Backend.GetDefaultContent().then(defaultContent => {

    var easyMDE = new EasyMDE({
    autoDownloadFontAwesome: true,
    element: document.getElementById('mde'),
    initialValue: atob(defaultContent),
    toolbar: [
      {
			name: "custom",
			action: function(){
        window.backend.Backend.OpenFile()
			},
			className: "fa fa-folder-open",
			title: "Open File",
		}, "preview", "|", "image", "link", "|", "bold", "italic", "heading", "|", "quote", "unordered-list", "ordered-list", "table"],
    
    minHeight: "500px",
    // toolbar: ["bold", "italic", "heading", "|", "quote"],
    sideBySideFullscreen: true,
    shortcuts: {
      "toggleOrderedList": "Ctrl-Alt-K", // alter the shortcut for toggleOrderedList
      "toggleCodeBlock": null, // unbind Ctrl-Alt-C
      "drawTable": "Ctrl-T", // bind Cmd-T to drawTable action, which doesn't come with a default shortcut
      // "openFile": "Ctrl-S", // bind Ctrl-S to save? null for now
      // "toggleSideBySide": "Ctrl-Shift-P", // bind Ctrl-S to save? null for now
    }
    });
    console.log(easyMDE)
    easyMDE.toggleFullScreen()
    easyMDE.codemirror.on("change", function(){
      window.wails.Events.Emit("file-saved", btoa(easyMDE.value()))
    });
    
    // Wails 
    window.wails.Events.On("file-opened", result => {
      console.log(result)
      easyMDE.value(atob(result))
    });


    });

  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
</style>
