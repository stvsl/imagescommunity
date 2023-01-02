import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueCookies from 'vue-cookies'
import { library } from '@fortawesome/fontawesome-svg-core'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// import UndrawUi from 'undraw-ui'
// import 'undraw-ui/dist/style.css'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
router.afterEach(() => {
    window.scrollTo(0, 0);
});
// 首页头部图标
import { faMagnifyingGlass, faArrowUpFromBracket, faStar, faImage, faUser, faListUl, faBars } from '@fortawesome/free-solid-svg-icons'
// 回到顶部图标
import { faArrowUp } from '@fortawesome/free-solid-svg-icons'
import { faEye, faHeart } from '@fortawesome/free-solid-svg-icons'
import { faFlickr } from '@fortawesome/free-brands-svg-icons'

library.add(faMagnifyingGlass, faArrowUpFromBracket, faStar, faImage, faUser, faListUl, faBars, faArrowUp)
library.add(faFlickr, faEye, faHeart)

createApp(App).
    use(store)
    .use(router)
    .use(ElementPlus)
    .use(VueCookies)
    .component('font-awesome-icon', FontAwesomeIcon)
    .mount('#app')
