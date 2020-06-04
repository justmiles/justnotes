<template>
  <div class="container">
    <div class="centered">
      <button @click="openFile" type="button">Open</button>
    </div>

    <div class="split right">
      <textarea id="mde"></textarea>
    </div>
  </div>
</template>

<script>
import EasyMDE from "easymde";
// import * as Wails from '@wailsapp/runtime';

export default {
  data() {
    return {
      defaultContent: ""
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
    // Custom Toolbars
    var openFileToolbar = {
      name: "custom",
      action: function() {
        window.backend.Backend.OpenFile();
      },
      className: "fa fa-folder-open",
      title: "Open File"
    };

    window.backend.Backend.GetDefaultContent().then(defaultContent => {
      var easyMDE = new EasyMDE({
        autoDownloadFontAwesome: true,
        element: document.getElementById("mde"),
        initialValue: atob(defaultContent),
        toolbar: [
          openFileToolbar,
          "preview",
          "|",
          "image",
          "link",
          "|",
          "bold",
          "italic",
          "heading",
          "|",
          "quote",
          "unordered-list",
          "ordered-list",
          "table"
        ],
        minHeight: "500px",
        autofocus: true,
        spellChecker: true,
        promptURLs: true,
        uploadImage: true,
        renderingConfig: {
          codeSyntaxHighlighting: true
        },
        imageUploadFunction: function(
          file,
          onImageUploadSuccess,
          onImageUploadError
        ) {
          console.log(file.name);
          console.log(file);
          console.log(onImageUploadSuccess);
          console.log(onImageUploadError);
          // window.backend.Backend.ImageUploadFunction(file)
        },
        sideBySideFullscreen: true,
        shortcuts: {
          toggleOrderedList: "Ctrl-Alt-K", // alter the shortcut for toggleOrderedList
          toggleCodeBlock: null, // unbind Ctrl-Alt-C
          drawTable: "Ctrl-T" // bind Cmd-T to drawTable action, which doesn't come with a default shortcut
          // "openFile": "Ctrl-S", // bind Ctrl-S to save? null for now
          // "toggleSideBySide": "Ctrl-Shift-P", // bind Ctrl-S to save? null for now
        }
      });

      easyMDE.toggleFullScreen();

      // Send changes to Wails backend
      easyMDE.codemirror.on("change", function() {
        window.wails.Events.Emit("file-changed", btoa(easyMDE.value()));
      });

      // Handle Wails file-opened events
      window.wails.Events.On("file-opened", result => {
        easyMDE.value(atob(result));
      });
    });
  }
};
</script>

<style>
.CodeMirror-code {
  font-family: Courier, monospace;
}
</style>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* Split the screen in half */
.split {
  height: 100%;
  width: 90%;
  position: fixed;
  z-index: 1;
  top: 0;
  overflow-x: hidden;
  padding-top: 20px;
}

/* Control the left side */
.left {
  left: 0;
  background-color: #111;
}

/* Control the right side */
.right {
  right: 0;
  background-color: red;
}

/* If you want the content centered horizontally and vertically */
.centered {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
}

/* Style the image inside the centered container, if needed */
.centered img {
  width: 150px;
  border-radius: 50%;
}
</style>
