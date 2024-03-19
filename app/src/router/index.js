import { createRouter, createWebHistory } from 'vue-router'
import NoteEditorView from '../views/NoteEditorView.vue';
import LoginView from '../views/LoginView.vue';
import NoteIndexView from '../views/NoteIndexView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/notes',
      name: 'noteIndex',
      component: NoteIndexView
    },
    {
      path: '/:user/:filename',
      name: 'editor',
      component: NoteEditorView
    }
  ]
})

export default router
