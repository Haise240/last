import { library } from '@fortawesome/fontawesome-svg-core'
import { faHome, faRoute, faInfoCircle, faPhone } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faHome, faRoute, faInfoCircle, faPhone)

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component('font-awesome-icon', FontAwesomeIcon)
})
