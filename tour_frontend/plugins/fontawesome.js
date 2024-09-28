import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

// Импорт нужных иконок
import { fas } from '@fortawesome/free-solid-svg-icons';
import { faHome, faRoute, faInfoCircle, faPhone } from '@fortawesome/free-solid-svg-icons';

library.add(fas, faHome, faRoute, faInfoCircle, faPhone);

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('font-awesome-icon', FontAwesomeIcon);
});
