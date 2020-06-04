import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import VueSplit from 'vue-split-panel'

Vue.config.productionTip = false;
Vue.config.devtools = true;

Vue.use(VueSplit)

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
