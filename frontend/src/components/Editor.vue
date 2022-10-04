

<script lang="ts" setup>
import {reactive} from 'vue'
import {InitialData} from '../../wailsjs/go/main/App'
import EasyMDE from "easymde"
import "../assets/css/easymde.min.css";
import "../assets/themes/gitbook-azure.css";

const data = reactive({
  name: "",
  mde: HTMLElement,
  resultText: "Please enter your name below 👇",
})

InitialData().then(result => {
  const easyMDE = new EasyMDE({
    element: document.getElementById("mde")!,
    initialValue: result,
    toolbar: ["preview", "side-by-side", "undo", "redo"],
    autoDownloadFontAwesome: true,
    autofocus: true,
    spellChecker: false,
    promptURLs: false,
    sideBySideFullscreen: false,
    shortcuts: {
      drawTable: "Ctrl-T" // bind Cmd-T to drawTable action
    }
  });

  EasyMDE.togglePreview(easyMDE);
  EasyMDE.toggleFullScreen(easyMDE);

})

</script>

<template>
  <main>
    <textarea id="mde" ></textarea>
  </main>
</template>
