import { createRouter, createWebHistory } from 'vue-router'
import NoteEditorView from '../views/NoteEditorView.vue';
import LoginView from '../views/LoginView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/editor',
      name: 'editor',
      component: NoteEditorView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    }
  ]
})

export default router
