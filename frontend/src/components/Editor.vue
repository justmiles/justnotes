<template>
  <span>
    <span v-shortkey.once="['f11']" @shortkey="toggleFullScreen()"></span>
    <span v-shortkey.once="['ctrl', 'o']" @shortkey="openFile()"></span>
    <textarea id="mde"></textarea>
  </span>
</template>

<script>
import EasyMDE from "easymde";

export default {
  methods: {
    openFile() {
      var self = this;
      window.backend.Backend.OpenFile().then(result => {
        self.message = result;
        self.unFullscreen();
      });
    },
    toggleFullScreen() {
      window.backend.Backend.ToggleFullScreen();
    }
  },
  mounted() {
    var self = this;
    // Custom Toolbars
    var openFileToolbar = {
      name: "custom",
      action: function() {
        window.backend.Backend.OpenFile();
      },
      className: "fa fa-folder-open",
      title: "Open File"
    };

    var fullscreen = {
      name: "custom",
      action: function() {
        self.toggleFullScreen();
      },
      className: "fa fa-arrows-alt",
      title: "Fullscreen"
    };

    window.backend.Backend.GetDefaultContent().then(defaultContent => {
      self.easyMDE = new EasyMDE({
        autoDownloadFontAwesome: true,
        element: document.getElementById("mde"),
        initialValue: atob(defaultContent),
        toolbar: [
          openFileToolbar,
          "preview",
          fullscreen,
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
        minHeight: window.innerHeight - 120 + "px",
        autofocus: true,
        spellChecker: true,
        promptURLs: false,
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
        },
        sideBySideFullscreen: true,
        shortcuts: {
          drawTable: "Ctrl-T" // bind Cmd-T to drawTable action
        }
      });

      self.easyMDE.toggleFullScreen();

      // Send changes to Wails backend
      self.easyMDE.codemirror.on("change", function() {
        window.wails.Events.Emit("file-changed", btoa(self.easyMDE.value()));
      });

      // Handle Wails file-opened events
      window.wails.Events.On("file-opened", result => {
        self.easyMDE.value(atob(result));
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