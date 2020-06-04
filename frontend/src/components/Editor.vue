<template>
  <span>
    <Split :style="windowHeight" :gutterSize="4">
      <SplitArea :size="10">
        <treeselect
          v-model="value"
          :options="files"
          :searchable="true"
          :clearable="false"
          :multiple="false"
          :default-expand-level="1"
          :flatten-search-results="false"
          placeholder="Select File..."
        >
          <label
            slot="option-label"
            slot-scope="{ node, shouldShowCount, count, labelClassName, countClassName }"
            :class="labelClassName"
          >
            {{ node.isBranch ? '📂' : '📜' }}: {{ node.label }}
            <span
              v-if="shouldShowCount"
              :class="countClassName"
            >({{ count }})</span>
          </label>
        </treeselect>
      </SplitArea>
      <SplitArea :size="90">
        <textarea id="mde"></textarea>
      </SplitArea>
    </Split>
  </span>
</template>

<script>
import EasyMDE from "easymde";
// import * as Wails from '@wailsapp/runtime';
// using ES6 modules
// import VueSplit from "vue-split-panel";
import Treeselect from "@riophae/vue-treeselect";
import "@riophae/vue-treeselect/dist/vue-treeselect.css";

export default {
  data() {
    return {
      defaultContent: "",
      windowHeight: "height: 800px;",
      files: [
        {
          id: "a",
          label: "http-signature",
          children: [
            {
              id: "r",
              label: "readme.md"
            },
            {
              id: "ab",
              label: "ab"
            }
          ]
        },
        {
          id: "b",
          label: "b"
        },
        {
          id: "c",
          label: "c"
        }
      ]
    };
  },
  components: {
    Treeselect
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

      // easyMDE.toggleFullScreen();

      // Send changes to Wails backend
      easyMDE.codemirror.on("change", function() {
        window.wails.Events.Emit("file-changed", btoa(easyMDE.value()));
      });

      // Handle Wails file-opened events
      window.wails.Events.On("file-opened", result => {
        easyMDE.value(atob(result));
      });
    });
  },
  // Handle Window size
  beforeDestroy() {
    window.removeEventListener("resize", this.onResize);
  },
  onResize() {
    this.windowHeight = "height: " + window.innerHeight + "px;";
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
.labelClassName {
}
</style>
